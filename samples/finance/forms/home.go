package forms

import (
	"log"
	"time"

	"github.com/FlashBoys/go-finance"
	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
	"golang.org/x/net/websocket"
)

const htmlFormHome = `
<h3>Quotes</h3>
<ul data-role="listview" data-inset="true">
  <li>
    <h2>Alphabet Inc.</h2>
		<p class="ui-li-aside"><strong id="GOOG"></strong></p>
  </li>
  <li>
    <h2>Apple Inc.</h2>
		<p class="ui-li-aside"><strong id="AAPL"></strong></p>
  </li>
  <li>
    <h2>Microsoft Inc.</h2>
    <p class="ui-li-aside"><strong id="MSFT"></strong></p>
  </li>
  <li>
    <h2>Facebook, Inc.</h2>
    <p class="ui-li-aside"><strong id="FB"></strong></p>
  </li>
</ul>
`

//FormHome example
type FormHome struct {
	valente.FormImpl
}

func updateQuote(ws *websocket.Conn, symbol string) {
	q, err := finance.GetQuote(symbol)
	if err == nil {
		val, _ := q.ChangeNominal.Float64()
		if val < 0.0 {
			action.HTML(ws, symbol, "<span style='color:#f00;'>"+q.LastTradePrice.String()+"</span>")
		} else {
			action.HTML(ws, symbol, q.LastTradePrice.String())
		}
	} else {
		action.HTML(ws, symbol, "--")
	}
	log.Println(q)
}

//Render the initial html form
func (form FormHome) Render(ws *websocket.Conn, app *valente.App, params []string) error {
	action.HTML(ws, "content", htmlFormHome)
	action.Exec(ws, "$('#content').appendTo('.ui-page').trigger('create');")

	go func() {
		for {
			updateQuote(ws, "GOOG")
			updateQuote(ws, "AAPL")
			updateQuote(ws, "MSFT")
			updateQuote(ws, "FB")
			time.Sleep(5 * time.Minute)
		}
	}()

	return nil
}

//Initialize inits the Home Form
func (form FormHome) Initialize(ws *websocket.Conn) valente.Form {
	log.Println("FormHome Initialize")

	return form
}
