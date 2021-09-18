package options

import "github.com/sirupsen/logrus"

func (s *CompletedServerRunOptions) installLogsOrDie() {

	err := s.logConfig.Complete().Apply()
	if err != nil {
		logrus.WithError(err).Fatalf("install Logs, exit")
		return
	}

}
