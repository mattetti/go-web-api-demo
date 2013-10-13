package main

import (
	"expvar"
	"fmt"
	"github.com/mattetti/go-web-api-demo/handlers"
	"net/http"
	"time"
  "os"
  "runtime"
)

var startedAt = expvar.NewString("startedAt")
var serverPid = expvar.NewInt("pid")
var serverBuild = expvar.NewString("build")

// Note:
// compile passing -ldflags "-X main.Build <build sha1>"
var Build string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	serverPid.Set(int64(os.Getpid()))
}

func main() {
	startedAt.Set(time.Now().String())
	serverBuild.Set(Build)

	fmt.Println("About to start the server on port 8080")
	http.HandleFunc("/hello/", handlers.HelloWorld)
  http.HandleFunc("/debug", handlers.DebugVars)
	http.ListenAndServe(":8080", nil)
}
