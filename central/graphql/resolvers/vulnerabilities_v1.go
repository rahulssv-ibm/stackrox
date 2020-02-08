package resolvers

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/nvdtools/cvefeed/nvd/schema"
	protoTypes "github.com/gogo/protobuf/types"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	clusterMappings "github.com/stackrox/rox/central/cluster/index/mappings"
	"github.com/stackrox/rox/central/cve/converter"
	"github.com/stackrox/rox/central/graphql/resolvers/loaders"
	"github.com/stackrox/rox/central/image/mappings"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	pkgMetrics "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/predicate"
)

var (
	vulnPredicateFactory = predicate.NewFactory("vulnerability", &storage.EmbeddedVulnerability{})
)

// V1 Connections to root.
//////////////////////////

// Vulnerability resolves a single vulnerability based on an id (the CVE value).
func (resolver *Resolver) vulnerabilityV1(ctx context.Context, args idQuery) (VulnerabilityResolver, error) {
	query := search.NewQueryBuilder().AddExactMatches(search.CVE, string(*args.ID)).ProtoQuery()
	vulns, err := vulnerabilities(ctx, resolver, query)
	if err != nil {
		return nil, err
	} else if len(vulns) == 0 {
		return nil, nil
	} else if len(vulns) > 1 {
		return nil, fmt.Errorf("multiple vulns matched: %s this should not happen", string(*args.ID))
	}
	return vulns[0], nil
}

// Vulnerabilities resolves a set of vulnerabilities based on a query.
func (resolver *Resolver) vulnerabilitiesV1(ctx context.Context, q PaginatedQuery) ([]VulnerabilityResolver, error) {
	// Convert to query, but link the fields for the search.
	query, err := q.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	resolvers, err := paginationWrapper{
		pv: query.Pagination,
	}.paginate(vulnerabilities(ctx, resolver, query))
	vulnRes := resolvers.([]*EmbeddedVulnerabilityResolver)

	ret := make([]VulnerabilityResolver, 0, len(vulnRes))
	for _, resolver := range vulnRes {
		ret = append(ret, resolver)
	}
	return ret, err
}

// VulnerabilityCount returns count of all clusters across infrastructure
func (resolver *Resolver) vulnerabilityCountV1(ctx context.Context, args RawQuery) (int32, error) {
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	vulns, err := vulnerabilities(ctx, resolver, query)
	if err != nil {
		return 0, err
	}
	return int32(len(vulns)), nil
}

func (resolver *Resolver) vulnCounterV1(ctx context.Context, args RawQuery) (*VulnerabilityCounterResolver, error) {
	vulnResolvers, err := resolver.vulnerabilitiesV1(ctx, PaginatedQuery{Query: args.Query})
	if err != nil {
		return nil, err
	}

	var vulns []*storage.EmbeddedVulnerability
	for _, vulnsResolver := range vulnResolvers {
		vulns = append(vulns, vulnsResolver.(*EmbeddedVulnerabilityResolver).data)
	}
	return mapVulnsToVulnerabilityCounter(vulns), nil
}

func (resolver *Resolver) wrapEmbeddedVulns(embeddedVulns []*storage.EmbeddedVulnerability) []*EmbeddedVulnerabilityResolver {
	evrs := make([]*EmbeddedVulnerabilityResolver, 0, len(embeddedVulns))
	for _, ev := range embeddedVulns {
		evrs = append(evrs, &EmbeddedVulnerabilityResolver{
			root: resolver,
			data: ev,
		})
	}
	return evrs
}

// Helper function that actually runs the queries and produces the resolvers from the images.
func vulnerabilities(ctx context.Context, root *Resolver, query *v1.Query) ([]*EmbeddedVulnerabilityResolver, error) {
	// Get the image loader from the context.
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil, err
	}

	// Run search on images.
	images, err := imageLoader.FromQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	imageVulnResolvers, err := mapImagesToVulnerabilityResolvers(root, images, query)
	if err != nil {
		return nil, err
	}

	k8sVulnResolvers, err := k8sIstioVulnerabilities(ctx, root, query, converter.K8s)
	if err != nil {
		return nil, err
	}

	istioVulnResolvers, err := k8sIstioVulnerabilities(ctx, root, query, converter.Istio)
	if err != nil {
		return nil, err
	}

	imageVulnResolvers = append(imageVulnResolvers, k8sVulnResolvers...)
	imageVulnResolvers = append(imageVulnResolvers, istioVulnResolvers...)
	return imageVulnResolvers, nil
}

