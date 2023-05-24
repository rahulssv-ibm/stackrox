package resolvers

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/graphql/resolvers/distroctx"
	"github.com/stackrox/rox/central/graphql/resolvers/loaders"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	pkgMetrics "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/scoped"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	imageWatchStatuses []string

	notWatchedImageWatchStatus = registerImageWatchStatus("NOT_WATCHED")
	watchedImageStatus         = registerImageWatchStatus("WATCHED")
)

func registerImageWatchStatus(s string) string {
	imageWatchStatuses = append(imageWatchStatuses, s)
	return s
}

func init() {
	schema := getBuilder()
	utils.Must(
		// NOTE: This list is and should remain alphabetically ordered
		schema.AddExtraResolvers("Image", []string{
			"deploymentCount(query: String): Int!",
			"deployments(query: String, pagination: Pagination): [Deployment!]!",
			"imageComponentCount(query: String): Int!",
			"imageComponents(query: String, pagination: Pagination): [ImageComponent!]!",
			"imageCVECountBySeverity(query: String): ResourceCountByCVESeverity!",
			"imageVulnerabilityCount(query: String): Int!",
			"imageVulnerabilityCounter(query: String): VulnerabilityCounter!",
			"imageVulnerabilities(query: String, scopeQuery: String, pagination: Pagination): [ImageVulnerability]!",
			"plottedImageVulnerabilities(query: String): PlottedImageVulnerabilities!",
			"scan: ImageScan",
			"topImageVulnerability(query: String): ImageVulnerability",
			"unusedVarSink(query: String): Int",
			"watchStatus: ImageWatchStatus!",

			// Image scan-related fields
			"dataSource: DataSource",
			"scanNotes: [ImageScan_Note!]!",
			"operatingSystem: String!",
			"scanTime: Time",
			"scannerVersion: String!",
		}),
		// deprecated fields
		schema.AddExtraResolvers("Image", []string{
			"topVuln(query: String): EmbeddedVulnerability " +
				"@deprecated(reason: \"use 'topImageVulnerability'\")",
			"vulnCount(query: String): Int! " +
				"@deprecated(reason: \"use 'imageVulnerabilityCount'\")",
			"vulnCounter(query: String): VulnerabilityCounter! " +
				"@deprecated(reason: \"use 'imageVulnerabilityCounter'\")",
			"vulns(query: String, scopeQuery: String, pagination: Pagination): [EmbeddedVulnerability]! " +
				"@deprecated(reason: \"use 'imageVulnerabilities'\")",
			"componentCount(query: String): Int!" +
				"@deprecated(reason: \"use 'imageComponentCount'\")",
			"components(query: String, pagination: Pagination): [EmbeddedImageScanComponent!]!" +
				"@deprecated(reason: \"use 'imageComponentCount'\")",
			"plottedVulns(query: String): PlottedVulnerabilities!" +
				"@deprecated(reason: \"use 'plottedImageVulnerabilities'\")",
		}),
		schema.AddQuery("image(id: ID!): Image"),
		schema.AddQuery("fullImage(id: ID!): Image"),
		schema.AddQuery("images(query: String, pagination: Pagination): [Image!]!"),
		schema.AddQuery("imageCount(query: String): Int!"),
		schema.AddEnumType("ImageWatchStatus", imageWatchStatuses),
	)
}

// Images returns GraphQL resolvers for all images
func (resolver *Resolver) Images(ctx context.Context, args PaginatedQuery) ([]*imageResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "Images")
	if err := readImages(ctx); err != nil {
		return nil, err
	}
	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil, err
	}
	images, err := imageLoader.FromQuery(ctx, q)
	return resolver.wrapImagesWithContext(ctx, images, err)
}

// ImageCount returns count of all images across deployments
func (resolver *Resolver) ImageCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "ImageCount")
	if err := readImages(ctx); err != nil {
		return 0, err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return 0, err
	}
	return imageLoader.CountFromQuery(ctx, q)
}

// Image returns a graphql resolver for the identified image, if it exists
func (resolver *Resolver) Image(ctx context.Context, args struct{ ID graphql.ID }) (*imageResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "Image")
	if err := readImages(ctx); err != nil {
		return nil, err
	}

	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil, err
	}
	image, err := imageLoader.FromID(ctx, string(args.ID))
	return resolver.wrapImageWithContext(ctx, image, image != nil, err)
}

// FullImage returns a graphql resolver for the identified image, if it exists
func (resolver *Resolver) FullImage(ctx context.Context, args struct{ ID graphql.ID }) (*imageResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "FullImage")
	if err := readImages(ctx); err != nil {
		return nil, err
	}

	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil, err
	}
	image, err := imageLoader.FullImageWithID(ctx, string(args.ID))
	return resolver.wrapImageWithContext(ctx, image, image != nil, err)
}

// Deployments returns the deployments which use this image for the identified image, if it exists
func (resolver *imageResolver) Deployments(ctx context.Context, args PaginatedQuery) ([]*deploymentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "Deployments")
	return resolver.root.Deployments(resolver.withImageScopeContext(ctx), args)
}

