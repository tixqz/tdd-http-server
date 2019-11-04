package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var server *PlayerServer
	store := flag.String("store", "in-memory", "players' score storage")

	flag.Parse()

	switch *store {
	case "in-memory":
		server = &PlayerServer{NewInMemoryPlayerStore()}
	case "boltdb":
		server = &PlayerServer{NewBoltDBPlayerStore()}
	default:
		log.Fatal("Wrong store type")
	}

	errServer := http.ListenAndServe(":5000", server)

	if errServer != nil {
		log.Fatalf("couldn't listen to port :5000 %v", errServer)
	}
}
