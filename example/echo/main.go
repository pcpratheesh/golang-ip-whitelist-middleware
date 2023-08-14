package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	ipgaurdmiddleware "github.com/pcpratheesh/ip-gaurd-middleware/echo"
	"github.com/pcpratheesh/ip-gaurd-middleware/options"
)

func main() {
	e := echo.New()
	e.Use(ipgaurdmiddleware.IPAccessControlMiddleware(
		options.WithWhiteLists([]string{
			"*",
		}),
		options.WithFallbackHandler(func(ctx echo.Context, clientIP string) error {
			return ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Not allowed to access.",
			})
		}),
	))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Start(":8080")
}
