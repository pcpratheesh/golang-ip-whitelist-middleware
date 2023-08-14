package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pcpratheesh/ip-guard-middleware/options"
	"github.com/pcpratheesh/ip-guard-middleware/whitelist"
)

var DefaultFallbackHandler = func(ctx echo.Context, clientIP string) error {
	// run custom handler
	return ctx.JSON(http.StatusForbidden, map[string]interface{}{
		"message": fmt.Sprintf("You ip %v are not authorized to access this resource", clientIP),
	})
}

func IPAccessControlMiddleware(opt ...options.Options) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// set the options
			for _, op := range opt {
				op()
			}

			// check the client is allowed to access the app
			clientIP := ctx.RealIP()
			if !whitelist.CheckAllowedAccess(options.WhiteLists, clientIP) {
				// set the default fallback handler
				if options.FallBackHandler == nil {
					options.FallBackHandler = options.EchoFallbackHandler(DefaultFallbackHandler)
				}

				return options.FallBackHandler.(options.EchoFallbackHandler)(ctx, clientIP)
			}

			return next(ctx)
		}
	}
}
