// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seadate

import (
	"context"

	"github.com/gin-gonic/gin"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/sea-date/application"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// can add handler in Controller
func NewController(app application.Application) *Controller {
	return &Controller{
		app: app,
	}
}

func (c *Controller) SetRoutes(ginRouter gin.IRouter, grpcRouter *gw_.GRPCGateway) {
	grpcRouter.RegisterGRPCHandler(func(srv *grpc.Server) {
		v1.RegisterSeaDateServiceServer(srv, &LocalController{Controller: c})
	})
	_ = grpcRouter.RegisterHTTPHandler(
		context.Background(),
		func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
			//return date.RegisterDateServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
			//https://github.com/grpc-ecosystem/grpc-gateway/issues/1458,
			//impove performace, but grpc interceptor is disabled
			//return v1.RegisterSeaDateServiceHandlerServer(ctx, mux, c)
			return v1.RegisterSeaDateServiceHandlerServer(ctx, mux, &LocalController{Controller: c})
		},
	)
}
