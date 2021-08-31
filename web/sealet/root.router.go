// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sealet

import (
	"github.com/gin-gonic/gin"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"

	//proxy_ "github.com/kaydxh/golang/pkg/proxy"
	"github.com/kaydxh/sea/web/modules/sealet/date"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// SetRoutes registers this handler's routes.
func (h *Handler) SetRoutes(ginRouter gin.IRouter, grpcRouter *gw_.GRPCGateway) {
	// bind grpcGateway as default

	//	middlewares.MiddlewaresRouter(ginRouter)
	//	index.Router(ginRouter)
	//	debug.Router(ginRouter, "")
	//	health.Router(ginRouter)
	// webapp static files
	//	webapp.Router(ginRouter)
	// doc
	//	swagger.Router(ginRouter)
	// API
	//	apiRouter := ginRouter.Group(values.APIPathPrefix)
	//	health.Router(apiRouter)

	//reverse proxy for  path "/proxy",
	/*
		addr := "http://127.0.0.1:1080"
		rp, err := proxy_.NewReverseProxy(ginRouter, proxy_.WithTargetUrl(addr), proxy_.WithRouterPatterns("/Proxy"))
		if err == nil {
			rp.SetProxy()
		}
	*/

	date.Router(grpcRouter)

	//// NOTE: It might be required to set Router.HandleMethodNotAllowed to false to avoid problems.
	//r.HandleMethodNotAllowed = false
	//r.NotFound = Routes(h.c, values.PathPrefix)
}
