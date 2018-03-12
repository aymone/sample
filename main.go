package main

import (
	"log"
	"net/http"

	"github.com/aymone/sample/api"
)

func main() {
	h := &api.AppHandler{}
	log.Fatal(http.ListenAndServe(":8080", api.Server(h)))
}
