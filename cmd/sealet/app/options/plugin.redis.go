package options

import (
	"github.com/kaydxh/sea/pkg/sealet/provider"
	"github.com/sirupsen/logrus"
)

func (s *CompletedServerRunOptions) installRedisOrDie() {
	c := s.redisConfig.Complete()
	if c.Proto.GetEnabled() {
		logrus.Infof("Installing Redis")
	}

	db, err := c.New()
	if err != nil {
		logrus.WithError(err).Fatalf("install Redis, exit")
		return
	}

	provider.GlobalProvider().RedisDB = db
}
