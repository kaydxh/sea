// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package monitor

import "github.com/gin-gonic/gin"

func Router(router gin.IRouter) gin.IRouter {
	s := Controller{}
	router.Any("/metrics/prometheus/*path", s.MetricsPrometheus())

	return router
}
