package forms

import (
	"log"
	"strconv"
	"time"

	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
	"github.com/trumae/valente/elements"
)

//FormHome example
type FormHome struct {
	valente.FormImpl
}

//Render the initial html form
func (form FormHome) Render(ws *action.WebSocket, app *valente.App, params []string) error {
	root := elements.Panel{}
	root.AddElement(elements.Heading1{Text: "H1 - Count"})

	root.AddElement(elements.Heading3{Text: "Update"})
	count := elements.Panel{}
	count.ID = "count"
	root.AddElement(count)

	root.AddElement(elements.Heading3{Text: "HorizontalRule"})
	root.AddElement(elements.HorizontalRule{})

	root.AddElement(elements.Heading3{Text: "Image"})
	image1 := elements.Image{
		Source: "http://i3.cpcache.com/product/34276024/valente_family_crest_tile_coaster.jpg?height=150&width=150",
	}
	image2 := elements.Image{
		Source: "http://www.filmeb.com.br/sites/default/files/calendario-de-estreias/fotos/Filme-Valente-Cartaz-1.jpg",
	}
	root.AddElement(image1)
	root.AddElement(image2)

	root.AddElement(elements.Heading3{Text: "Link"})
	link := elements.Link{
		Text:      "click here - (new window)",
		URL:       "https://github.com/trumae/valente",
		NewWindow: true,
	}
	link2 := elements.Link{
		Text: "click here",
		URL:  "https://github.com/trumae/valente",
	}
	root.AddElement(link)
	root.AddElement(elements.Break{})
	root.AddElement(link2)

	root.AddElement(elements.Heading3{Text: "List"})
	list := elements.List{}
	list.AddElement(elements.ListItem{Text: "item1"})
	list.AddElement(elements.ListItem{Text: "item2"})
	list.AddElement(elements.ListItem{Text: "item3"})
	root.AddElement(list)

	root.AddElement(elements.Heading3{Text: "Button"})
	btn := elements.Button{Text: "valente"}
	root.AddElement(btn)

	root.AddElement(elements.Heading3{Text: "InputText"})
	it := elements.InputText{Value: "Valente"}
	root.AddElement(it)

	root.AddElement(elements.Heading3{Text: "InputPassword"})
	pa := elements.InputPassword{}
	root.AddElement(pa)

	root.AddElement(elements.Heading3{Text: "TextArea"})
	ta := elements.TextArea{Text: "Valente"}
	root.AddElement(ta)

	root.AddElement(elements.Heading3{Text: "Sparkline"})
	spark := elements.Sparkline{}
	spark.ID = "sparkline"
	spark.Values = []int{1, 2, 3, 4, 5, 3, 2, 3, 4, 3}
	root.AddElement(spark)

	root.AddElement(elements.Break{})
	sparkBar := elements.Sparkline{Options: map[string]string{"type": "bar"}}
	sparkBar.ID = "sparklineBar"
	sparkBar.Values = []int{1, 2, 3, 4, 5, 3, 2, 3, 4, 3}
	root.AddElement(sparkBar)

	action.HTML(ws, "content", root.String())

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
	return nil
}

//Initialize inits the Home Form
func (form FormHome) Initialize(ws *action.WebSocket) valente.Form {
	log.Println("FormHome Initialize")

	return form
}
