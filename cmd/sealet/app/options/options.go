package options

import (
	"context"

	"github.com/go-playground/validator/v10"
	errors_ "github.com/kaydxh/golang/go/errors"
	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	redis_ "github.com/kaydxh/golang/pkg/database/redis"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"
	logs_ "github.com/kaydxh/golang/pkg/logs"
	reslover_ "github.com/kaydxh/golang/pkg/reslover"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
)

type ServerRunOptions struct {
	webServerConfig *webserver_.Config
	logConfig       *logs_.Config
	mysqlConfig     *mysql_.Config
	redisConfig     *redis_.Config
	resloverConfig  *reslover_.Config
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
		webServerConfig: webserver_.NewConfig(gatewayOpts...),
		logConfig:       logs_.NewConfig(logs_.WithViper(viper_.GetViper(configFile, "log"))),
		mysqlConfig:     mysql_.NewConfig(mysql_.WithViper(viper_.GetViper(configFile, "database.mysql"))),
		redisConfig:     redis_.NewConfig(redis_.WithViper(viper_.GetViper(configFile, "database.redis"))),
		resloverConfig:  reslover_.NewConfig(reslover_.WithViper(viper_.GetViper(configFile, "reslover"))),
	}

}

// Validate checks ServerRunOptions and return a slice of found errs.
func (s *ServerRunOptions) Validate(validate *validator.Validate) error {
	var errs []error
	return errors_.NewAggregate(errs)
}

// Complete set default ServerRunOptions.
func (s *ServerRunOptions) Complete() (CompletedServerRunOptions, error) {

	//api30 response formatter
	s.webServerConfig.WithWebConfigOptions(
		webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsTCloud30HTTPResponseOptions()),
	)

	//format error response
	s.webServerConfig.WithWebConfigOptions(
		webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsHttpErrorOptions()),
	)

	return CompletedServerRunOptions{&completedServerRunOptions{s}}, nil
}

// Run runs the specified APIServer.  This should never exit.
func (s *CompletedServerRunOptions) Run(ctx context.Context) error {
	ws, err := s.webServerConfig.Complete().New()
	if err != nil {
		return err
	}

	s.installLogsOrDie()

	//auto installed depend on yaml configure with enabled field
	s.installMysqlOrDie(ctx)
	s.installRedisOrDie(ctx)
	//	s.installPrometheusOrDie(ctx)

	s.installResloverOrDie(ctx, ws)
	//install web handler
	s.installWebHandler(ws)

	prepared, err := ws.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(ctx)
}