// k8sIstioVulnerabilities returns the k8s/istio vulnerabilities that match the input query.
func k8sIstioVulnerabilities(ctx context.Context, root *Resolver, query *v1.Query, ct converter.CVEType) ([]*EmbeddedVulnerabilityResolver, error) {
	var cves []*storage.CVE
	_, containsUnmatchableFields := search.FilterQueryWithMap(query, search.CombineOptionsMaps(clusterMappings.OptionsMap, mappings.VulnerabilityOptionsMap))
	if containsUnmatchableFields {
		return nil, nil
	}

	vulnQuery, _ := search.FilterQueryWithMap(query, mappings.VulnerabilityOptionsMap)
	vulnPred, err := vulnPredicateFactory.GeneratePredicate(vulnQuery)
	if err != nil {
		return nil, err
	}

	clusterQuery, _ := search.FilterQueryWithMap(query, clusterMappings.OptionsMap)
	clusters, err := root.ClusterDataStore.SearchRawClusters(ctx, clusterQuery)
	if err != nil {
		return nil, err
	}

	var checkImpact func(context.Context, *storage.Cluster, *schema.NVDCVEFeedJSON10DefCVEItem) (bool, error)

	if ct == converter.K8s {
		cves, err = root.k8sIstioCVEManager.GetK8sCVEs(ctx, vulnQuery)
		if err != nil {
			return nil, err
		}
		checkImpact = func(ctx context.Context, cluster *storage.Cluster, cve *schema.NVDCVEFeedJSON10DefCVEItem) (bool, error) {
			ok := root.cveMatcher.IsClusterAffectedByK8sCVE(cluster, cve)
			return ok, nil
		}
	} else if ct == converter.Istio {
		cves, err = root.k8sIstioCVEManager.GetIstioCVEs(ctx, vulnQuery)
		if err != nil {
			return nil, err
		}
		checkImpact = func(ctx context.Context, cluster *storage.Cluster, cve *schema.NVDCVEFeedJSON10DefCVEItem) (bool, error) {
			ok, err := root.cveMatcher.IsClusterAffectedByIstioCVE(ctx, cluster, cve)
			return ok, err
		}
	} else {
		return nil, errors.Errorf("unknown CVE type: %d", ct)
	}

	ret := make([]*EmbeddedVulnerabilityResolver, 0, len(cves))

	for _, cve := range cves {
		for _, cluster := range clusters {
			nvdCVE := root.getNvdCVE(cve.Id)
			ok, err := checkImpact(ctx, cluster, nvdCVE)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}

			embedded := converter.ProtoCVEToEmbeddedCVE(cve)
			if !features.Dackbox.Enabled() && !vulnPred.Matches(embedded) {
				continue
			}

			r := &EmbeddedVulnerabilityResolver{
				data: embedded,
				root: root,
			}
			ret = append(ret, r)
			break // No need to continue the clusters loop since the CVE was already added to the list.
		}
	}
	return ret, nil
}

func (resolver *Resolver) wrapEmbeddedVulnerability(value *storage.EmbeddedVulnerability, err error) (*EmbeddedVulnerabilityResolver, error) {
	if err != nil {
		return nil, err
	}
	return &EmbeddedVulnerabilityResolver{root: resolver, data: value}, nil
}

// EmbeddedVulnerabilityResolver resolves data about a CVE/Vulnerability.
// If using the top level vulnerability resolver (as opposed to the resolver under the top level image resolver) you get
// a couple of extensions that allow you to resolve some relationships.
type EmbeddedVulnerabilityResolver struct {
	root        *Resolver
	lastScanned *protoTypes.Timestamp
	data        *storage.EmbeddedVulnerability
}

// Vectors returns either the CVSSV2 or CVSSV3 data.
func (evr *EmbeddedVulnerabilityResolver) Vectors() *EmbeddedVulnerabilityVectorsResolver {
	if val := evr.data.GetCvssV3(); val != nil {
		return &EmbeddedVulnerabilityVectorsResolver{
			resolver: &cVSSV3Resolver{evr.root, val},
		}
	}
	if val := evr.data.GetCvssV2(); val != nil {
		return &EmbeddedVulnerabilityVectorsResolver{
			resolver: &cVSSV2Resolver{evr.root, val},
		}
	}
	return nil
}

// ID returns the CVE string (which is effectively an id)
func (evr *EmbeddedVulnerabilityResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(evr.data.GetCve())
}

// Cve returns the CVE string (which is effectively an id)
func (evr *EmbeddedVulnerabilityResolver) Cve(ctx context.Context) string {
	return evr.data.GetCve()
}

