package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	guard "github.com/pcpratheesh/ip-guard-middleware/middleware/fiber"
	"github.com/pcpratheesh/ip-guard-middleware/options"
)

var fallbackHandler = func(ctx *fiber.Ctx, clientIP string) error {
	return ctx.Status(http.StatusForbidden).JSON(map[string]string{
		"message": "Not allowed to access.",
	})
}

func main() {
	app := fiber.New()
	var opts = []options.Options{
		options.WithWhiteListIPs([]string{
			"127.0.0.52",
		}),
	}

	app.Use(guard.IPAccessControlMiddleware(
		opts...,
	))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})
	log.Fatal(app.Listen(":8080"))

}
