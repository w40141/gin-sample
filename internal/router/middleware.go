// Package router はHTTPルーターの設定とミドルウェアを提供する
package router

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/w40141/gin-sample/internal/util"
)

// Logger はリクエストとレスポンスのログを記録するミドルウェアです。
func Logger(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := util.Now()
		startStr := start.Format(time.RFC3339)
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		params := map[string]string{}
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}

		attrs := []slog.Attr{
			slog.String("Time", startStr),
			slog.String("Method", method),
			slog.String("Path", path),
			slog.String("Query", query),
			slog.Any("Params", params),
			slog.String("ClientIP", clientIP),
			slog.String("User Agent", userAgent),
		}

		c.Next()

		end := util.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()

		var logLevel slog.Level

		if status >= http.StatusInternalServerError {
			logLevel = slog.LevelError
		} else if status >= http.StatusBadRequest {
			logLevel = slog.LevelWarn
		} else {
			logLevel = slog.LevelInfo
		}

		attrs = append(
			attrs,
			slog.String("Latency", latency.String()),
			slog.Int("Status", status),
		)

		l.LogAttrs(c, logLevel, "GIN", attrs...)
	}
}
