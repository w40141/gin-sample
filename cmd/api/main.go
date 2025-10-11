package main

import (
	"log/slog"
	"os"

	"github.com/w40141/gin-sample/internal/server"
	"github.com/w40141/gin-sample/internal/util"
)

const notOk = 1

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if e := util.Initialize(); e != nil {
		slog.Error("failed to initialize util:", slog.String("cause", e.Error()))
		os.Exit(notOk)
	}

	srv := server.New(logger)

	if e := server.Start(srv); e != nil {
		slog.Error("failed to start server:", slog.String("cause", e.Error()))
		os.Exit(notOk)
	}
}
