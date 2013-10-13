package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	fmt.Println("checkout http://localhost:8080/debug/vars")
}

type Page struct {
	Title string
	Body  string
}

func resolveContentType(r *http.Request) string {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		return "text/html"
	} else {
		return strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
	}
}
