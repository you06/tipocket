package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pingcap/tipocket/pkg/pocket/config"
	"github.com/pingcap/tipocket/pkg/pocket/core"
)

var (
	configPath = flag.String("config", "", "config file path")
)

func init() {
	flag.Parse()
}

func main() {
	cfg := config.Init()
	if *configPath != "" {
		if err := cfg.Load(*configPath); err != nil {
			panic(err)
		}
	}

	pCore := core.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc,
			os.Interrupt,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)

		fmt.Printf("Got signal %d to exit.\n", <-sc)
		cancel()
	}()

	fmt.Println(pCore.Start(ctx))
}
