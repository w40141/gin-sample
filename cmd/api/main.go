package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/w40141/gin-sample/internal/router"
	"github.com/w40141/gin-sample/internal/web"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, e := router.LoadConfig()
	if e != nil {
		logger.Error(fmt.Sprintf("failed to load config: %v", e))
		os.Exit(1)
	}

	r := web.SetupRouter(cfg, logger)

	r = web.Handler(r)

	if e := cfg.Start(r); e != nil {
		logger.Error(fmt.Sprintf("failed to run server: %v", e))
		os.Exit(1)
	}
}
