package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"time"

	os_ "github.com/kaydxh/golang/go/os"
	profile_ "github.com/kaydxh/golang/pkg/profile"
	"github.com/kaydxh/sea/cmd/sea-date/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := signal.NotifyContext(context.Background(), os_.ShutdownSignals...)
	defer cancel()

	command := app.NewCommand(ctx)
	{
		defer profile_.StartWithEnv().Stop()
	}

	if err := command.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
