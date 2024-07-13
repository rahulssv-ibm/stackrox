// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/process_listening_on_port_service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_ListeningEndpointsService_GetListeningEndpoints_0(ctx context.Context, marshaler runtime.Marshaler, client ListeningEndpointsServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProcessesListeningOnPortsRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["deployment_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id")
	}

	protoReq.DeploymentId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id", err)
	}

	msg, err := client.GetListeningEndpoints(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ListeningEndpointsService_GetListeningEndpoints_0(ctx context.Context, marshaler runtime.Marshaler, server ListeningEndpointsServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProcessesListeningOnPortsRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["deployment_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id")
	}

	protoReq.DeploymentId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id", err)
	}

	msg, err := server.GetListeningEndpoints(ctx, &protoReq)
	return msg, metadata, err

}

func request_ListeningEndpointsService_CountListeningEndpoints_0(ctx context.Context, marshaler runtime.Marshaler, client ListeningEndpointsServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.CountListeningEndpoints(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ListeningEndpointsService_CountListeningEndpoints_0(ctx context.Context, marshaler runtime.Marshaler, server ListeningEndpointsServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := server.CountListeningEndpoints(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterListeningEndpointsServiceHandlerServer registers the http handlers for service ListeningEndpointsService to "mux".
// UnaryRPC     :call ListeningEndpointsServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterListeningEndpointsServiceHandlerFromEndpoint instead.
func RegisterListeningEndpointsServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ListeningEndpointsServiceServer) error {

	mux.Handle("GET", pattern_ListeningEndpointsService_GetListeningEndpoints_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/v1.ListeningEndpointsService/GetListeningEndpoints", runtime.WithHTTPPathPattern("/v1/listening_endpoints/deployment/{deployment_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ListeningEndpointsService_GetListeningEndpoints_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ListeningEndpointsService_GetListeningEndpoints_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ListeningEndpointsService_CountListeningEndpoints_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/v1.ListeningEndpointsService/CountListeningEndpoints", runtime.WithHTTPPathPattern("/v1/listening_endpoints/count"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ListeningEndpointsService_CountListeningEndpoints_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ListeningEndpointsService_CountListeningEndpoints_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterListeningEndpointsServiceHandlerFromEndpoint is same as RegisterListeningEndpointsServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterListeningEndpointsServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterListeningEndpointsServiceHandler(ctx, mux, conn)
}

// RegisterListeningEndpointsServiceHandler registers the http handlers for service ListeningEndpointsService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterListeningEndpointsServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterListeningEndpointsServiceHandlerClient(ctx, mux, NewListeningEndpointsServiceClient(conn))
}

// RegisterListeningEndpointsServiceHandlerClient registers the http handlers for service ListeningEndpointsService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ListeningEndpointsServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ListeningEndpointsServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ListeningEndpointsServiceClient" to call the correct interceptors.
func RegisterListeningEndpointsServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ListeningEndpointsServiceClient) error {

	mux.Handle("GET", pattern_ListeningEndpointsService_GetListeningEndpoints_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/v1.ListeningEndpointsService/GetListeningEndpoints", runtime.WithHTTPPathPattern("/v1/listening_endpoints/deployment/{deployment_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ListeningEndpointsService_GetListeningEndpoints_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ListeningEndpointsService_GetListeningEndpoints_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ListeningEndpointsService_CountListeningEndpoints_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/v1.ListeningEndpointsService/CountListeningEndpoints", runtime.WithHTTPPathPattern("/v1/listening_endpoints/count"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ListeningEndpointsService_CountListeningEndpoints_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ListeningEndpointsService_CountListeningEndpoints_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ListeningEndpointsService_GetListeningEndpoints_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "listening_endpoints", "deployment", "deployment_id"}, ""))

	pattern_ListeningEndpointsService_CountListeningEndpoints_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "listening_endpoints", "count"}, ""))
)

var (
	forward_ListeningEndpointsService_GetListeningEndpoints_0 = runtime.ForwardResponseMessage

	forward_ListeningEndpointsService_CountListeningEndpoints_0 = runtime.ForwardResponseMessage
)
