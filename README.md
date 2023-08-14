# ip-guard-middleware
IP access control middleware for Gin, echo and fiber web frameworks. This middleware enhances your application's security by allowing you to define a whitelist of permitted IP addresses while providing extended features for handling non-whitelisted IPs.

## Installation
To install GoHealthWatch, use the following command:
```
go get github.com/pcpratheesh/ip-guard-middleware/
```

## Usage
### For Gin

*Import the package*
```go
    import (
        guard "github.com/pcpratheesh/ip-guard-middleware/middleware/gin"
    )
```

*Initializing the options*
```go
var opts = []options.Options{
    options.WithWhiteListIPs([]string{
        "your-ip",
    }),
}

```

*Custom fallback handler function*
```go
var fallbackHandler = func(ctx *gin.Context, clientIP string) {
	ctx.JSON(http.StatusForbidden, map[string]string{
		"message": "Not allowed to access.",
	})
	ctx.Abort()
	return
}

...

options.SetFallbackHandler(
    options.GinFallbackHandler(fallbackHandler),
)
```

## Examples
- [Gin](example/gin/main.go)
- [Echo](example/echo/main.go)
- [Fiber](example/fiber/main.go)