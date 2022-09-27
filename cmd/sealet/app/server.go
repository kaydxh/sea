/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
package app

import (
	"context"
	"fmt"
	"os"

	app_ "github.com/kaydxh/golang/pkg/webserver/app"
	"github.com/kaydxh/sea/cmd/sealet/app/options"
	"github.com/spf13/cobra"
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

	appFlag := app_.NewAppFlags(cmd)
	appFlag.Install()

	/*
		var cfgFile string
		cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultConfigPath(),
			fmt.Sprintf("Config file (default is %q)", defaultConfigPath()))
	*/

	/*
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			cmd.Flags().VisitAll(func(flag *pflag.Flag) {
				fmt.Printf("FLAG: --%s=%q\n", flag.Name, flag.Value)
			})
			return nil
		}
	*/

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

	if err := completedOptions.Run(ctx); err != nil {
		fmt.Printf("failed to run server")
		os.Exit(1)
	}

	return nil
}
