package main

import (
	"log"
	"net/http"

	"github.com/aymone/sample/api"
)

func main() {
	http.HandleFunc("/api", api.MainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
