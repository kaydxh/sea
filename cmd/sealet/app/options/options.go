package options

import (
	"context"

	"github.com/go-playground/validator/v10"
	errors_ "github.com/kaydxh/golang/go/errors"
	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	gw_ "github.com/kaydxh/golang/pkg/grpc-gateway"
	logs_ "github.com/kaydxh/golang/pkg/logs"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/sirupsen/logrus"
)

type ServerRunOptions struct {
	webServerConfig *webserver_.Config
	logConfig       *logs_.Config
	mysqlConfig     *mysql_.Config
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
	gatewayOpts := []webserver_.ConfigOption{}
	gatewayOpts = append(gatewayOpts, webserver_.WithViper(viper_.GetViper(configFile, "web")))

	//api response formatter
	gatewayOpts = append(
		gatewayOpts,
		webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsTCloud30HTTPResponseOptions()),
	)

	//format error response
	gatewayOpts = append(
		gatewayOpts,
		webserver_.WithGRPCGatewayOptions(gw_.WithServerInterceptorsHttpErrorOptions()),
	)

	//auto generate requestId
	gatewayOpts = append(
		gatewayOpts,
		webserver_.WithGRPCGatewayOptions(gw_.WithServerUnaryInterceptorsRequestIdOptions()),
	)

	return &ServerRunOptions{
		webServerConfig: webserver_.NewConfig(gatewayOpts...),
		logConfig:       logs_.NewConfig(logs_.WithViper(viper_.GetViper(configFile, "log"))),
		mysqlConfig:     mysql_.NewConfig(mysql_.WithViper(viper_.GetViper(configFile, "database.mysql"))),
	}
}

// Validate checks ServerRunOptions and return a slice of found errs.
func (s *ServerRunOptions) Validate(validate *validator.Validate) error {
	var errs []error
	return errors_.NewAggregate(errs)
}

// Complete set default ServerRunOptions.
func (s *ServerRunOptions) Complete() (CompletedServerRunOptions, error) {
	_ = s
	return CompletedServerRunOptions{&completedServerRunOptions{s}}, nil
}

// Run runs the specified APIServer.  This should never exit.
func (s *CompletedServerRunOptions) Run(ctx context.Context) error {
	ws, err := s.webServerConfig.Complete().New()
	if err != nil {
		return err
	}

	logrus.Infof("Installing Logs")
	s.installLogsOrDie()

	//install web handler
	logrus.Infof("Installing WebHandler")
	s.installWebHandler(ws)

	//below, auto installed depend on yaml configure with enabled field
	s.installMysqlOrDie()

	prepared, err := ws.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(ctx)
}
