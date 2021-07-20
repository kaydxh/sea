package options

import (
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/web/app"
)

func (s *CompletedServerRunOptions) installWebHandler(ws *webserver_.GenericWebServer) {
	ws.InstallWebHandlers(app.NewHandler())
}
