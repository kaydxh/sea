package options

import (
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/web/modules/sealet/date"
	sealet_ "github.com/kaydxh/sea/web/sealet"
)

func (s *CompletedServerRunOptions) installWebHandler(ws *webserver_.GenericWebServer) {
	dateApp := date.NewController()
	sealet_.NewWebHandlers(ws, dateApp)
}
