// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sealet

import (
	"context"

	"github.com/kaydxh/golang/pkg/webserver"
	//proxy_ "github.com/kaydxh/golang/pkg/proxy"
	"github.com/kaydxh/sea/web/modules/seadate"
)

func NewWebHandlers(ws *webserver.GenericWebServer, c *seadate.Controller) []webserver.WebHandler {
	ws.AddPostStartHookOrDie("web_handler", func(ctx context.Context) error {
		ws.InstallWebHandlers(c)
		return nil
	})
	return []webserver.WebHandler{c}
}
