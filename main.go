package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	err := http.ListenAndServe(":5000", handler)

	if err != nil {
		log.Fatalf("couldn't listen to port :5000 %v", err)
	}
}
