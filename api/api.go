package api

import (
	"fmt"
	"net/http"
)

// Auth ...
func Auth(token string) bool {
	if token == "password" {
		return true
	}

	return false
}

// Server will link net/http server with request handler
func Server() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/api", MainHandler)

	return r
}

// MainHandler ...
func MainHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("X-Access-Token")
	if Auth(accessToken) {
		fmt.Fprint(w, "authenticated with success.")
		return
	}

	http.Error(w, "you don't have access.", http.StatusForbidden)
}
