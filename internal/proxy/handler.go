package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/dylandhw/L-RPEC/internal/cache"
)

type Handler struct {
	routes []Route
	cache  *cache.Cache
}

func New(routes []Route, cache *cache.Cache) *Handler {
	return &Handler{
		routes: routes,
		cache:  cache,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + r.URL.Path
	fmt.Println("THIS IS THE KEY: ", key)
	upstream, ok := Match(h.routes, r.URL.Path)
	if !ok {
		http.Error(w, "no matching route", http.StatusBadGateway)
		return
	}

	target, _ := url.Parse(upstream)
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println("PROXY ERROR: ", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	fmt.Println("forwarding to:", target.String(), r.URL.Path)
	proxy.ServeHTTP(w, r)
}
