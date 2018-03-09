package main

import (
	"log"
	"net/http"

	"github.com/aymone/sample/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", api.Server()))
}
