package app

import (
	"context"
	"fmt"
	"os"

	"github.com/kaydxh/sea/cmd/sealet/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	componentSealet = "sealet"
)

// NewSealetCommand creates a *cobra.Command object with default parameters
func NewSealetCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sea-let",
		Short: "sea-let Public HTTP/2 and GRPC APIs",
		// stop printing usage when the command errors
		Long: `The Sea let is a gateway serve which you can use curl over HTTP 1.1 or grpc protocal on the same host:port.
Example: curl -X POST -k https://localhost:port/Now
See [Sea](https://github.com/kaydxh/sea/blob/master/README.md) for more information.`,
		//SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			return runCommand(ctx, cmd)
		},

		PostRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("server exit")
			return nil
		},

		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					//%q a single-quoted character literal safely escaped with Go syntax
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	cobra.OnInitialize(func() {})

	var cfgFile string
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultConfigPath(),
		fmt.Sprintf("Config file (default is %q)", defaultConfigPath()))

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("FLAG: --%s=%q\n", flag.Name, flag.Value)
		})
		return nil
	}

	return cmd
}

func runCommand(ctx context.Context, cmd *cobra.Command) error {
	cfgFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	s := options.NewServerRunOptions(cfgFile)

	// set default options
	completedOptions, err := s.Complete()
	if err != nil {
		return err
	}

	// validate options
	if err := completedOptions.Validate(nil); err != nil {
		return err
	}

	if err := completedOptions.Run(ctx); err != nil {
		fmt.Printf("failed to run server")
		os.Exit(1)
	}

	return nil
}

// defaultConfigPath returns config file's default path
func defaultConfigPath() string {
	return fmt.Sprintf("./conf/%s.yaml", componentSealet)
}
