package app

import (
	"context"
	"fmt"
	"os"

	"github.com/kaydxh/sea/cmd/app/options"
	"github.com/spf13/cobra"
)

const (
	componentSealet = "sealet"
)

// NewSealetCommand creates a *cobra.Command object with default parameters
func NewSealetCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "sealet",
		Short: "sealet Public HTTP/2 and GRPC APIs",
		// stop printing usage when the command errors
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s := options.NewServerRunOptions()

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
				fmt.Printf("failed to run")
				os.Exit(1)
			}

			return nil
		},
	}
}

/*
// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedServerRunOptions struct {
	*options.ServerRunOptions
}

// Complete set default ServerRunOptions.
func Complete(s *options.ServerRunOptions) (completedServerRunOptions, error) {
	var options completedServerRunOptions
	if err := completeServer(s); err != nil {
		return options, err
	}

	options.ServerRunOptions = s
	return options, nil
}

func completeServer(s *options.ServerRunOptions) error {
	return nil
}

// Run runs the specified APIServer.  This should never exit.
func Run(ctx context.Context, completeOptions completedServerRunOptions) error {
	// To help debugging, immediately log version
	server, err := CreateServerChain(completeOptions, stopCh)
	if err != nil {
		return err
	}

	prepared, err := server.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(ctx)
}
*/
