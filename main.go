package main

import (
	"expvar"
	"fmt"
	"github.com/mattetti/go-web-api-demo/handlers"
	"net/http"
	"time"
)

var startedAt = expvar.NewString("startedAt")
var serverBuild = expvar.NewString("build")

// Note:
// compile passing -ldflags "-X main.Build <build sha1>"
var Build string

func main() {
	startedAt.Set(time.Now().String())
	serverBuild.Set(Build)

	fmt.Println("About to start the server on port 8080")
	http.HandleFunc("/hello/", handlers.HelloWorld)
	http.ListenAndServe(":8080", nil)
}
