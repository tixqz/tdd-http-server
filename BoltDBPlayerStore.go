package main

import (
	"github.com/boltdb/bolt"
	"log"
	"strconv"
)

type BoltDBPlayerStore struct {
	db *bolt.DB
}

func NewBoltDBPlayerStore() *BoltDBPlayerStore {
	db, err := bolt.Open("players.db", 0600, nil)
	if err != nil {
		log.Fatal("Couldn't connect to DB")
		return nil
	}
	defer db.Close()

	return &BoltDBPlayerStore{db}
}

func (b *BoltDBPlayerStore) GetPlayerScore(name string) int {
	var score int
	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PlayerScore"))
		v := b.Get([]byte(name))
		score, _ = strconv.Atoi(string(v))
		return nil
	})

	return score
}

func (b *BoltDBPlayerStore) RecordWin(name string) {

}
