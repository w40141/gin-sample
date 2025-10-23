package router

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouterGin はGinのルーターを設定します。
func SetupRouterGin(logger *slog.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(Logger(logger))

	r = handler(r)

	return r
}

// handler は基本的なハンドラーを設定します。
func handler(r *gin.Engine) *gin.Engine {
	r.GET("/checkhealth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	return r
}
