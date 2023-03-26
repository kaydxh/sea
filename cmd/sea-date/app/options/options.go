/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
package options

import (
	"context"

	config_ "github.com/kaydxh/golang/pkg/config"
	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	redis_ "github.com/kaydxh/golang/pkg/database/redis"
	logs_ "github.com/kaydxh/golang/pkg/logs"
	opentelemetry_ "github.com/kaydxh/golang/pkg/monitor/opentelemetry"
	resolver_ "github.com/kaydxh/golang/pkg/resolver"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	app_ "github.com/kaydxh/golang/pkg/webserver/app"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"

	//"github.com/kaydxh/sea/cmd/sea-date/app/config"
	"github.com/sirupsen/logrus"
)

type ServerRunOptions struct {
	//Config              *config.Config
	Config              *config_.Config[*v1.Configuration]
	webServerConfig     *webserver_.Config
	logConfig           *logs_.Config
	mysqlConfig         *mysql_.Config
	redisConfig         *redis_.Config
	resolverConfig      *resolver_.Config
	opentelemetryConfig *opentelemetry_.Config
}

// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedServerRunOptions struct {
	*ServerRunOptions
}

// CompletedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type CompletedServerRunOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedServerRunOptions
}

func NewServerRunOptions(configFile string) *ServerRunOptions {
	var gatewayOpts []webserver_.ConfigOption
	gatewayOpts = append(gatewayOpts, webserver_.WithViper(viper_.GetViper(configFile, "web")))

	var config v1.Configuration

	return &ServerRunOptions{
		Config:              config_.NewConfig(&config, config_.WithViper[*v1.Configuration](viper_.GetViper(configFile, ""))),
		webServerConfig:     webserver_.NewConfig(gatewayOpts...),
		logConfig:           logs_.NewConfig(logs_.WithViper(viper_.GetViper(configFile, "log"))),
		mysqlConfig:         mysql_.NewConfig(mysql_.WithViper(viper_.GetViper(configFile, "database.mysql"))),
		redisConfig:         redis_.NewConfig(redis_.WithViper(viper_.GetViper(configFile, "database.redis"))),
		resolverConfig:      resolver_.NewConfig(resolver_.WithViper(viper_.GetViper(configFile, "reslover"))),
		opentelemetryConfig: opentelemetry_.NewConfig(opentelemetry_.WithViper(viper_.GetViper(configFile, "web.monitor.open_telemetry"))),
	}

}

// Complete set default ServerRunOptions.
func (s *ServerRunOptions) Complete() (CompletedServerRunOptions, error) {

	return CompletedServerRunOptions{&completedServerRunOptions{s}}, nil
}

// Run runs the specified APIServer.  This should never exit.
func (s *CompletedServerRunOptions) Run(ctx context.Context) error {
	logrus.Infof("Starting Sea let version: %v", app_.GetVersion())
	ws, err := s.webServerConfig.Complete().New(ctx)
	if err != nil {
		return err
	}

	s.installLogsOrDie()
	s.installConfigOrDie()

	//auto installed depend on yaml configure with enabled field
	s.installMysqlOrDie(ctx)
	s.installRedisOrDie(ctx)
	s.installOpenTelemetryOrDie(ctx)
	//	s.installPrometheusOrDie(ctx)

	s.installResolverOrDie(ctx, ws)
	//install web handler
	s.installWebHandlerOrDie(ws)

	prepared, err := ws.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(ctx)
}
