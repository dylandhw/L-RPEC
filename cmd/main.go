package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*
 * server running receives requests, forwards, then returns response
 * hardcoded for now **
 */

func main() {
	target, _ := url.Parse("https://httpbin.org/")
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.Handle("/", proxy) // takes reverseproxy obj

	fmt.Println("server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error serving", err)
	}
}
