package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, e := LoadConfig()
	if e != nil {
		logger.Error(fmt.Sprintf("failed to load config: %v", e))
		os.Exit(1)
	}

	r := SetupRouter(cfg, logger)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // デフォルトで0.0.0.0:8080で待機します
}

type Config struct {
	loc           *time.Location
	timeFormatter string
}

func LoadConfig() (Config, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return Config{}, err
	}

	timeFormatter := time.RFC3339

	return Config{
		loc:           loc,
		timeFormatter: timeFormatter,
	}, nil
}

func SetupRouter(cfg Config, logger *slog.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(Logger(cfg, logger))

	return r
}

// Logger はリクエストとレスポンスのログを記録するミドルウェアです。
func Logger(cfg Config, l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().In(cfg.loc)
		startStr := start.Format(cfg.timeFormatter)
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

		end := time.Now().In(cfg.loc)
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
