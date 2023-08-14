package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	guard "github.com/pcpratheesh/ip-guard-middleware/middleware/echo"
	"github.com/pcpratheesh/ip-guard-middleware/options"
)

var fallbackHandler = func(ctx echo.Context, clientIP string) error {
	return ctx.JSON(http.StatusForbidden, gin.H{
		"message": "Not allowed to access.",
	})
}

func main() {
	e := echo.New()
	var opts = []options.Options{
		options.WithWhiteListIPs([]string{
			"123",
		}),
		// options.SetFallbackHandler(
		// 	options.EchoFallbackHandler(fallbackHandler),
		// ),
	}

	e.Use(guard.IPAccessControlMiddleware(
		opts...,
	))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Start(":8080")
}
