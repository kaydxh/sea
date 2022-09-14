package options

import (
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	sealet_ "github.com/kaydxh/sea/web/app/sealet"
	"github.com/kaydxh/sea/web/modules/date"
)

func (s *CompletedServerRunOptions) installWebHandler(ws *webserver_.GenericWebServer) {
	dateApp := date.NewController()
	sealet_.NewWebHandlers(ws, dateApp)
}
