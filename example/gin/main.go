package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	guard "github.com/pcpratheesh/ip-guard-middleware/middleware/gin"
	"github.com/pcpratheesh/ip-guard-middleware/options"
)

var fallbackHandler = func(ctx *gin.Context, clientIP string) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"message": "Not allowed to access.",
	})
	ctx.Abort()
	return
}

func main() {
	r := gin.Default()

	var opts = []options.Options{
		options.WithWhiteLists([]string{
			"*",
		}),
		options.SetFallbackHandler(
			options.GinFallbackHandler(fallbackHandler),
		),
	}

	r.Use(guard.IPAccessControlMiddleware(
		opts...,
	))
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}
