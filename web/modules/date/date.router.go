// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"context"

	"github.com/gin-gonic/gin"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"
	"github.com/kaydxh/sea/api/openapi-spec/date/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

/*
func Router(router *gw_.GRPCGateway) *gw_.GRPCGateway {
	s := Controller{}
	router.RegisterGRPCHandler(func(srv *grpc.Server) {
		date.RegisterDateServiceServer(srv, &s)
	})
	err := router.RegisterHTTPHandler(
		context.Background(),
		func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
			//		return date.RegisterDateServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
			//https://github.com/grpc-ecosystem/grpc-gateway/issues/1458,
			//impove performace, but grpc interceptor is disabled
			return date.RegisterDateServiceHandlerServer(ctx, mux, &s)
		},
	)
	if err != nil {
		return nil
	}

	return router
}
*/

// can add handler in Controller
func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetRoutes(ginRouter gin.IRouter, grpcRouter *gw_.GRPCGateway) {
	grpcRouter.RegisterGRPCHandler(func(srv *grpc.Server) {
		date.RegisterDateServiceServer(srv, c)
	})
	_ = grpcRouter.RegisterHTTPHandler(
		context.Background(),
		func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
			//return date.RegisterDateServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
			//https://github.com/grpc-ecosystem/grpc-gateway/issues/1458,
			//impove performace, but grpc interceptor is disabled
			return date.RegisterDateServiceHandlerServer(ctx, mux, c)
		},
	)
}
