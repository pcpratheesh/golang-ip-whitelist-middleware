package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pcpratheesh/ip-guard-middleware/options"
	"github.com/pcpratheesh/ip-guard-middleware/whitelist"
)

var DefaultFallbackHandler = func(ctx *gin.Context, clientIP string) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"message": fmt.Sprintf("You ip %v are not authorized to access this resource", clientIP),
	})
	ctx.Abort()
	return
}

func IPAccessControlMiddleware(opt ...options.Options) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// set the options
		for _, op := range opt {
			op()
		}

		// check the client is allowed to access the app
		clientIP := ctx.ClientIP()
		if !whitelist.CheckAllowedAccess(options.WhiteLists, clientIP) {

			// set the default fallback handler
			if options.FallBackHandler == nil {
				options.FallBackHandler = options.GinFallbackHandler(DefaultFallbackHandler)
			}

			options.FallBackHandler.(options.GinFallbackHandler)(ctx, clientIP)
			return
		}

		ctx.Next()
	}
}
