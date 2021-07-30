// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package health

import (
	"github.com/gin-gonic/gin"
)

func Router(router gin.IRouter) gin.IRouter {
	health := NewController()
	_ = health
	/*
		router.GET(values.HealthAliveCheckPath, health.Alive())
		router.GET(values.HealthReadyCheckPath, health.Ready(true))
		router.GET(values.HealthMetricsPrometheusPath, health.MetricsPrometheus())

		router.GET(values.HealthVersionPath, health.Version())
	*/
	return router
}
