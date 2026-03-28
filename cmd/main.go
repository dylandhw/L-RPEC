package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dylandhw/L-RPEC/internal/cache"
	"github.com/dylandhw/L-RPEC/internal/proxy"
	"github.com/dylandhw/L-RPEC/metrics"
	"github.com/spf13/viper"
)

/*
 * server running receives requests, forwards, then returns response
 * hardcoded for now **
 */

// small reverse proxy set up. client reaches out to the server, server goes to
// httpbin.org and gives the client the response
func main() {

	secretKey := os.Getenv("SECRET_KEY")
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
	cache := cache.NewCache()
	http.Handle("/", proxy.New(routes, cache, secretKey)) // need to handle requests to url

	fmt.Println("server started on port 8080")

	go func() {
		time.Sleep(5 * time.Second)
		metrics.Tests()
	}()

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error serving", err)
	}
}
