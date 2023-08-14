package options

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type Options func()

type FallbackHandlerInterface interface {
	func(ctx *gin.Context, clientIP string) | func(ctx echo.Context, clientIP string) error
}

// WhiteLists
// A list of whitelisting ip addresses
var (
	WhiteLists                 []string
	FallBackHandler            interface{}
	SuccessRedirectionCallback interface{}
)

// set the WhiteLists
func WithWhiteLists(ips []string) Options {
	return func() {
		WhiteLists = ips
	}
}

// set the FallbackHandler
func WithFallbackHandler[T FallbackHandlerInterface](handler T) Options {
	return func() {
		FallBackHandler = handler
	}
}
