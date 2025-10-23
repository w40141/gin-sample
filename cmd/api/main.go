package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/w40141/gin-sample/internal/server"
	"github.com/w40141/gin-sample/internal/util"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if e := util.Initialize(); e != nil {
		log.Fatal("failed to initialize util: ", e)
	}

	srv := server.New(logger)

	if e := srv.Start(); e != nil {
		log.Fatal("failed to start server: ", e)
	}
}
