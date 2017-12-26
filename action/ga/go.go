package ga

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/trumae/valente/action"
)

//SendGA send a event to Google Analytics
func SendGA(ws *websocket.Conn, page string) error {
	s := fmt.Sprintf("ga('send', 'pageview', '%s');", page)
	return action.Exec(ws, s)
}