// DeploymentCount returns the number of deployments which use this image for the identified image, if it exists
func (resolver *imageResolver) DeploymentCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "DeploymentCount")
	return resolver.root.DeploymentCount(resolver.withImageScopeContext(ctx), args)
}

// TopImageVulnerability returns the image vulnerability with the top CVSS score.
func (resolver *imageResolver) TopImageVulnerability(ctx context.Context, args RawQuery) (ImageVulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "TopImageVulnerability")
	return resolver.root.TopImageVulnerability(resolver.withImageScopeContext(ctx), args)
}

// TopVuln returns the first vulnerability with the top CVSS score.
func (resolver *imageResolver) TopVuln(ctx context.Context, args RawQuery) (VulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "TopVuln")
	return nil, errors.New("TopVuln not supported with postgres enabled. Please use TopImageVulnerability.")
}

func (resolver *imageResolver) topVulnV2(ctx context.Context, args RawQuery) (*cVEResolver, error) {
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	query = search.ConjunctionQuery(query, resolver.getImageQuery())
	query.Pagination = &v1.QueryPagination{
		SortOptions: []*v1.QuerySortOption{
			{
				Field:    search.CVSS.String(),
				Reversed: true,
			},
			{
				Field:    search.CVE.String(),
				Reversed: true,
			},
		},
		Limit:  1,
		Offset: 0,
	}

	vulnLoader, err := loaders.GetCVELoader(ctx)
	if err != nil {
		return nil, err
	}
	vulns, err := vulnLoader.FromQuery(ctx, query)
	if err != nil {
		return nil, err
	} else if len(vulns) == 0 {
		return nil, err
	} else if len(vulns) > 1 {
		return nil, errors.New("multiple vulnerabilities matched for top image vulnerability")
	}
	return &cVEResolver{root: resolver.root, data: vulns[0]}, nil
}

func (resolver *imageResolver) vulnQueryScoping(ctx context.Context) context.Context {
	ctx = scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGES,
		ID:    resolver.data.GetId(),
	})
	ctx = distroctx.Context(ctx, resolver.data.GetScan().GetOperatingSystem())

	return ctx
}

// ImageVulnerabilities returns, as ImageVulnerabilityResolver, the vulnerabilities for the image
func (resolver *imageResolver) ImageVulnerabilities(ctx context.Context, args PaginatedQuery) ([]ImageVulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ImageVulnerabilities")
	return resolver.root.ImageVulnerabilities(resolver.withImageScopeContext(ctx), args)
}

func (resolver *imageResolver) ImageVulnerabilityCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ImageVulnerabilityCount")
	return resolver.root.ImageVulnerabilityCount(resolver.withImageScopeContext(ctx), args)
}

func (resolver *imageResolver) ImageVulnerabilityCounter(ctx context.Context, args RawQuery) (*VulnerabilityCounterResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ImageVulnerabilityCounter")
	return resolver.root.ImageVulnerabilityCounter(resolver.withImageScopeContext(ctx), args)
}

func (resolver *imageResolver) ImageCVECountBySeverity(ctx context.Context, q RawQuery) (*resourceCountBySeverityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ImageCVECountBySeverity")

	if !features.VulnMgmtWorkloadCVEs.Enabled() {
		return nil, errors.Errorf("%s=false. Set %s=true and retry", features.VulnMgmtWorkloadCVEs.Name(), features.VulnMgmtWorkloadCVEs.Name())
	}
	if err := readImages(ctx); err != nil {
		return nil, err
	}
	query, err := q.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}
	val, err := resolver.root.ImageCVEView.CountBySeverity(resolver.withImageScopeContext(ctx), query)
	return resolver.root.wrapResourceCountByCVESeverityWithContext(ctx, val, err)
}

// Vulns returns all the vulnerabilities in the image.
func (resolver *imageResolver) Vulns(ctx context.Context, args PaginatedQuery) ([]VulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "Vulns")
	return nil, errors.New("Vulns not supported with postgres enabled. Please use ImageVulnerabilities.")
}

// VulnCount returns the number of vulnerabilities the image has.
func (resolver *imageResolver) VulnCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "VulnCount")
	return 0, errors.New("VulnCount not supported with postgres enabled. Please use ImageVulnerabilityCount.")
}

// VulnCounter resolves the number of different types of vulnerabilities contained in an image component.
func (resolver *imageResolver) VulnCounter(ctx context.Context, args RawQuery) (*VulnerabilityCounterResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "VulnCounter")
	return nil, errors.New("VulnCounter not supported with postgres enabled. Please use ImageVulnerabilityCounter.")
}

func (resolver *imageResolver) ImageComponents(ctx context.Context, args PaginatedQuery) ([]ImageComponentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ImageComponents")
	return resolver.root.ImageComponents(resolver.withImageScopeContext(ctx), args)
}

