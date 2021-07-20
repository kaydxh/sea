// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package health

import (
	_ "expvar"
	_ "net/http/pprof"

	//	"git.code.oa.com/YouTu-BizRD-1/yt-transfer-proxy/internal/pkg/provider"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/x/healthx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/searKing/golang/pkg/net/http"
)

type Controller struct {
	h *healthx.Handler
}

func NewController() *Controller {
	h := &Controller{}
	//	h.init()
	return h
}

/*
func (d *Controller) init() {
	c := provider.GlobalProvider()
	//dependency.ExpectDependency(c.Logger(), map[string]interface{}{"service_discovery": c.ServiceDiscoveryConnection})

	logger := logrusx.New("", "")
	logger.Logger = logrus.StandardLogger()
	w := herodot.NewJSONWriter(logger)
	d.h = healthx.NewHandler(w, c.Proto().GetAppInfo().GetBuildVersion(), healthx.ReadyCheckers{
		"database": provider.GlobalProvider().SqlDBPing,
		//"zookeeper": ctx.ServiceDiscoveryConnection.Ping,
	})
}
*/

func (d *Controller) Health() gin.HandlerFunc {
	router := httprouter.New()
	d.h.SetRoutes(router, true)
	return gin.WrapF(router.ServeHTTP)
}

func (d *Controller) MetricsPrometheus() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}

func (d *Controller) Alive() gin.HandlerFunc {
	return gin.WrapH(http.WrapHTTPRouterF(d.h.Alive))
}

func (d *Controller) Ready(shareErrors bool) gin.HandlerFunc {
	return gin.WrapH(http.WrapHTTPRouterF(d.h.Ready(shareErrors)))
}

func (d *Controller) Version() gin.HandlerFunc {
	return gin.WrapH(http.WrapHTTPRouterF(d.h.Version))
}
