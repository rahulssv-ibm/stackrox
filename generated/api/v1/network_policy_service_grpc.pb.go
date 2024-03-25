// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/network_policy_service.proto

package v1

import (
	context "context"
	storage "github.com/stackrox/rox/generated/storage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NetworkPolicyService_GetNetworkPolicy_FullMethodName                                  = "/v1.NetworkPolicyService/GetNetworkPolicy"
	NetworkPolicyService_GetNetworkPolicies_FullMethodName                                = "/v1.NetworkPolicyService/GetNetworkPolicies"
	NetworkPolicyService_GetNetworkGraph_FullMethodName                                   = "/v1.NetworkPolicyService/GetNetworkGraph"
	NetworkPolicyService_GetNetworkGraphEpoch_FullMethodName                              = "/v1.NetworkPolicyService/GetNetworkGraphEpoch"
	NetworkPolicyService_ApplyNetworkPolicy_FullMethodName                                = "/v1.NetworkPolicyService/ApplyNetworkPolicy"
	NetworkPolicyService_GetUndoModification_FullMethodName                               = "/v1.NetworkPolicyService/GetUndoModification"
	NetworkPolicyService_SimulateNetworkGraph_FullMethodName                              = "/v1.NetworkPolicyService/SimulateNetworkGraph"
	NetworkPolicyService_SendNetworkPolicyYAML_FullMethodName                             = "/v1.NetworkPolicyService/SendNetworkPolicyYAML"
	NetworkPolicyService_GenerateNetworkPolicies_FullMethodName                           = "/v1.NetworkPolicyService/GenerateNetworkPolicies"
	NetworkPolicyService_GetBaselineGeneratedNetworkPolicyForDeployment_FullMethodName    = "/v1.NetworkPolicyService/GetBaselineGeneratedNetworkPolicyForDeployment"
	NetworkPolicyService_GetAllowedPeersFromCurrentPolicyForDeployment_FullMethodName     = "/v1.NetworkPolicyService/GetAllowedPeersFromCurrentPolicyForDeployment"
	NetworkPolicyService_ApplyNetworkPolicyYamlForDeployment_FullMethodName               = "/v1.NetworkPolicyService/ApplyNetworkPolicyYamlForDeployment"
	NetworkPolicyService_GetUndoModificationForDeployment_FullMethodName                  = "/v1.NetworkPolicyService/GetUndoModificationForDeployment"
	NetworkPolicyService_GetDiffFlowsBetweenPolicyAndBaselineForDeployment_FullMethodName = "/v1.NetworkPolicyService/GetDiffFlowsBetweenPolicyAndBaselineForDeployment"
	NetworkPolicyService_GetDiffFlowsFromUndoModificationForDeployment_FullMethodName     = "/v1.NetworkPolicyService/GetDiffFlowsFromUndoModificationForDeployment"
)

