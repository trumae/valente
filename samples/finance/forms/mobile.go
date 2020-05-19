package forms

import (
	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
)

func applyJQM(ws *action.WebSocket) {
	action.Exec(ws, "$('#content').appendTo('.ui-page').trigger('create');")
}

func notImplementedHandle(ws *action.WebSocket, app *valente.App, params []string) {
	action.Alert(ws, "Not Implemented!")
}
