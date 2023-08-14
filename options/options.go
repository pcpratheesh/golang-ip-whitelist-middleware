package options

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

type Options func()

type (
	EchoFallbackHandler  func(ctx echo.Context, clientIP string) error
	GinFallbackHandler   func(ctx *gin.Context, clientIP string)
	FiberFallbackHandler func(ctx *fiber.Ctx, clientIP string) error
)

type FallbackHandlerInterface interface {
	EchoFallbackHandler | GinFallbackHandler | FiberFallbackHandler
}

// WhiteLists
// A list of whitelisting ip addresses
var (
	WhiteLists                 []string
	FallBackHandler            interface{}
	SuccessRedirectionCallback interface{}
)

// set the WhiteLists IPs
func WithWhiteListIPs(ips []string) Options {
	return func() {
		WhiteLists = ips
	}
}

// set the FallbackHandler
func SetFallbackHandler[T FallbackHandlerInterface](handler T) Options {
	return func() {
		FallBackHandler = handler
	}
}
