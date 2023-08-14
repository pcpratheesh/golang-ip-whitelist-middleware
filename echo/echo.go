package gin

// func IPAccessControlMiddleware(opt ...options.Options) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(ctx echo.Context) error {
// 			// set the options
// 			for _, op := range opt {
// 				op()
// 			}

// 			// check the client is allowed to access the app
// 			clientIP := ctx.ClientIP()
// 			if !whitelist.CheckAllowedAccess(options.WhiteLists, clientIP) {
// 				if options.FallBackHandler != nil {
// 					return options.TriggerFallbackHandler[func(ctx echo.Context) error, echo.Context](ctx)
// 				}

// 				// run custom handler
// 				ctx.JSON(http.StatusForbidden, gin.H{
// 					"message": "Access denied.",
// 				})
// 				ctx.Abort()
// 				return
// 			}
// 		}
// 	}
// }
