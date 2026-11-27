package metrics

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryClientInterceptor records count and latency of unary gRPC calls.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)

		grpcClientRequestsTotal.WithLabelValues(method, status.Code(err).String()).Inc()
		grpcClientRequestDuration.WithLabelValues(method).Observe(time.Since(start).Seconds())

		return err
	}
}
