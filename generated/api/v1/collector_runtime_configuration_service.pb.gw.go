// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/collector_runtime_configuration_service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0(ctx context.Context, marshaler runtime.Marshaler, client CollectorRuntimeConfigurationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetCollectorRuntimeConfiguration(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0(ctx context.Context, marshaler runtime.Marshaler, server CollectorRuntimeConfigurationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetCollectorRuntimeConfiguration(ctx, &protoReq)
	return msg, metadata, err

}

func request_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0(ctx context.Context, marshaler runtime.Marshaler, client CollectorRuntimeConfigurationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PostCollectorRuntimeConfiguration(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0(ctx context.Context, marshaler runtime.Marshaler, server CollectorRuntimeConfigurationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.PostCollectorRuntimeConfiguration(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCollectorRuntimeConfigurationServiceHandlerServer registers the http handlers for service CollectorRuntimeConfigurationService to "mux".
// UnaryRPC     :call CollectorRuntimeConfigurationServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCollectorRuntimeConfigurationServiceHandlerFromEndpoint instead.
func RegisterCollectorRuntimeConfigurationServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CollectorRuntimeConfigurationServiceServer) error {

	mux.Handle("GET", pattern_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCollectorRuntimeConfigurationServiceHandlerFromEndpoint is same as RegisterCollectorRuntimeConfigurationServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCollectorRuntimeConfigurationServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCollectorRuntimeConfigurationServiceHandler(ctx, mux, conn)
}

// RegisterCollectorRuntimeConfigurationServiceHandler registers the http handlers for service CollectorRuntimeConfigurationService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCollectorRuntimeConfigurationServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCollectorRuntimeConfigurationServiceHandlerClient(ctx, mux, NewCollectorRuntimeConfigurationServiceClient(conn))
}

// RegisterCollectorRuntimeConfigurationServiceHandlerClient registers the http handlers for service CollectorRuntimeConfigurationService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CollectorRuntimeConfigurationServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CollectorRuntimeConfigurationServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CollectorRuntimeConfigurationServiceClient" to call the correct interceptors.
func RegisterCollectorRuntimeConfigurationServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CollectorRuntimeConfigurationServiceClient) error {

	mux.Handle("GET", pattern_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "collector_runtime_configuration"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "collector_runtime_configuration"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_CollectorRuntimeConfigurationService_GetCollectorRuntimeConfiguration_0 = runtime.ForwardResponseMessage

	forward_CollectorRuntimeConfigurationService_PostCollectorRuntimeConfiguration_0 = runtime.ForwardResponseMessage
)
