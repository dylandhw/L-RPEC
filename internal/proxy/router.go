package internal

import (
	"strings"
)

type Route struct {
	Path     string
	Upstream string
}

// deteriministic route matching system that is based on length
func Match(routes []Route, path string) (string, bool) {
	var best Route
	for _, route := range routes {
		if strings.HasPrefix(path, route.Path) && len(route.Path) > len(best.Path) {
			best = route
		}
	}
	if best.Upstream == "" {
		return "", false
	}

	return best.Upstream, true
}
