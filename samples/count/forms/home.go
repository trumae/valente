package forms

import (
	"log"
	"strconv"
	"time"

	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
	"golang.org/x/net/websocket"
)

const htmlFormHome = `
<h3>Home</h3>
  <span id="count"></span>
`

//FormHome example
type FormHome struct {
	valente.FormImpl
}

//Initialize inits the Home Form
func (form FormHome) Initialize(ws *websocket.Conn) valente.Form {
	log.Println("FormHome Initialize")

	action.HTML(ws, "content", htmlFormHome)

	go func() {
		i := 0
		c := time.Tick(1 * time.Second)
		for range c {
			i = i + 1
			err := action.HTML(ws, "count", strconv.Itoa(i))
			if err != nil {
				log.Println("Error sending count ", err)
				return
			}
		}
	}()

	return form
}