func (resolver *imageResolver) ImageComponentCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ImageComponentCount")
	return resolver.root.ImageComponentCount(resolver.withImageScopeContext(ctx), args)
}

func (resolver *imageResolver) withImageScopeContext(ctx context.Context) context.Context {
	if ctx == nil {
		err := utils.ShouldErr(errors.New("argument 'ctx' is nil"))
		if err != nil {
			log.Error(err)
		}
	}
	if resolver.ctx == nil {
		resolver.ctx = ctx
	}
	return scoped.Context(resolver.ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGES,
		ID:    resolver.data.GetId(),
	})
}

// Components returns all the components in the image.
func (resolver *imageResolver) Components(ctx context.Context, args PaginatedQuery) ([]ComponentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "Components")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getImageRawQuery())

	ctx = scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGES,
		ID:    resolver.data.GetId(),
	})
	ctx = distroctx.Context(ctx, resolver.data.GetScan().GetOperatingSystem())
	return resolver.root.Components(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *imageResolver) ComponentCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "ComponentCount")

	if args.IsEmpty() {
		return int32(len(resolver.data.GetScan().GetComponents())), nil
	}

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getImageRawQuery())
	return resolver.root.ComponentCount(scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGES,
		ID:    resolver.data.GetId(),
	}), RawQuery{Query: &query})
}

func (resolver *Resolver) getImage(ctx context.Context, id string) *storage.Image {
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil
	}
	image, err := imageLoader.FromID(ctx, id)
	if err != nil {
		return nil
	}
	return image
}

func (resolver *imageResolver) getImageRawQuery() string {
	return search.NewQueryBuilder().AddExactMatches(search.ImageSHA, resolver.data.GetId()).Query()
}

func (resolver *imageResolver) getImageQuery() *v1.Query {
	return search.NewQueryBuilder().AddExactMatches(search.ImageSHA, resolver.data.GetId()).ProtoQuery()
}

func (resolver *imageResolver) PlottedVulns(ctx context.Context, args RawQuery) (*PlottedVulnerabilitiesResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "PlottedVulns")
	return nil, errors.New("PlottedVulns resolver is not support on postgres. Use PlottedImageVulnerabilities.")
}

// PlottedImageVulnerabilities returns the data required by top risky entity scatter-plot on vuln mgmt dashboard
func (resolver *imageResolver) PlottedImageVulnerabilities(ctx context.Context, args RawQuery) (*PlottedImageVulnerabilitiesResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "PlottedImageVulnerabilities")
	return resolver.root.PlottedImageVulnerabilities(resolver.withImageScopeContext(ctx), args)
}

func (resolver *imageResolver) Scan(ctx context.Context) (*imageScanResolver, error) {
	resolver.ensureData(ctx)
	scan := resolver.data.GetScan()
	// If scan is pulled, it is most likely to fetch all components and vulns contained in image.
	// Therefore, load the image again with full scan.
	imageLoader, err := loaders.GetImageLoader(ctx)
	if err != nil {
		return nil, err
	}

	// If Postgres is not enabled, image loader always pull full image.
	image, err := imageLoader.FullImageWithID(ctx, resolver.data.GetId())
	if err != nil {
		return nil, err
	}
	scan = image.GetScan()

	res, err := resolver.root.wrapImageScan(scan, true, nil)
	if err != nil || res == nil {
		return nil, err
	}
	res.ctx = resolver.withImageScopeContext(ctx)
	return res, nil
}

func (resolver *imageResolver) WatchStatus(ctx context.Context) (string, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Images, "WatchStatus")
	if err := readAuth(resources.WatchedImage)(ctx); err != nil {
		return "", err
	}
	watched, err := resolver.root.WatchedImageDataStore.Exists(ctx, resolver.data.GetName().GetFullName())
	if err != nil {
		return "", err
	}
	if watched {
		return watchedImageStatus, nil
	}
	return notWatchedImageWatchStatus, nil
}

func (resolver *imageResolver) UnusedVarSink(_ context.Context, _ RawQuery) *int32 {
	return nil
}

//// Image scan-related fields pulled as direct sub-resolvers of image.

func (resolver *imageResolver) DataSource(_ context.Context) (*dataSourceResolver, error) {
	value := resolver.data.GetScan().GetDataSource()
	return resolver.root.wrapDataSource(value, true, nil)
}

func (resolver *imageResolver) ScanNotes(_ context.Context) []string {
	value := resolver.data.GetScan().GetNotes()
	return stringSlice(value)
}

func (resolver *imageResolver) OperatingSystem(_ context.Context) string {
	value := resolver.data.GetScan().GetOperatingSystem()
	return value
}

func (resolver *imageResolver) ScanTime(_ context.Context) (*graphql.Time, error) {
	value := resolver.data.GetScan().GetScanTime()
	return timestamp(value)
}

func (resolver *imageResolver) ScannerVersion(_ context.Context) string {
	value := resolver.data.GetScan().GetScannerVersion()
	return value
}
