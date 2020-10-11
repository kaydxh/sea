package main

import (
	"github.com/kaydxh/sea/pkg/cmd/sea"
	"os"
)

func main() {

	command := sea.NewCommand()

	//	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	//	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	// utilflag.InitFlags()
	//	logs.InitLogs()
	//	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
