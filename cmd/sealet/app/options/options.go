package options

import (
	"context"

	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	redis_ "github.com/kaydxh/golang/pkg/database/redis"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"
	logs_ "github.com/kaydxh/golang/pkg/logs"
	resolver_ "github.com/kaydxh/golang/pkg/resolver"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	app_ "github.com/kaydxh/golang/pkg/webserver/app"
	"github.com/kaydxh/sea/cmd/sealet/app/config"
	"github.com/sirupsen/logrus"
)

type ServerRunOptions struct {
	Config          *config.Config
	webServerConfig *webserver_.Config
	logConfig       *logs_.Config
	mysqlConfig     *mysql_.Config
	redisConfig     *redis_.Config
	resolverConfig  *resolver_.Config
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

	return &ServerRunOptions{
		Config:          config.NewConfig(config.WithViper(viper_.GetViper(configFile, ""))),
		webServerConfig: webserver_.NewConfig(gatewayOpts...),
		logConfig:       logs_.NewConfig(logs_.WithViper(viper_.GetViper(configFile, "log"))),
		mysqlConfig:     mysql_.NewConfig(mysql_.WithViper(viper_.GetViper(configFile, "database.mysql"))),
		redisConfig:     redis_.NewConfig(redis_.WithViper(viper_.GetViper(configFile, "database.redis"))),
		resolverConfig:  resolver_.NewConfig(resolver_.WithViper(viper_.GetViper(configFile, "reslover"))),
	}

}

// Complete set default ServerRunOptions.
func (s *ServerRunOptions) Complete() (CompletedServerRunOptions, error) {

	s.webServerConfig.WithWebConfigOptions(
		//api30 response formatter
		//webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsTCloud30HTTPResponseOptions()),
		//webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsHttpErrorOptions()),

		//trivial v1 response formatter
		webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsTrivialV1HTTPResponseOptions()),
		//format error response
		webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsTrivialV1HttpErrorOptions()),
	)

	return CompletedServerRunOptions{&completedServerRunOptions{s}}, nil
}

// Run runs the specified APIServer.  This should never exit.
func (s *CompletedServerRunOptions) Run(ctx context.Context) error {
	logrus.Infof("Starting Sea let version: %v", app_.GetVersion())
	ws, err := s.webServerConfig.Complete().New()
	if err != nil {
		return err
	}

	s.installLogsOrDie()
	s.installConfigOrDie()

	//auto installed depend on yaml configure with enabled field
	s.installMysqlOrDie(ctx)
	s.installRedisOrDie(ctx)
	//	s.installPrometheusOrDie(ctx)

	s.installResolverOrDie(ctx, ws)
	//install web handler
	s.installWebHandler(ws)

	prepared, err := ws.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(ctx)
}