// Cvss returns the CVSS score.
func (evr *EmbeddedVulnerabilityResolver) Cvss(ctx context.Context) float64 {
	return float64(evr.data.GetCvss())
}

// Link returns a link to the vulnerability.
func (evr *EmbeddedVulnerabilityResolver) Link(ctx context.Context) string {
	return evr.data.GetLink()
}

// Summary returns the summary of the vulnerability.
func (evr *EmbeddedVulnerabilityResolver) Summary(ctx context.Context) string {
	return evr.data.GetSummary()
}

// ScoreVersion returns the version of the CVSS score returned.
func (evr *EmbeddedVulnerabilityResolver) ScoreVersion(ctx context.Context) string {
	value := evr.data.GetScoreVersion()
	return value.String()
}

// FixedByVersion returns the version of the parent component that removes this CVE.
func (evr *EmbeddedVulnerabilityResolver) FixedByVersion(ctx context.Context) (string, error) {
	return evr.data.GetFixedBy(), nil
}

// IsFixable returns whether or not a component with a fix exists.
func (evr *EmbeddedVulnerabilityResolver) IsFixable(ctx context.Context) (bool, error) {
	return evr.data.GetFixedBy() != "", nil
}

// LastScanned is the last time the vulnerability was scanned in an image.
func (evr *EmbeddedVulnerabilityResolver) LastScanned(ctx context.Context) (*graphql.Time, error) {
	return timestamp(evr.lastScanned)
}

// CreatedAt is the firsts time the vulnerability was scanned in an image. Unavailable in an image context.
func (evr *EmbeddedVulnerabilityResolver) CreatedAt(ctx context.Context) (*graphql.Time, error) {
	return timestamp(evr.lastScanned)
}

// VulnerabilityType returns the type of vulnerability
func (evr *EmbeddedVulnerabilityResolver) VulnerabilityType() string {
	return evr.data.VulnerabilityType.String()
}

