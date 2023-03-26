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
	"github.com/kaydxh/sea/cmd/sea-date/app/options"
	"github.com/spf13/cobra"
)

// NewCommand creates a *cobra.Command object with default parameters
func NewCommand(ctx context.Context) *cobra.Command {
	return app_.NewCommand(ctx, runCommand)
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
		fmt.Printf("failed to run server, err: %v\n", err)
		os.Exit(1)
	}

	return nil
}
