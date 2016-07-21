package forms

import (
	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
	"golang.org/x/net/websocket"
)

func applyJQM(ws *websocket.Conn) {
	action.Exec(ws, "$('#content').appendTo('.ui-page').trigger('create');")
}

func notImplementedHandle(ws *websocket.Conn, app *valente.App, params []string) {
	action.Alert(ws, "Not Implemented!")
}
