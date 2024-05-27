package service

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	profileDS "github.com/stackrox/rox/central/complianceoperator/v2/profiles/datastore"
	"github.com/stackrox/rox/central/convert/storagetov2"
	v1 "github.com/stackrox/rox/generated/api/v1"
	v2 "github.com/stackrox/rox/generated/api/v2"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/grpc/authz"
	"github.com/stackrox/rox/pkg/grpc/authz/perrpc"
	"github.com/stackrox/rox/pkg/grpc/authz/user"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/paginated"
	"google.golang.org/grpc"
)

const (
	maxPaginationLimit = 1000
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.Compliance)): {
			"/v2.ComplianceProfileService/GetComplianceProfile",   // TODO: Lists Standard in Proto, necessary here?
			"/v2.ComplianceProfileService/ListComplianceProfiles", // TODO: Lists Standard in Proto, necessary here?
			"/v2.ComplianceProfileService/ListProfileSummaries",   // does not include standard/control
			"/v2.ComplianceProfileService/GetComplianceProfileCount",
		},
	})
)

// New returns a service object for registering with grpc.
func New(complianceProfilesDS profileDS.DataStore) Service {
	return &serviceImpl{
		complianceProfilesDS: complianceProfilesDS,
	}
}

type serviceImpl struct {
	v2.UnimplementedComplianceProfileServiceServer

	complianceProfilesDS profileDS.DataStore
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v2.RegisterComplianceProfileServiceServer(grpcServer, s)
}

// RegisterServiceHandler registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return v2.RegisterComplianceProfileServiceHandler(ctx, mux, conn)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, authorizer.Authorized(ctx, fullMethodName)
}

// GetComplianceProfile retrieves the specified compliance profile
func (s *serviceImpl) GetComplianceProfile(ctx context.Context, req *v2.ResourceByID) (*v2.ComplianceProfile, error) {
	if req.GetId() == "" {
		return nil, errors.Wrap(errox.InvalidArgs, "Scan configuration name is required for retrieval")
	}

	profile, found, err := s.complianceProfilesDS.GetProfile(ctx, req.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve compliance profile with id %q.", req.GetId())
	}
	if !found {
		return nil, errors.Wrapf(errox.NotFound, "compliance profile with id %q does not exist", req.GetId())
	}

	return storagetov2.ComplianceV2Profile(profile), nil
}

// ListComplianceProfiles returns profiles matching given query
func (s *serviceImpl) ListComplianceProfiles(ctx context.Context, request *v2.ProfilesForClusterRequest) (*v2.ListComplianceProfilesResponse, error) {
	if request.GetClusterId() == "" {
		return nil, errors.Wrap(errox.InvalidArgs, "cluster is required")
	}

	// Fill in Query.
	parsedQuery, err := search.ParseQuery(request.GetQuery().GetQuery(), search.MatchAllIfEmpty())
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to parse query %v", err)
	}

	// Add the cluster ids as an exact match
	parsedQuery = search.ConjunctionQuery(
		search.NewQueryBuilder().AddSelectFields().AddExactMatches(search.ClusterID, request.GetClusterId()).ProtoQuery(),
		parsedQuery,
	)

	// To get total count, need the parsed query without the paging.
	countQuery := parsedQuery.Clone()

	// Fill in pagination.
	paginated.FillPaginationV2(parsedQuery, request.GetQuery().GetPagination(), maxPaginationLimit)

	profiles, err := s.complianceProfilesDS.SearchProfiles(ctx, parsedQuery)
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance profiles for cluster %v", request.GetClusterId())
	}

	totalCount, err := s.complianceProfilesDS.CountProfiles(ctx, countQuery)
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance profiles counts for %v", request)
	}

	return &v2.ListComplianceProfilesResponse{
		Profiles:   storagetov2.ComplianceV2Profiles(profiles),
		TotalCount: int32(totalCount),
	}, nil
}

// ListProfileSummaries returns profile summaries matching incoming clusters
func (s *serviceImpl) ListProfileSummaries(ctx context.Context, request *v2.ClustersProfileSummaryRequest) (*v2.ListComplianceProfileSummaryResponse, error) {
	if len(request.GetClusterIds()) == 0 {
		return nil, errors.Wrap(errox.InvalidArgs, "cluster is required")
	}

	// Fill in Query.
	parsedQuery, err := search.ParseQuery(request.GetQuery().GetQuery(), search.MatchAllIfEmpty())
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to parse query %v", err)
	}

	// To get total count, need the parsed query without the paging.
	countQuery := parsedQuery.Clone()

	// Fill in pagination.
	paginated.FillPaginationV2(parsedQuery, request.GetQuery().GetPagination(), maxPaginationLimit)
	// make sure we sort by profile name at a minimum
	if parsedQuery.GetPagination().GetSortOptions() == nil {
		parsedQuery.Pagination.SortOptions = []*v1.QuerySortOption{
			{
				Field: search.ComplianceOperatorProfileName.String(),
			},
		}
	}

	profileNames, err := s.complianceProfilesDS.GetProfilesNames(ctx, parsedQuery, request.GetClusterIds())
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance profiles for %v", request)
	}

	// Build query to get the filtered list by profile names
	profileQuery := search.NewQueryBuilder().AddSelectFields().AddExactMatches(search.ComplianceOperatorProfileName, profileNames...).ProtoQuery()
	// Bring the sort options only, paging is handled in step one when we get the distinct profiles.
	profileQuery.Pagination = &v1.QueryPagination{}
	profileQuery.Pagination.SortOptions = parsedQuery.GetPagination().GetSortOptions()

	profiles, err := s.complianceProfilesDS.SearchProfiles(ctx, profileQuery)
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance profiles for %v", request)
	}

	totalCount, err := s.complianceProfilesDS.CountDistinctProfiles(ctx, countQuery, request.GetClusterIds())
	if err != nil {
		return nil, errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance profiles counts for %v", request)
	}

	return &v2.ListComplianceProfileSummaryResponse{
		Profiles:   storagetov2.ComplianceProfileSummary(profiles, request.GetClusterIds()),
		TotalCount: int32(totalCount),
	}, nil
}

// GetComplianceProfileCount returns counts of profiles matching query
func (s *serviceImpl) GetComplianceProfileCount(ctx context.Context, request *v2.RawQuery) (*v2.CountComplianceProfilesResponse, error) {
	parsedQuery, err := search.ParseQuery(request.GetQuery(), search.MatchAllIfEmpty())
	if err != nil {
		return nil, errors.Wrap(errox.InvalidArgs, err.Error())
	}

	profileCount, err := s.complianceProfilesDS.CountProfiles(ctx, parsedQuery)
	if err != nil {
		return nil, errors.Wrap(errox.NotFound, err.Error())
	}
	return &v2.CountComplianceProfilesResponse{
		Count: int32(profileCount),
	}, nil
}
