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
func Server(h Handler) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/api", h.MainHandler)

	return r
}

// Handler interface
type Handler interface {
	MainHandler(http.ResponseWriter, *http.Request)
}

// AppHandler will handle api
type AppHandler struct{}

// MainHandler ...
func (h *AppHandler) MainHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("X-Access-Token")
	if Auth(accessToken) {
		fmt.Fprint(w, "authenticated with success.")
		return
	}

	http.Error(w, "you don't have access.", http.StatusForbidden)
}
