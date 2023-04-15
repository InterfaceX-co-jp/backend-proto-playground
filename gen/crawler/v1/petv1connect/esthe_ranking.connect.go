// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: crawler/v1/esthe_ranking.proto

package petv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/InterfaceX-co-jp/backend-proto-playground/gen/crawler/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EstheRankingServiceName is the fully-qualified name of the EstheRankingService service.
	EstheRankingServiceName = "pet.v1.EstheRankingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EstheRankingServiceSyncTherapistsBySpanProcedure is the fully-qualified name of the
	// EstheRankingService's SyncTherapistsBySpan RPC.
	EstheRankingServiceSyncTherapistsBySpanProcedure = "/pet.v1.EstheRankingService/SyncTherapistsBySpan"
)

// EstheRankingServiceClient is a client for the pet.v1.EstheRankingService service.
type EstheRankingServiceClient interface {
	SyncTherapistsBySpan(context.Context, *connect_go.Request[v1.SyncTherapistsBySpanRequest]) (*connect_go.Response[v1.SyncTherapistsBySpanResponse], error)
}

// NewEstheRankingServiceClient constructs a client for the pet.v1.EstheRankingService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEstheRankingServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EstheRankingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &estheRankingServiceClient{
		syncTherapistsBySpan: connect_go.NewClient[v1.SyncTherapistsBySpanRequest, v1.SyncTherapistsBySpanResponse](
			httpClient,
			baseURL+EstheRankingServiceSyncTherapistsBySpanProcedure,
			opts...,
		),
	}
}

// estheRankingServiceClient implements EstheRankingServiceClient.
type estheRankingServiceClient struct {
	syncTherapistsBySpan *connect_go.Client[v1.SyncTherapistsBySpanRequest, v1.SyncTherapistsBySpanResponse]
}

// SyncTherapistsBySpan calls pet.v1.EstheRankingService.SyncTherapistsBySpan.
func (c *estheRankingServiceClient) SyncTherapistsBySpan(ctx context.Context, req *connect_go.Request[v1.SyncTherapistsBySpanRequest]) (*connect_go.Response[v1.SyncTherapistsBySpanResponse], error) {
	return c.syncTherapistsBySpan.CallUnary(ctx, req)
}

// EstheRankingServiceHandler is an implementation of the pet.v1.EstheRankingService service.
type EstheRankingServiceHandler interface {
	SyncTherapistsBySpan(context.Context, *connect_go.Request[v1.SyncTherapistsBySpanRequest]) (*connect_go.Response[v1.SyncTherapistsBySpanResponse], error)
}

// NewEstheRankingServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEstheRankingServiceHandler(svc EstheRankingServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(EstheRankingServiceSyncTherapistsBySpanProcedure, connect_go.NewUnaryHandler(
		EstheRankingServiceSyncTherapistsBySpanProcedure,
		svc.SyncTherapistsBySpan,
		opts...,
	))
	return "/pet.v1.EstheRankingService/", mux
}

// UnimplementedEstheRankingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEstheRankingServiceHandler struct{}

func (UnimplementedEstheRankingServiceHandler) SyncTherapistsBySpan(context.Context, *connect_go.Request[v1.SyncTherapistsBySpanRequest]) (*connect_go.Response[v1.SyncTherapistsBySpanResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("pet.v1.EstheRankingService.SyncTherapistsBySpan is not implemented"))
}
