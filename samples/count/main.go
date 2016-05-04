package main

import (
	"log"
	"runtime"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/trumae/valente"
	"github.com/trumae/valente/samples/count/forms"

	"golang.org/x/net/websocket"
)

//App is a Web Application representation
type App struct {
	valente.App
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	e.Get("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		app := App{}
		app.WS = ws
		app.Initialize()
		app.Run()
	})))

	log.Println("Server running")
	e.Run(standard.New(":8000"))
}
