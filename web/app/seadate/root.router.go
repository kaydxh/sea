package seadate

import (
	"context"

	"github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/web/modules/seadate"
)

// NewWebHandlers 注册 SeaDate Controller 到 WebServer。
func NewWebHandlers(ws *webserver.GenericWebServer, c *seadate.Controller) []webserver.WebHandler {
	ws.AddPostStartHookOrDie("web_handler", func(ctx context.Context) error {
		ws.InstallWebHandlers(c)
		return nil
	})
	return []webserver.WebHandler{c}
}