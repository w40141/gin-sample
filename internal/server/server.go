// Package server はHTTPサーバーの設定と起動を提供する
package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/w40141/gin-sample/internal/router"
)

const (
	serverTimeout     = 5 * time.Second
	errChanBufferSize = 1
)

type Server struct {
	*http.Server
}

// New は新しいサーバー設定を作成する
func New(l *slog.Logger) Server {
	// TODO: 環境変数や設定ファイルから取得するようにする
	addr := ":8081"

	handler := router.SetupRouterGin(l)

	return Server{
		&http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

// Start はサーバーを起動します。
func (srv Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errChan := make(chan error, errChanBufferSize)

	go func() {
		if e := srv.ListenAndServe(); e != nil && !errors.Is(e, http.ErrServerClosed) {
			errChan <- e
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), serverTimeout)
		defer cancel()

		slog.Info("shutting down server...")

		if e := srv.Shutdown(shutdownCtx); e != nil {
			return fmt.Errorf("failed to shutdown server: %w", e)
		}
	case e := <-errChan:
		return fmt.Errorf("server error: %w", e)
	}

	return nil
}
