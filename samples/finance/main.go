package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/satori/go.uuid"
	"github.com/trumae/valente"
	"github.com/trumae/valente/samples/finance/forms"
	"github.com/trumae/valente/status"

	"golang.org/x/net/websocket"
)

//App is a Web Application representation
type App struct {
	valente.App
}

const timeout = 300
const gctime = 1

var (
	sessions map[string]*App
	mutex    sync.Mutex
)

//addiSession include a new app on sessions
func addSession(key string, app *App) {
	mutex.Lock()
	defer mutex.Unlock()

	sessions[key] = app
}

//getSession return the app by key
func getSession(key string) *App {
	return sessions[key]
}

func gcStepSession() {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().Unix()
	for key, app := range sessions {
		if now-app.LastAccess.Unix() > timeout {
			log.Println("Collecting", key)
			delete(sessions, key)
		}
	}
}

//Initialize inits the App
func (app *App) Initialize() {
	log.Println("App Initialize")

	app.AddForm("login", forms.FormLogin{})
	app.AddForm("home", forms.FormHome{})

	app.GoTo("login", nil)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("Init sessions")
	sessions = make(map[string]*App)
	mutex = sync.Mutex{}

	go func() {
		for {
			time.Sleep(gctime * time.Second)
			gcStepSession()
		}
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	e.GET("/status", status.ValenteStatusHandler)
	e.GET("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		err := websocket.Message.Send(ws, "__GETSESSION__")
		if err != nil {
			return
		}

		idSession := ""
		err = websocket.Message.Receive(ws, &idSession)
		if err != nil {
			return
		}

		var app *App
		app = getSession(idSession)
		if app == nil {
			u1 := uuid.NewV4().String()
			log.Println("New session", u1)
			app = &App{}
			app.LastAccess = time.Now()
			addSession(u1, app)
			err := websocket.Message.Send(ws, u1)
			if err != nil {
				return
			}
			app.WebSocket(ws)
			app.Initialize()
		} else {
			app.LastAccess = time.Now()
			log.Println("Reusing session", idSession)
			err := websocket.Message.Send(ws, idSession)
			if err != nil {
				return
			}
			app.WebSocket(ws)
		}
		app.Run()
	})))

	log.Println("Server running")
	e.Run(standard.New(":8000"))
}
