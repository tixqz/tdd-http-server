package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(":5000", server)

	if err != nil {
		log.Fatalf("couldn't listen to port :5000 %v", err)
	}
}
