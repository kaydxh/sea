package options

import (
	"github.com/kaydxh/sea/pkg/sealet/provider"
	"github.com/sirupsen/logrus"
)

func (s *CompletedServerRunOptions) installConfigOrDie() {

	config, err := s.Config.Complete().New()
	if err != nil {
		logrus.WithError(err).Fatalf("failed to install Config, exit")
		return
	}
	provider.GlobalProvider().Config = config
}
