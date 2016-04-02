package forms

import (
	"log"

	"github.com/trumae/carcara/ws/action"
	"github.com/trumae/valente"
	"golang.org/x/net/websocket"
)

const htmlFormHome = `
<h3>Home</h3>
`

//FormHome example
type FormHome struct {
	valente.FormImpl
}

//Initialize inits the Home Form
func (form FormHome) Initialize(ws *websocket.Conn) valente.Form {
	log.Println("FormHome Initialize")

	action.Html(ws, "content", htmlFormHome)

	return form
}
