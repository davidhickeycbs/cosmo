// https://protobuf.dev/programming-guides/style/

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: wg/cosmo/graphqlmetrics/v1/graphqlmetrics.proto

package graphqlmetricsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/wundergraph/cosmo/graphqlmetrics/gen/proto/wg/cosmo/graphqlmetrics/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GraphQLMetricsServiceName is the fully-qualified name of the GraphQLMetricsService service.
	GraphQLMetricsServiceName = "wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GraphQLMetricsServicePublishGraphQLMetricsProcedure is the fully-qualified name of the
	// GraphQLMetricsService's PublishGraphQLMetrics RPC.
	GraphQLMetricsServicePublishGraphQLMetricsProcedure = "/wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService/PublishGraphQLMetrics"
	// GraphQLMetricsServicePublishAggregatedGraphQLMetricsProcedure is the fully-qualified name of the
	// GraphQLMetricsService's PublishAggregatedGraphQLMetrics RPC.
	GraphQLMetricsServicePublishAggregatedGraphQLMetricsProcedure = "/wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService/PublishAggregatedGraphQLMetrics"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	graphQLMetricsServiceServiceDescriptor                               = v1.File_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto.Services().ByName("GraphQLMetricsService")
	graphQLMetricsServicePublishGraphQLMetricsMethodDescriptor           = graphQLMetricsServiceServiceDescriptor.Methods().ByName("PublishGraphQLMetrics")
	graphQLMetricsServicePublishAggregatedGraphQLMetricsMethodDescriptor = graphQLMetricsServiceServiceDescriptor.Methods().ByName("PublishAggregatedGraphQLMetrics")
)

// GraphQLMetricsServiceClient is a client for the wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService
// service.
type GraphQLMetricsServiceClient interface {
	// PublishGraphQLMetrics publishes the GraphQL metrics to the metrics service
	PublishGraphQLMetrics(context.Context, *connect.Request[v1.PublishGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishOperationCoverageReportResponse], error)
	PublishAggregatedGraphQLMetrics(context.Context, *connect.Request[v1.PublishAggregatedGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishAggregatedGraphQLRequestMetricsResponse], error)
}

// NewGraphQLMetricsServiceClient constructs a client for the
// wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGraphQLMetricsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GraphQLMetricsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &graphQLMetricsServiceClient{
		publishGraphQLMetrics: connect.NewClient[v1.PublishGraphQLRequestMetricsRequest, v1.PublishOperationCoverageReportResponse](
			httpClient,
			baseURL+GraphQLMetricsServicePublishGraphQLMetricsProcedure,
			connect.WithSchema(graphQLMetricsServicePublishGraphQLMetricsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		publishAggregatedGraphQLMetrics: connect.NewClient[v1.PublishAggregatedGraphQLRequestMetricsRequest, v1.PublishAggregatedGraphQLRequestMetricsResponse](
			httpClient,
			baseURL+GraphQLMetricsServicePublishAggregatedGraphQLMetricsProcedure,
			connect.WithSchema(graphQLMetricsServicePublishAggregatedGraphQLMetricsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// graphQLMetricsServiceClient implements GraphQLMetricsServiceClient.
type graphQLMetricsServiceClient struct {
	publishGraphQLMetrics           *connect.Client[v1.PublishGraphQLRequestMetricsRequest, v1.PublishOperationCoverageReportResponse]
	publishAggregatedGraphQLMetrics *connect.Client[v1.PublishAggregatedGraphQLRequestMetricsRequest, v1.PublishAggregatedGraphQLRequestMetricsResponse]
}

// PublishGraphQLMetrics calls
// wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService.PublishGraphQLMetrics.
func (c *graphQLMetricsServiceClient) PublishGraphQLMetrics(ctx context.Context, req *connect.Request[v1.PublishGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishOperationCoverageReportResponse], error) {
	return c.publishGraphQLMetrics.CallUnary(ctx, req)
}

// PublishAggregatedGraphQLMetrics calls
// wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService.PublishAggregatedGraphQLMetrics.
func (c *graphQLMetricsServiceClient) PublishAggregatedGraphQLMetrics(ctx context.Context, req *connect.Request[v1.PublishAggregatedGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishAggregatedGraphQLRequestMetricsResponse], error) {
	return c.publishAggregatedGraphQLMetrics.CallUnary(ctx, req)
}

// GraphQLMetricsServiceHandler is an implementation of the
// wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService service.
type GraphQLMetricsServiceHandler interface {
	// PublishGraphQLMetrics publishes the GraphQL metrics to the metrics service
	PublishGraphQLMetrics(context.Context, *connect.Request[v1.PublishGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishOperationCoverageReportResponse], error)
	PublishAggregatedGraphQLMetrics(context.Context, *connect.Request[v1.PublishAggregatedGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishAggregatedGraphQLRequestMetricsResponse], error)
}

// NewGraphQLMetricsServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGraphQLMetricsServiceHandler(svc GraphQLMetricsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	graphQLMetricsServicePublishGraphQLMetricsHandler := connect.NewUnaryHandler(
		GraphQLMetricsServicePublishGraphQLMetricsProcedure,
		svc.PublishGraphQLMetrics,
		connect.WithSchema(graphQLMetricsServicePublishGraphQLMetricsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	graphQLMetricsServicePublishAggregatedGraphQLMetricsHandler := connect.NewUnaryHandler(
		GraphQLMetricsServicePublishAggregatedGraphQLMetricsProcedure,
		svc.PublishAggregatedGraphQLMetrics,
		connect.WithSchema(graphQLMetricsServicePublishAggregatedGraphQLMetricsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GraphQLMetricsServicePublishGraphQLMetricsProcedure:
			graphQLMetricsServicePublishGraphQLMetricsHandler.ServeHTTP(w, r)
		case GraphQLMetricsServicePublishAggregatedGraphQLMetricsProcedure:
			graphQLMetricsServicePublishAggregatedGraphQLMetricsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGraphQLMetricsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGraphQLMetricsServiceHandler struct{}

func (UnimplementedGraphQLMetricsServiceHandler) PublishGraphQLMetrics(context.Context, *connect.Request[v1.PublishGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishOperationCoverageReportResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService.PublishGraphQLMetrics is not implemented"))
}

func (UnimplementedGraphQLMetricsServiceHandler) PublishAggregatedGraphQLMetrics(context.Context, *connect.Request[v1.PublishAggregatedGraphQLRequestMetricsRequest]) (*connect.Response[v1.PublishAggregatedGraphQLRequestMetricsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService.PublishAggregatedGraphQLMetrics is not implemented"))
}
