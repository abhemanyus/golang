package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":8000", server))
}
