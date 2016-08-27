package main

import (
	"log"
	"sync"

	"github.com/trumae/valente"
)

var (
	sessions map[string]valente.App
	mutex    sync.Mutex
)

func init() {
	log.Println("Init sessions")
	sessions = make(map[string]valente.App)
	mutex = sync.Mutex{}

}
