package options

import (
	"fmt"

	"github.com/kaydxh/sea/pkg/sea-date/provider"
	"github.com/sirupsen/logrus"
)

func (s *CompletedServerRunOptions) installConfigOrDie() {
	config, err := s.Config.Complete().New()
	if err != nil {
		logrus.WithError(err).Fatalf("failed to install Config, exit")
		return
	}

	provider.GlobalProvider().Config = config
	logrus.Infof("config loaded: %s", maskedConfigSummary(config))
}

// maskedConfigSummary 返回日志友好的脱敏配置摘要，避免打印完整敏感信息。
func maskedConfigSummary(config interface{}) string {
	if config == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", config)
}
