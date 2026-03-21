package main

import (
	"fmt"
	"net/http"

	"github.com/dylandhw/L-RPEC/internal/proxy"
	"github.com/spf13/viper"
)

/*
 * server running receives requests, forwards, then returns response
 * hardcoded for now **
 */

// small reverse proxy set up. client reaches out to the server, server goes to
// httpbin.org and gives the client the response
func main() {

	// viper setup
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err reading config file: %s\n", err)
	}
	fmt.Printf("config settings: \n%+v\n", viper.AllSettings())

	var routes []proxy.Route
	viper.UnmarshalKey("routes", &routes)
	for route := range routes {
		fmt.Println("route: ", route)

	}

	http.Handle("/", proxy.New(routes)) // need to handle requests to url

	fmt.Println("server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error serving", err)
	}
}
