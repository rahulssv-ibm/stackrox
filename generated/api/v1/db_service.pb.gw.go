// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/db_service.proto

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

func request_DBService_GetExportCapabilities_0(ctx context.Context, marshaler runtime.Marshaler, client DBServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetExportCapabilities(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DBService_GetExportCapabilities_0(ctx context.Context, marshaler runtime.Marshaler, server DBServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetExportCapabilities(ctx, &protoReq)
	return msg, metadata, err

}

func request_DBService_GetActiveRestoreProcess_0(ctx context.Context, marshaler runtime.Marshaler, client DBServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetActiveRestoreProcess(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DBService_GetActiveRestoreProcess_0(ctx context.Context, marshaler runtime.Marshaler, server DBServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetActiveRestoreProcess(ctx, &protoReq)
	return msg, metadata, err

}

func request_DBService_InterruptRestoreProcess_0(ctx context.Context, marshaler runtime.Marshaler, client DBServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InterruptDBRestoreProcessRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["process_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "process_id")
	}

	protoReq.ProcessId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "process_id", err)
	}

	val, ok = pathParams["attempt_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "attempt_id")
	}

	protoReq.AttemptId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "attempt_id", err)
	}

	msg, err := client.InterruptRestoreProcess(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DBService_InterruptRestoreProcess_0(ctx context.Context, marshaler runtime.Marshaler, server DBServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InterruptDBRestoreProcessRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["process_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "process_id")
	}

	protoReq.ProcessId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "process_id", err)
	}

	val, ok = pathParams["attempt_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "attempt_id")
	}

	protoReq.AttemptId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "attempt_id", err)
	}

	msg, err := server.InterruptRestoreProcess(ctx, &protoReq)
	return msg, metadata, err

}

func request_DBService_CancelRestoreProcess_0(ctx context.Context, marshaler runtime.Marshaler, client DBServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ResourceByID
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.CancelRestoreProcess(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DBService_CancelRestoreProcess_0(ctx context.Context, marshaler runtime.Marshaler, server DBServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ResourceByID
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := server.CancelRestoreProcess(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterDBServiceHandlerServer registers the http handlers for service DBService to "mux".
// UnaryRPC     :call DBServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterDBServiceHandlerFromEndpoint instead.
func RegisterDBServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server DBServiceServer) error {

	mux.Handle("GET", pattern_DBService_GetExportCapabilities_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := local_request_DBService_GetExportCapabilities_0(rctx, inboundMarshalVTer, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_GetExportCapabilities_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_DBService_GetActiveRestoreProcess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := local_request_DBService_GetActiveRestoreProcess_0(rctx, inboundMarshalVTer, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_GetActiveRestoreProcess_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DBService_InterruptRestoreProcess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := local_request_DBService_InterruptRestoreProcess_0(rctx, inboundMarshalVTer, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_InterruptRestoreProcess_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_DBService_CancelRestoreProcess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := local_request_DBService_CancelRestoreProcess_0(rctx, inboundMarshalVTer, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_CancelRestoreProcess_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterDBServiceHandlerFromEndpoint is same as RegisterDBServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterDBServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterDBServiceHandler(ctx, mux, conn)
}

// RegisterDBServiceHandler registers the http handlers for service DBService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterDBServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterDBServiceHandlerClient(ctx, mux, NewDBServiceClient(conn))
}

// RegisterDBServiceHandlerClient registers the http handlers for service DBService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "DBServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "DBServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "DBServiceClient" to call the correct interceptors.
func RegisterDBServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client DBServiceClient) error {

	mux.Handle("GET", pattern_DBService_GetExportCapabilities_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := request_DBService_GetExportCapabilities_0(rctx, inboundMarshalVTer, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_GetExportCapabilities_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_DBService_GetActiveRestoreProcess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := request_DBService_GetActiveRestoreProcess_0(rctx, inboundMarshalVTer, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_GetActiveRestoreProcess_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DBService_InterruptRestoreProcess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := request_DBService_InterruptRestoreProcess_0(rctx, inboundMarshalVTer, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_InterruptRestoreProcess_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_DBService_CancelRestoreProcess_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshalVTer, outboundMarshalVTer := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}
		resp, md, err := request_DBService_CancelRestoreProcess_0(rctx, inboundMarshalVTer, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshalVTer, w, req, err)
			return
		}

		forward_DBService_CancelRestoreProcess_0(ctx, mux, outboundMarshalVTer, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_DBService_GetExportCapabilities_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "db", "exportcaps"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_DBService_GetActiveRestoreProcess_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "db", "restore"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_DBService_InterruptRestoreProcess_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "db", "interruptrestore", "process_id", "attempt_id"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_DBService_CancelRestoreProcess_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "db", "restore", "id"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_DBService_GetExportCapabilities_0 = runtime.ForwardResponseMessage

	forward_DBService_GetActiveRestoreProcess_0 = runtime.ForwardResponseMessage

	forward_DBService_InterruptRestoreProcess_0 = runtime.ForwardResponseMessage

	forward_DBService_CancelRestoreProcess_0 = runtime.ForwardResponseMessage
)
