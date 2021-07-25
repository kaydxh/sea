package options

import (
	"context"

	"github.com/go-playground/validator/v10"
	errors_ "github.com/kaydxh/golang/go/errors"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
)

type ServerRunOptions struct {
	webServerConfig *webserver_.Config
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
	return &ServerRunOptions{
		webServerConfig: webserver_.NewConfig(webserver_.WithViper(viper_.GetViper(configFile, "web"))),
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

	s.installWebHandler(ws)

	prepared, err := ws.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(ctx)
}
