package fiber

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pcpratheesh/ip-guard-middleware/options"
	"github.com/pcpratheesh/ip-guard-middleware/whitelist"
)

var DefaultFallbackHandler = func(ctx *fiber.Ctx, clientIP string) error {
	// run custom handler
	return ctx.Status(http.StatusForbidden).JSON(map[string]interface{}{
		"message": fmt.Sprintf("You ip %v are not authorized to access this resource", clientIP),
	})
}

func IPAccessControlMiddleware(opt ...options.Options) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// set the options
		for _, op := range opt {
			op()
		}

		// check the client is allowed to access the app
		clientIP := ctx.IP()
		if !whitelist.CheckAllowedAccess(options.WhiteLists, clientIP) {
			// set the default fallback handler
			if options.FallBackHandler == nil {
				options.FallBackHandler = options.FiberFallbackHandler(DefaultFallbackHandler)
			}

			return options.FallBackHandler.(options.FiberFallbackHandler)(ctx, clientIP)
		}

		return ctx.Next()
	}
}
