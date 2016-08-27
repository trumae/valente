package session

import (
	"log"
	"sync"
	"time"

	"github.com/trumae/valente"
)

const timeout = 300
const gctime = 1

var (
	sessions map[string]*valente.App
	mutex    sync.Mutex
)

//Add include a new app on sessions
func Add(key string, app *valente.App) {
	mutex.Lock()
	defer mutex.Unlock()

	sessions[key] = app
}

//Get return the app by key
func Get(key string) *valente.App {
	return sessions[key]
}

func gcstep() {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().Unix()
	for key, app := range sessions {
		if now-app.LastAccess.Unix() > gctime {
			delete(sessions, key)
		}
	}
}

func init() {
	log.Println("Init sessions")
	sessions = make(map[string]*valente.App)
	mutex = sync.Mutex{}

	go func() {
		for {
			time.Sleep(gctime * time.Second)
			gcstep()
		}
	}()
}
