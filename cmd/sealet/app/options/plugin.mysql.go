package options

import (
	"context"

	"github.com/kaydxh/sea/pkg/sealet/provider"
	"github.com/sirupsen/logrus"
)

func (s *CompletedServerRunOptions) installMysqlOrDie(ctx context.Context) {
	c := s.mysqlConfig.Complete()
	if !c.Proto.GetEnabled() {
		return
	}

	db, err := c.New(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("install Mysql, exit")
		return
	}

	provider.GlobalProvider().SqlDB = db

}