// Components are the components that contain the CVE/Vulnerability.
func (evr *EmbeddedVulnerabilityResolver) Components(ctx context.Context, args PaginatedQuery) ([]ComponentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.CVEs, "Components")

	query := search.AddRawQueriesAsConjunction(args.String(), evr.vulnRawQuery())

	return evr.root.Components(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

// ComponentCount is the number of components that contain the CVE/Vulnerability.
func (evr *EmbeddedVulnerabilityResolver) ComponentCount(ctx context.Context, args RawQuery) (int32, error) {
	components, err := evr.Components(ctx, PaginatedQuery{Query: args.Query})
	if err != nil {
		return 0, err
	}
	return int32(len(components)), nil
}

// Images are the images that contain the CVE/Vulnerability.
func (evr *EmbeddedVulnerabilityResolver) Images(ctx context.Context, args PaginatedQuery) ([]*imageResolver, error) {
	// Convert to query, but link the fields for the search.
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}
	images, err := evr.loadImages(ctx, query)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// ImageCount is the number of images that contain the CVE/Vulnerability.
func (evr *EmbeddedVulnerabilityResolver) ImageCount(ctx context.Context, args RawQuery) (int32, error) {
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return 0, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	query, err = search.AddAsConjunction(evr.vulnQuery(), query)
	if err != nil {
		return 0, err
	}
	return imageLoader.CountFromQuery(ctx, query)
}

// Deployments are the deployments that contain the CVE/Vulnerability.
func (evr *EmbeddedVulnerabilityResolver) Deployments(ctx context.Context, args PaginatedQuery) ([]*deploymentResolver, error) {
	if err := readDeployments(ctx); err != nil {
		return nil, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}
	return evr.loadDeployments(ctx, query)
}

// DeploymentCount is the number of deployments that contain the CVE/Vulnerability.
func (evr *EmbeddedVulnerabilityResolver) DeploymentCount(ctx context.Context, args RawQuery) (int32, error) {
	if err := readDeployments(ctx); err != nil {
		return 0, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	deploymentBaseQuery, err := evr.getDeploymentBaseQuery(ctx)
	if err != nil || deploymentBaseQuery == nil {
		return 0, err
	}
	deploymentLoader, err := loaders.GetDeploymentLoader(ctx)
	if err != nil {
		return 0, err
	}
	return deploymentLoader.CountFromQuery(ctx, search.ConjunctionQuery(deploymentBaseQuery, query))
}

func (resolver *Resolver) getNvdCVE(id string) *schema.NVDCVEFeedJSON10DefCVEItem {
	for _, cve := range resolver.k8sIstioCVEManager.GetK8sAndIstioCVEs() {
		if cve.CVE.CVEDataMeta.ID == id {
			return cve
		}
	}
	return nil
}

func (resolver *Resolver) getAffectedClusterPercentage(ctx context.Context, cve *schema.NVDCVEFeedJSON10DefCVEItem, ct converter.CVEType) (float64, error) {
	clusters, err := resolver.ClusterDataStore.GetClusters(ctx)
	if err != nil {
		return 0, err
	}

	var checkImpact func(context.Context, *storage.Cluster, *schema.NVDCVEFeedJSON10DefCVEItem) (bool, error)

	if ct == converter.K8s {
		checkImpact = func(ctx context.Context, cluster *storage.Cluster, entry *schema.NVDCVEFeedJSON10DefCVEItem) (bool, error) {
			return resolver.cveMatcher.IsClusterAffectedByK8sCVE(cluster, cve), nil
		}
	} else if ct == converter.Istio {
		checkImpact = func(ctx context.Context, cluster *storage.Cluster, entry *schema.NVDCVEFeedJSON10DefCVEItem) (bool, error) {
			ok, err := resolver.cveMatcher.IsClusterAffectedByIstioCVE(ctx, cluster, cve)
			return ok, err
		}
	} else {
		return 0.0, errors.Errorf("unknown CVE type: %d", ct)
	}

	affectedClusterCount := 0
	for _, cluster := range clusters {
		ok, err := checkImpact(ctx, cluster, cve)
		if err != nil {
			return 0.0, errors.Errorf("unknown CVE type: %d", ct)
		}
		if ok {
			affectedClusterCount++
		}
	}
	return float64(affectedClusterCount) / float64(len(clusters)), nil
}

func (evr *EmbeddedVulnerabilityResolver) getEnvImpactForK8sIstioVuln(ctx context.Context, ct converter.CVEType) (float64, error) {
	cve := evr.root.getNvdCVE(evr.data.Cve)
	if cve == nil {
		return 0.0, fmt.Errorf("cve: %q not found", evr.data.Cve)
	}
	p, err := evr.root.getAffectedClusterPercentage(ctx, cve, ct)
	if err != nil {
		return 0.0, err
	}
	return p, nil
}

// EnvImpact is the fraction of deployments that contains the CVE
func (evr *EmbeddedVulnerabilityResolver) EnvImpact(ctx context.Context) (float64, error) {
	if evr.data.VulnerabilityType == storage.EmbeddedVulnerability_K8S_VULNERABILITY {
		return evr.getEnvImpactForK8sIstioVuln(ctx, converter.K8s)
	} else if evr.data.VulnerabilityType == storage.EmbeddedVulnerability_ISTIO_VULNERABILITY {
		return evr.getEnvImpactForK8sIstioVuln(ctx, converter.Istio)
	}

	allDepsCount, err := evr.root.DeploymentDataStore.CountDeployments(ctx)
	if err != nil {
		return 0, err
	}
	if allDepsCount == 0 {
		return 0, errors.New("deployments count not available")
	}
	deploymentBaseQuery, err := evr.getDeploymentBaseQuery(ctx)
	if err != nil || deploymentBaseQuery == nil {
		return 0, err
	}
	deploymentLoader, err := loaders.GetDeploymentLoader(ctx)
	if err != nil {
		return 0, err
	}
	withThisCVECount, err := deploymentLoader.CountFromQuery(ctx, search.ConjunctionQuery(deploymentBaseQuery, deploymentBaseQuery))
	if err != nil {
		return 0, err
	}
	return float64(float64(withThisCVECount) / float64(allDepsCount)), nil
}

// Severity return the severity of the vulnerability (CVSSv3 or CVSSv2).
func (evr *EmbeddedVulnerabilityResolver) Severity(ctx context.Context) string {
	if val := evr.data.GetCvssV3(); val != nil {
		return evr.data.GetCvssV3().GetSeverity().String()
	}
	if val := evr.data.GetCvssV2(); val != nil {
		return evr.data.GetCvssV2().GetSeverity().String()
	}
	return storage.CVSSV2_UNKNOWN.String()
}

// PublishedOn is the time the vulnerability was published (ref: NVD).
func (evr *EmbeddedVulnerabilityResolver) PublishedOn(ctx context.Context) (*graphql.Time, error) {
	return timestamp(evr.data.GetPublishedOn())
}

// LastModified is the time the vulnerability was last modified (ref: NVD).
func (evr *EmbeddedVulnerabilityResolver) LastModified(ctx context.Context) (*graphql.Time, error) {
	return timestamp(evr.data.GetLastModified())
}

// ImpactScore returns the impact score of the vulnerability.
func (evr *EmbeddedVulnerabilityResolver) ImpactScore(ctx context.Context) float64 {
	if val := evr.data.GetCvssV3(); val != nil {
		return float64(evr.data.GetCvssV3().GetImpactScore())
	}
	if val := evr.data.GetCvssV2(); val != nil {
		return float64(evr.data.GetCvssV2().GetImpactScore())
	}
	return float64(0.0)
}

func (evr *EmbeddedVulnerabilityResolver) loadImages(ctx context.Context, query *v1.Query) ([]*imageResolver, error) {
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil, err
	}

	pagination := query.GetPagination()
	query.Pagination = nil

	query, err = search.AddAsConjunction(evr.vulnQuery(), query)
	if err != nil {
		return nil, err
	}

	query.Pagination = pagination

	return evr.root.wrapImages(imageLoader.FromQuery(ctx, query))
}

func (evr *EmbeddedVulnerabilityResolver) loadDeployments(ctx context.Context, query *v1.Query) ([]*deploymentResolver, error) {
	deploymentBaseQuery, err := evr.getDeploymentBaseQuery(ctx)
	if err != nil || deploymentBaseQuery == nil {
		return nil, err
	}
	// Search the deployments.
	ListDeploymentLoader, err := loaders.GetListDeploymentLoader(ctx)
	if err != nil {
		return nil, err
	}

	pagination := query.GetPagination()
	query.Pagination = nil

	query, err = search.AddAsConjunction(deploymentBaseQuery, query)
	if err != nil {
		return nil, err
	}

	query.Pagination = pagination

	return evr.root.wrapListDeployments(ListDeploymentLoader.FromQuery(ctx, query))
}

func (evr *EmbeddedVulnerabilityResolver) getDeploymentBaseQuery(ctx context.Context) (*v1.Query, error) {
	imageQuery := evr.vulnQuery()
	results, err := evr.root.ImageDataStore.Search(ctx, imageQuery)
	if err != nil || len(results) == 0 {
		return nil, err
	}

	// Create a query that finds all of the deployments that contain at least one of the infected images.
	var qb []*v1.Query
	for _, id := range search.ResultsToIDs(results) {
		qb = append(qb, search.NewQueryBuilder().AddExactMatches(search.ImageSHA, id).ProtoQuery())
	}
	return search.DisjunctionQuery(qb...), nil
}

func (evr *EmbeddedVulnerabilityResolver) vulnQuery() *v1.Query {
	return search.NewQueryBuilder().AddExactMatches(search.CVE, evr.data.GetCve()).ProtoQuery()
}

func (evr *EmbeddedVulnerabilityResolver) vulnRawQuery() string {
	return search.NewQueryBuilder().AddExactMatches(search.CVE, evr.data.GetCve()).Query()
}

// Static helpers.
//////////////////

// Map the images that matched a query to the vulnerabilities it contains.
func mapImagesToVulnerabilityResolvers(root *Resolver, images []*storage.Image, query *v1.Query) ([]*EmbeddedVulnerabilityResolver, error) {
	vulnQuery, _ := search.FilterQueryWithMap(query, mappings.VulnerabilityOptionsMap)
	vulnPred, err := vulnPredicateFactory.GeneratePredicate(vulnQuery)
	if err != nil {
		return nil, err
	}

	componentQuery, _ := search.FilterQueryWithMap(query, mappings.ComponentOptionsMap)
	componentPred, err := componentPredicateFactory.GeneratePredicate(componentQuery)
	if err != nil {
		return nil, err
	}

	// Use the images to map CVEs to the images and components.
	cveToResolver := make(map[string]*EmbeddedVulnerabilityResolver)
	for _, image := range images {
		for _, component := range image.GetScan().GetComponents() {
			if _, matches := componentPred(component); !matches {
				continue
			}
			for _, vuln := range component.GetVulns() {
				if !vulnPred.Matches(vuln) {
					continue
				}
				if _, exists := cveToResolver[vuln.GetCve()]; !exists {
					cveToResolver[vuln.GetCve()] = &EmbeddedVulnerabilityResolver{
						data: vuln,
						root: root,
					}
				}
				latestTime := cveToResolver[vuln.GetCve()].lastScanned
				if latestTime == nil || image.GetScan().GetScanTime().Compare(latestTime) > 0 {
					cveToResolver[vuln.GetCve()].lastScanned = image.GetScan().GetScanTime()
				}
			}
		}
	}

	// Create the resolvers.
	resolvers := make([]*EmbeddedVulnerabilityResolver, 0, len(cveToResolver))
	for _, vuln := range cveToResolver {
		resolvers = append(resolvers, vuln)
	}
	return resolvers, nil
}
