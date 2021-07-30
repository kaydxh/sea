package options

import (
	"github.com/spf13/pflag"
)

const defaultSeaConfigFile = "./conf/sea.yaml"

type SeaFlags struct {
	SeaConfigFile string
}

func NewSeaFlags() *SeaFlags {
	return &SeaFlags{
		SeaConfigFile: defaultSeaConfigFile,
	}
}

func (f *SeaFlags) AddFlags(mainfs *pflag.FlagSet) {
	fs := pflag.NewFlagSet("", pflag.ExitOnError)
	fs.StringVar(&f.SeaConfigFile, "config", f.SeaConfigFile, "sea config file")

}

/*
func (s *CompletedServerRunOptions) installFlagsOrDie() {
	config, err := s.Config.Complete().New()
	if err != nil {
		logrus.WithError(err).Fatalf("install Config, exit")
		return
	}

	provider.GlobalProvider().Config = config

	logrus.Infof("install Config: %v", config.String())

}
*/
