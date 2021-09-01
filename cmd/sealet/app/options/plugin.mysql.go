package options

import (
	"github.com/kaydxh/sea/pkg/sealet/provider"
	"github.com/sirupsen/logrus"
)

func (s *CompletedServerRunOptions) installMysqlOrDie() {
	c := s.mysqlConfig.Complete()
	if c.Proto.GetEnabled() {
		logrus.Infof("Installing Mysql")
	}

	db, err := c.New()
	if err != nil {
		logrus.WithError(err).Fatalf("install Mysql, exit")
		return
	}

	provider.GlobalProvider().SqlDB = db
}
