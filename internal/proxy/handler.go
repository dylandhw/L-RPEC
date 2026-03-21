package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Handler struct {
	routes []Route
}

func New(routes []Route) *Handler {
	return &Handler{routes: routes}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upstream, ok := Match(h.routes, r.URL.Path)
	if !ok {
		http.Error(w, "no matching route", http.StatusBadGateway)
		return
	}

	target, _ := url.Parse(upstream)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
