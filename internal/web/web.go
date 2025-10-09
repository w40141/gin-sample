package web

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w40141/gin-sample/internal/router"
)

// SetupRouter はGinのルーターを設定します。
func SetupRouter(cfg router.Config, logger *slog.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(Logger(cfg, logger))

	return r
}

// Handler は基本的なハンドラーを設定します。
func Handler(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
