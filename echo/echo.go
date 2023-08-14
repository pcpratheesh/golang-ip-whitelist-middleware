package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/pcpratheesh/ip-gaurd-middleware/options"
	"github.com/pcpratheesh/ip-gaurd-middleware/whitelist"
)

var DefaultFallbackHandler = func(ctx echo.Context, clientIP string) error {
	// run custom handler
	return ctx.JSON(http.StatusForbidden, gin.H{
		"message": fmt.Sprintf("You ip %v are not authorized to access this resource", clientIP),
	})
}

func IPAccessControlMiddleware(ips []string, opt ...options.Options) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			// set the ip white-lists

			// set the options
			for _, op := range opt {
				op()
			}

			// check the client is allowed to access the app
			clientIP := ctx.RealIP()
			if !whitelist.CheckAllowedAccess(options.WhiteLists, clientIP) {
				// set the default fallback handler
				if options.FallBackHandler == nil {
					options.FallBackHandler = DefaultFallbackHandler
				}

				return options.FallBackHandler.(func(ctx echo.Context, clientIP string) error)(ctx, clientIP)
			}

			return next(ctx)
		}
	}
}
