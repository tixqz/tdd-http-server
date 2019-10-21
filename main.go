package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(":5000", server)

	if err != nil {
		log.Fatalf("couldn't listen to port :5000 %v", err)
	}
}
