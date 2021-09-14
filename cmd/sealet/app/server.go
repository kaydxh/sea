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
		Use:   "sealet",
		Short: "sealet Public HTTP/2 and GRPC APIs",
		// stop printing usage when the command errors
		Long: `To get started run the serve subcommand which will start a gateway server
			   You can use curl over HTTP 1.1, 
			   eg: curl -X POST -k https://localhost:port/Now `,
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
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

// defaultConfigPath returns config file's default path
func defaultConfigPath() string {
	return fmt.Sprintf("./conf/%s.yaml", componentSealet)
}
