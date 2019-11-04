package main

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	errServer := http.ListenAndServe(":5000", server)

	if errServer != nil {
		log.Fatalf("couldn't listen to port :5000 %v", errServer)
	}
}
