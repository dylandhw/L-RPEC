package main

import (
	"fmt"
	"net/http"
)

/*
 * server running receives requests, forwards, then returns response
 * hardcoded for now **
 */

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "requested url: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error serving", err)
	}
}
