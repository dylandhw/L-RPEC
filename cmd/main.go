package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/spf13/viper"
)

/*
 * server running receives requests, forwards, then returns response
 * hardcoded for now **
 */

// small reverse proxy set up. client reaches out to the server, server goes to
// httpbin.org and gives the client the response
func main() {
	viper.AddConfigPath("/L-REPC/config/config.yaml")

	// targets http testing service
	target, _ := url.Parse("https://httpbin.org/")
	proxy := httputil.NewSingleHostReverseProxy(target)

	// handles a reverseproxy object
	http.Handle("/", proxy)

	fmt.Println("server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error serving", err)
	}
}
