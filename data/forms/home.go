package forms

import (
	"log"

	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
)

const htmlFormHome = `
<h3>Home</h3>
`

//FormHome example
type FormHome struct {
	valente.FormImpl
}

//Initialize inits the Home Form
func (form FormHome) Initialize(ws *action.WebSocket) valente.Form {
	log.Println("FormHome Initialize")

	action.HTML(ws, "content", htmlFormHome)

	return form
}
