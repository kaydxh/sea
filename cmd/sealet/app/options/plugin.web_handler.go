package options

import (
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/web/sealet"
)

func (s *CompletedServerRunOptions) installWebHandler(ws *webserver_.GenericWebServer) {
	ws.InstallWebHandlers(sealet.NewHandler())
}