// NetworkPolicyServiceClient is the client API for NetworkPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkPolicyServiceClient interface {
	GetNetworkPolicy(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.NetworkPolicy, error)
	GetNetworkPolicies(ctx context.Context, in *GetNetworkPoliciesRequest, opts ...grpc.CallOption) (*NetworkPoliciesResponse, error)
	GetNetworkGraph(ctx context.Context, in *GetNetworkGraphRequest, opts ...grpc.CallOption) (*NetworkGraph, error)
	GetNetworkGraphEpoch(ctx context.Context, in *GetNetworkGraphEpochRequest, opts ...grpc.CallOption) (*NetworkGraphEpoch, error)
	ApplyNetworkPolicy(ctx context.Context, in *ApplyNetworkPolicyYamlRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUndoModification(ctx context.Context, in *GetUndoModificationRequest, opts ...grpc.CallOption) (*GetUndoModificationResponse, error)
	SimulateNetworkGraph(ctx context.Context, in *SimulateNetworkGraphRequest, opts ...grpc.CallOption) (*SimulateNetworkGraphResponse, error)
	SendNetworkPolicyYAML(ctx context.Context, in *SendNetworkPolicyYamlRequest, opts ...grpc.CallOption) (*Empty, error)
	GenerateNetworkPolicies(ctx context.Context, in *GenerateNetworkPoliciesRequest, opts ...grpc.CallOption) (*GenerateNetworkPoliciesResponse, error)
	GetBaselineGeneratedNetworkPolicyForDeployment(ctx context.Context, in *GetBaselineGeneratedPolicyForDeploymentRequest, opts ...grpc.CallOption) (*GetBaselineGeneratedPolicyForDeploymentResponse, error)
	GetAllowedPeersFromCurrentPolicyForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetAllowedPeersFromCurrentPolicyForDeploymentResponse, error)
	ApplyNetworkPolicyYamlForDeployment(ctx context.Context, in *ApplyNetworkPolicyYamlForDeploymentRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUndoModificationForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetUndoModificationForDeploymentResponse, error)
	GetDiffFlowsBetweenPolicyAndBaselineForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetDiffFlowsResponse, error)
	GetDiffFlowsFromUndoModificationForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetDiffFlowsResponse, error)
}

type networkPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkPolicyServiceClient(cc grpc.ClientConnInterface) NetworkPolicyServiceClient {
	return &networkPolicyServiceClient{cc}
}

