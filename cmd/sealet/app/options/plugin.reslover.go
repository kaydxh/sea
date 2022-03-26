package options

import (
	"context"

	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/pkg/sealet/provider"
	"github.com/sirupsen/logrus"
)

func (s *CompletedServerRunOptions) installResloverOrDie(ctx context.Context, ws *webserver_.GenericWebServer) {
	c := s.resloverConfig.Complete()
	if !c.Proto.GetEnabled() {
		return
	}

	rs, err := c.New(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("install Reslover, exit")
		return
	}

	ws.AddPostStartHookOrDie("reslover-service", func(ctx context.Context) error {
		return rs.Run(ctx)
	})
	ws.AddPreShutdownHookOrDie("reslover-service", func() error {
		rs.Shutdown()
		return nil
	})

	//use default domains in reslover filed from yaml file
	//you can also add by self, use rs.AddService method
	provider.GlobalProvider().ResolverService = rs
}
