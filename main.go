package main

import (
	"log"
)

//Underlying storage (in memory, on disk, s3, anything
//server (http, tcp)
func main() {
	cfg := &Config{
		ListenAddr: ":3000",
		StoreProducerFunc: func() Storer {
			return NewMemoryStore()
		},
	}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	s.Start()

}
