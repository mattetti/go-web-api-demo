package handlers

import (
	"encoding/json"
	"fmt"
	html "html/template"
	"net/http"
)

var helloTemplate *html.Template

func init() {
	fmt.Println("Try: /hello/world")
	var err error
	helloTemplate, err = html.ParseFiles("templates/hello.html")
	if err != nil {
		panic(err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// write the response
	//fmt.Fprintf(w, "Hello %s!", r.URL.Path[len("/hello/"):])
	page := Page{Title: r.URL.Path[len("/hello/"):], Body: "This is a test"}
	contentType := resolveContentType(r)
	if contentType == "text/html" {
		helloTemplate.Execute(w, page)
	} else {
		// json
		var b []byte
		var err error
		w.Header().Set("Content-Type", "application/json")
		b, err = json.Marshal(page)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}
