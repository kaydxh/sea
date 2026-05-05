package options

import (
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/pkg/sea-date/application"
	local_ "github.com/kaydxh/sea/pkg/sea-date/infrastructure/local"
	seadateapp_ "github.com/kaydxh/sea/web/app/seadate"
	"github.com/kaydxh/sea/web/modules/seadate"
	"github.com/sirupsen/logrus"
)

// installWebHandlerOrDie 安装 Web Handler（Controller 层），初始化完整的依赖链：
// DateRepository → SeaDateHandler → Application → Controller → NewWebHandlers
func (s *CompletedServerRunOptions) installWebHandlerOrDie(ws *webserver_.GenericWebServer) {
	// 1. 初始化 Date 相关依赖
	dateRepo := &local_.Repository{}

	// 2. 组装 Application
	app := application.Application{
		Commands: application.Commands{
			SeaDateHandler: application.NewSeaDateHandler(dateRepo),
		},
	}

	// 3. 创建 Controller 并注册到 WebServer
	dateCtrl := seadate.NewController(app)
	seadateapp_.NewWebHandlers(ws, dateCtrl)

	logrus.Info("[WebHandler] SeaDate web handlers installed")
}