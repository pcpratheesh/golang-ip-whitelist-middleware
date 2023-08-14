package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ipgaurdmiddleware "github.com/pcpratheesh/ip-gaurd-middleware/gin"
	"github.com/pcpratheesh/ip-gaurd-middleware/options"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(ipgaurdmiddleware.IPAccessControlMiddleware(
		options.WithWhiteLists([]string{
			"187547",
		}),
		options.WithFallbackHandler(func(ctx *gin.Context, clientIP string) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Not allowed to access.",
			})
			ctx.Abort()
			return
		}),
	))
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}
