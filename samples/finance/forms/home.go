package forms

import (
	"log"
	"time"

	"github.com/FlashBoys/go-finance"
	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
	"github.com/trumae/valente/elements"
)

//FormHome example
type FormHome struct {
	valente.FormImpl
}

func updateQuote(ws *action.WebSocket, symbol string) {
	span := elements.Span{}
	q, err := finance.GetQuote(symbol)
	if err == nil {
		val, _ := q.ChangeNominal.Float64()
		if val < 0.0 {
			span.SetStyle("color", "#f00")
			span.Text = q.LastTradePrice.String()
		} else {
			span.Text = q.LastTradePrice.String()
		}
	} else {
		span.Text = "--"
	}
	action.HTML(ws, symbol, span.String())
	log.Println(q)
}

func itemQuote(title, id string) elements.Element {
	el := elements.ListItem{}
	el.AddElement(elements.Heading2{Text: title})
	pel := elements.Paragraph{Text: "-"}
	pel.ID = id
	pel.AddClass("ui-li-aside")
	el.AddElement(pel)

	return el
}

//Render the initial html form
func (form FormHome) Render(ws *action.WebSocket, app *valente.App, params []string) error {
	root := elements.Panel{}
	root.AddElement(elements.Heading3{Text: "Quotes"})

	list := elements.List{}
	list.SetData("data-role", "listview")
	list.SetData("data-inset", "true")
	list.AddElement(itemQuote("Alphabet Inc.", "GOOG"))
	list.AddElement(itemQuote("Apple Inc.", "AAPL"))
	list.AddElement(itemQuote("Microsoft Inc.", "MSFT"))
	list.AddElement(itemQuote("Facebook Inc.", "FB"))
	root.AddElement(list)

	action.HTML(ws, "content", root.String())
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
func (form FormHome) Initialize(ws *action.WebSocket) valente.Form {
	log.Println("FormHome Initialize")

	return form
}
