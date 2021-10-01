// Copyright 2021 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Controller struct {
}

func (d *Controller) MetricsPrometheus() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
