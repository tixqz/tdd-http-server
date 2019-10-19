package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(":5000", server)

	if err != nil {
		log.Fatalf("couldn't listen to port :5000 %v", err)
	}
}