func (c *networkPolicyServiceClient) GetNetworkPolicy(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*storage.NetworkPolicy, error) {
	out := new(storage.NetworkPolicy)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetNetworkPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetNetworkPolicies(ctx context.Context, in *GetNetworkPoliciesRequest, opts ...grpc.CallOption) (*NetworkPoliciesResponse, error) {
	out := new(NetworkPoliciesResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetNetworkPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetNetworkGraph(ctx context.Context, in *GetNetworkGraphRequest, opts ...grpc.CallOption) (*NetworkGraph, error) {
	out := new(NetworkGraph)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetNetworkGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetNetworkGraphEpoch(ctx context.Context, in *GetNetworkGraphEpochRequest, opts ...grpc.CallOption) (*NetworkGraphEpoch, error) {
	out := new(NetworkGraphEpoch)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetNetworkGraphEpoch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) ApplyNetworkPolicy(ctx context.Context, in *ApplyNetworkPolicyYamlRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NetworkPolicyService_ApplyNetworkPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetUndoModification(ctx context.Context, in *GetUndoModificationRequest, opts ...grpc.CallOption) (*GetUndoModificationResponse, error) {
	out := new(GetUndoModificationResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetUndoModification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) SimulateNetworkGraph(ctx context.Context, in *SimulateNetworkGraphRequest, opts ...grpc.CallOption) (*SimulateNetworkGraphResponse, error) {
	out := new(SimulateNetworkGraphResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_SimulateNetworkGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) SendNetworkPolicyYAML(ctx context.Context, in *SendNetworkPolicyYamlRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NetworkPolicyService_SendNetworkPolicyYAML_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GenerateNetworkPolicies(ctx context.Context, in *GenerateNetworkPoliciesRequest, opts ...grpc.CallOption) (*GenerateNetworkPoliciesResponse, error) {
	out := new(GenerateNetworkPoliciesResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GenerateNetworkPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetBaselineGeneratedNetworkPolicyForDeployment(ctx context.Context, in *GetBaselineGeneratedPolicyForDeploymentRequest, opts ...grpc.CallOption) (*GetBaselineGeneratedPolicyForDeploymentResponse, error) {
	out := new(GetBaselineGeneratedPolicyForDeploymentResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetBaselineGeneratedNetworkPolicyForDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetAllowedPeersFromCurrentPolicyForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetAllowedPeersFromCurrentPolicyForDeploymentResponse, error) {
	out := new(GetAllowedPeersFromCurrentPolicyForDeploymentResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetAllowedPeersFromCurrentPolicyForDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) ApplyNetworkPolicyYamlForDeployment(ctx context.Context, in *ApplyNetworkPolicyYamlForDeploymentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NetworkPolicyService_ApplyNetworkPolicyYamlForDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetUndoModificationForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetUndoModificationForDeploymentResponse, error) {
	out := new(GetUndoModificationForDeploymentResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetUndoModificationForDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetDiffFlowsBetweenPolicyAndBaselineForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetDiffFlowsResponse, error) {
	out := new(GetDiffFlowsResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetDiffFlowsBetweenPolicyAndBaselineForDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkPolicyServiceClient) GetDiffFlowsFromUndoModificationForDeployment(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetDiffFlowsResponse, error) {
	out := new(GetDiffFlowsResponse)
	err := c.cc.Invoke(ctx, NetworkPolicyService_GetDiffFlowsFromUndoModificationForDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkPolicyServiceServer is the server API for NetworkPolicyService service.
// All implementations must embed UnimplementedNetworkPolicyServiceServer
// for forward compatibility
type NetworkPolicyServiceServer interface {
	GetNetworkPolicy(context.Context, *ResourceByID) (*storage.NetworkPolicy, error)
	GetNetworkPolicies(context.Context, *GetNetworkPoliciesRequest) (*NetworkPoliciesResponse, error)
	GetNetworkGraph(context.Context, *GetNetworkGraphRequest) (*NetworkGraph, error)
	GetNetworkGraphEpoch(context.Context, *GetNetworkGraphEpochRequest) (*NetworkGraphEpoch, error)
	ApplyNetworkPolicy(context.Context, *ApplyNetworkPolicyYamlRequest) (*Empty, error)
	GetUndoModification(context.Context, *GetUndoModificationRequest) (*GetUndoModificationResponse, error)
	SimulateNetworkGraph(context.Context, *SimulateNetworkGraphRequest) (*SimulateNetworkGraphResponse, error)
	SendNetworkPolicyYAML(context.Context, *SendNetworkPolicyYamlRequest) (*Empty, error)
	GenerateNetworkPolicies(context.Context, *GenerateNetworkPoliciesRequest) (*GenerateNetworkPoliciesResponse, error)
	GetBaselineGeneratedNetworkPolicyForDeployment(context.Context, *GetBaselineGeneratedPolicyForDeploymentRequest) (*GetBaselineGeneratedPolicyForDeploymentResponse, error)
	GetAllowedPeersFromCurrentPolicyForDeployment(context.Context, *ResourceByID) (*GetAllowedPeersFromCurrentPolicyForDeploymentResponse, error)
	ApplyNetworkPolicyYamlForDeployment(context.Context, *ApplyNetworkPolicyYamlForDeploymentRequest) (*Empty, error)
	GetUndoModificationForDeployment(context.Context, *ResourceByID) (*GetUndoModificationForDeploymentResponse, error)
	GetDiffFlowsBetweenPolicyAndBaselineForDeployment(context.Context, *ResourceByID) (*GetDiffFlowsResponse, error)
	GetDiffFlowsFromUndoModificationForDeployment(context.Context, *ResourceByID) (*GetDiffFlowsResponse, error)
	mustEmbedUnimplementedNetworkPolicyServiceServer()
}

// UnimplementedNetworkPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkPolicyServiceServer struct {
}

func (UnimplementedNetworkPolicyServiceServer) GetNetworkPolicy(context.Context, *ResourceByID) (*storage.NetworkPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkPolicy not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetNetworkPolicies(context.Context, *GetNetworkPoliciesRequest) (*NetworkPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkPolicies not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetNetworkGraph(context.Context, *GetNetworkGraphRequest) (*NetworkGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkGraph not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetNetworkGraphEpoch(context.Context, *GetNetworkGraphEpochRequest) (*NetworkGraphEpoch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkGraphEpoch not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) ApplyNetworkPolicy(context.Context, *ApplyNetworkPolicyYamlRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyNetworkPolicy not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetUndoModification(context.Context, *GetUndoModificationRequest) (*GetUndoModificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUndoModification not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) SimulateNetworkGraph(context.Context, *SimulateNetworkGraphRequest) (*SimulateNetworkGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimulateNetworkGraph not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) SendNetworkPolicyYAML(context.Context, *SendNetworkPolicyYamlRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNetworkPolicyYAML not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GenerateNetworkPolicies(context.Context, *GenerateNetworkPoliciesRequest) (*GenerateNetworkPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateNetworkPolicies not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetBaselineGeneratedNetworkPolicyForDeployment(context.Context, *GetBaselineGeneratedPolicyForDeploymentRequest) (*GetBaselineGeneratedPolicyForDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaselineGeneratedNetworkPolicyForDeployment not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetAllowedPeersFromCurrentPolicyForDeployment(context.Context, *ResourceByID) (*GetAllowedPeersFromCurrentPolicyForDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllowedPeersFromCurrentPolicyForDeployment not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) ApplyNetworkPolicyYamlForDeployment(context.Context, *ApplyNetworkPolicyYamlForDeploymentRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyNetworkPolicyYamlForDeployment not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetUndoModificationForDeployment(context.Context, *ResourceByID) (*GetUndoModificationForDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUndoModificationForDeployment not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetDiffFlowsBetweenPolicyAndBaselineForDeployment(context.Context, *ResourceByID) (*GetDiffFlowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiffFlowsBetweenPolicyAndBaselineForDeployment not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) GetDiffFlowsFromUndoModificationForDeployment(context.Context, *ResourceByID) (*GetDiffFlowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiffFlowsFromUndoModificationForDeployment not implemented")
}
func (UnimplementedNetworkPolicyServiceServer) mustEmbedUnimplementedNetworkPolicyServiceServer() {}

// UnsafeNetworkPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkPolicyServiceServer will
// result in compilation errors.
type UnsafeNetworkPolicyServiceServer interface {
	mustEmbedUnimplementedNetworkPolicyServiceServer()
}

func RegisterNetworkPolicyServiceServer(s grpc.ServiceRegistrar, srv NetworkPolicyServiceServer) {
	s.RegisterService(&NetworkPolicyService_ServiceDesc, srv)
}

func _NetworkPolicyService_GetNetworkPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetNetworkPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetNetworkPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetNetworkPolicy(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetNetworkPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetNetworkPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetNetworkPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetNetworkPolicies(ctx, req.(*GetNetworkPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetNetworkGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetNetworkGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetNetworkGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetNetworkGraph(ctx, req.(*GetNetworkGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetNetworkGraphEpoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkGraphEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetNetworkGraphEpoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetNetworkGraphEpoch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetNetworkGraphEpoch(ctx, req.(*GetNetworkGraphEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_ApplyNetworkPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNetworkPolicyYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).ApplyNetworkPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_ApplyNetworkPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).ApplyNetworkPolicy(ctx, req.(*ApplyNetworkPolicyYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetUndoModification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUndoModificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetUndoModification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetUndoModification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetUndoModification(ctx, req.(*GetUndoModificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_SimulateNetworkGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateNetworkGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).SimulateNetworkGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_SimulateNetworkGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).SimulateNetworkGraph(ctx, req.(*SimulateNetworkGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_SendNetworkPolicyYAML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNetworkPolicyYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).SendNetworkPolicyYAML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_SendNetworkPolicyYAML_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).SendNetworkPolicyYAML(ctx, req.(*SendNetworkPolicyYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GenerateNetworkPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateNetworkPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GenerateNetworkPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GenerateNetworkPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GenerateNetworkPolicies(ctx, req.(*GenerateNetworkPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetBaselineGeneratedNetworkPolicyForDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBaselineGeneratedPolicyForDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetBaselineGeneratedNetworkPolicyForDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetBaselineGeneratedNetworkPolicyForDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetBaselineGeneratedNetworkPolicyForDeployment(ctx, req.(*GetBaselineGeneratedPolicyForDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetAllowedPeersFromCurrentPolicyForDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetAllowedPeersFromCurrentPolicyForDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetAllowedPeersFromCurrentPolicyForDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetAllowedPeersFromCurrentPolicyForDeployment(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_ApplyNetworkPolicyYamlForDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNetworkPolicyYamlForDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).ApplyNetworkPolicyYamlForDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_ApplyNetworkPolicyYamlForDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).ApplyNetworkPolicyYamlForDeployment(ctx, req.(*ApplyNetworkPolicyYamlForDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetUndoModificationForDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetUndoModificationForDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetUndoModificationForDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetUndoModificationForDeployment(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetDiffFlowsBetweenPolicyAndBaselineForDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetDiffFlowsBetweenPolicyAndBaselineForDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetDiffFlowsBetweenPolicyAndBaselineForDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetDiffFlowsBetweenPolicyAndBaselineForDeployment(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkPolicyService_GetDiffFlowsFromUndoModificationForDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyServiceServer).GetDiffFlowsFromUndoModificationForDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkPolicyService_GetDiffFlowsFromUndoModificationForDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyServiceServer).GetDiffFlowsFromUndoModificationForDeployment(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkPolicyService_ServiceDesc is the grpc.ServiceDesc for NetworkPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.NetworkPolicyService",
	HandlerType: (*NetworkPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkPolicy",
			Handler:    _NetworkPolicyService_GetNetworkPolicy_Handler,
		},
		{
			MethodName: "GetNetworkPolicies",
			Handler:    _NetworkPolicyService_GetNetworkPolicies_Handler,
		},
		{
			MethodName: "GetNetworkGraph",
			Handler:    _NetworkPolicyService_GetNetworkGraph_Handler,
		},
		{
			MethodName: "GetNetworkGraphEpoch",
			Handler:    _NetworkPolicyService_GetNetworkGraphEpoch_Handler,
		},
		{
			MethodName: "ApplyNetworkPolicy",
			Handler:    _NetworkPolicyService_ApplyNetworkPolicy_Handler,
		},
		{
			MethodName: "GetUndoModification",
			Handler:    _NetworkPolicyService_GetUndoModification_Handler,
		},
		{
			MethodName: "SimulateNetworkGraph",
			Handler:    _NetworkPolicyService_SimulateNetworkGraph_Handler,
		},
		{
			MethodName: "SendNetworkPolicyYAML",
			Handler:    _NetworkPolicyService_SendNetworkPolicyYAML_Handler,
		},
		{
			MethodName: "GenerateNetworkPolicies",
			Handler:    _NetworkPolicyService_GenerateNetworkPolicies_Handler,
		},
		{
			MethodName: "GetBaselineGeneratedNetworkPolicyForDeployment",
			Handler:    _NetworkPolicyService_GetBaselineGeneratedNetworkPolicyForDeployment_Handler,
		},
		{
			MethodName: "GetAllowedPeersFromCurrentPolicyForDeployment",
			Handler:    _NetworkPolicyService_GetAllowedPeersFromCurrentPolicyForDeployment_Handler,
		},
		{
			MethodName: "ApplyNetworkPolicyYamlForDeployment",
			Handler:    _NetworkPolicyService_ApplyNetworkPolicyYamlForDeployment_Handler,
		},
		{
			MethodName: "GetUndoModificationForDeployment",
			Handler:    _NetworkPolicyService_GetUndoModificationForDeployment_Handler,
		},
		{
			MethodName: "GetDiffFlowsBetweenPolicyAndBaselineForDeployment",
			Handler:    _NetworkPolicyService_GetDiffFlowsBetweenPolicyAndBaselineForDeployment_Handler,
		},
		{
			MethodName: "GetDiffFlowsFromUndoModificationForDeployment",
			Handler:    _NetworkPolicyService_GetDiffFlowsFromUndoModificationForDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/network_policy_service.proto",
}
