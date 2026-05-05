package seadate

import (
	"context"

	"github.com/gin-gonic/gin"
	http_ "github.com/kaydxh/golang/go/net/http"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"
	httpinterceptordebug_ "github.com/kaydxh/golang/pkg/middleware/http-middleware/debug"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// SetRoutes 注册 gRPC 和 HTTP 路由，包括中间件。
func (c *Controller) SetRoutes(ginRouter gin.IRouter, grpcRouter *gw_.GRPCGateway) {
	// 1. 启用 HTTP 请求/响应 body 打印中间件（公共库带截断的版本）
	grpcRouter.ApplyOptions(
		gw_.WithHttpHandlerInterceptorOptions(http_.HandlerInterceptor{
			Interceptor: httpinterceptordebug_.InOutputPrinterWithTruncate,
		}),
	)

	// 2. 注册 gRPC handler
	grpcRouter.RegisterGRPCHandler(func(srv *grpc.Server) {
		v1.RegisterSeaDateServiceServer(srv, c)
	})

	// 3. 注册 HTTP handler（直接使用 HandlerServer，无需 LocalController 包装）
	_ = grpcRouter.RegisterHTTPHandler(
		context.Background(),
		func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
			return v1.RegisterSeaDateServiceHandlerServer(ctx, mux, c)
		},
	)
}