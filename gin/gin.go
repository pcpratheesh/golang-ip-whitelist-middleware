package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pcpratheesh/ip-gaurd-middleware/options"
	"github.com/pcpratheesh/ip-gaurd-middleware/whitelist"
)

func IPAccessControlMiddleware(opt ...options.Options) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// set the options
		for _, op := range opt {
			op()
		}

		// check the client is allowed to access the app
		clientIP := ctx.ClientIP()
		if !whitelist.CheckAllowedAccess(options.WhiteLists, clientIP) {
			if options.FallBackHandler != nil {
				options.TriggerFallbackHandler[func(ctx *gin.Context), *gin.Context](ctx)
				return
			}

			// run custom handler
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Access denied.",
			})
			ctx.Abort()
			return
		}

	}
}
