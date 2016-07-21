package valente

import (
	"log"

	"golang.org/x/net/websocket"
)

var (
	tablews   []*websocket.Conn
	wschannel chan wsmessage
)

const (
	//GET is a tag message
	GET = iota

	//PUT is a tag message
	PUT

	//DROP is a tag message
	DROP
)

type wsmessage struct {
	Command int
	WS      *websocket.Conn
	Table   []*websocket.Conn
}

//tableWSServer handle the data in safe way
func tableWSServer() {
	for {
		msg := <-wschannel
		switch {
		case msg.Command == DROP:
			for idx, val := range tablews {
				if val == msg.WS {
					tablews[idx] = tablews[len(tablews)-1]
					tablews = tablews[:len(tablews)-1]
				}
			}

		case msg.Command == PUT:
			tablews = append(tablews, msg.WS)

		case msg.Command == GET:
			table := []*websocket.Conn{}
			copy(table, tablews)
			nmsg := wsmessage{Table: table}
			wschannel <- nmsg

		default:
			log.Println("Unknow message")
		}
	}
}

//DropWS in safe way
func DropWS(ws *websocket.Conn) {
	msg := wsmessage{
		Command: DROP,
		WS:      ws}
	wschannel <- msg
}

//PutWS in a safe way
func PutWS(ws *websocket.Conn) {
	msg := wsmessage{
		Command: PUT,
		WS:      ws}
	wschannel <- msg

}

//GetWSTable in a safe way
func GetWSTable() []*websocket.Conn {
	msg := wsmessage{
		Command: GET}

	wschannel <- msg
	ans := <-wschannel
	return ans.Table
}
