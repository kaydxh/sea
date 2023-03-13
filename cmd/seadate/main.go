/*
Copyright 2020 The kaydxh Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"time"

	os_ "github.com/kaydxh/golang/go/os"
	profile_ "github.com/kaydxh/golang/pkg/profile"
	"github.com/kaydxh/sea/cmd/seadate/app"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	ctx, cancel := signal.NotifyContext(context.Background(), os_.ShutdownSignals...)
	defer cancel()

	command := app.NewCommand(ctx)
	// profile
	{
		//env variable PROFILING=cpu[mem,mutex,block,trace,thread_create,goroutine]
		//sudo PROFILING="cpu" PROFILEPATH="./profile" ./bin/seadate --config ./conf/seadate.yaml
		//<ctrl-c>
		defer profile_.StartWithEnv().Stop()
	}

	if err := command.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
