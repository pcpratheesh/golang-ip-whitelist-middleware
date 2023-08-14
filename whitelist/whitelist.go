package whitelist

import (
	"slices"

	"github.com/pcpratheesh/ip-guard-middleware/options"
)

func CheckAllowedAccess(ips []string, clientIP string) bool {
	// check accessible for all IP's
	if slices.Contains(options.WhiteLists, "*") {
		return true
	}

	return slices.Contains(ips, clientIP)
}
