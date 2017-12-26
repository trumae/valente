package elements

import (
	valente "github.com/trumae/valente/elements"
)

//Card is a materializecss component
type Card struct {
	valente.Panel
}

//NewCard create and initialize a new card
func NewCard(title string, content *valente.Panel, actions []*valente.Link) *Card {
	card := &Card{}

	d := valente.Panel{}

	card.AddClass("card")
	d.AddClass("card-content")

	stitle := valente.Span{Text: title}
	stitle.AddClass("card-title")

	d.AddElement(stitle)
	d.AddElement(content)

	card.AddElement(d)

	if len(actions) > 0 {
		caction := valente.Panel{}
		caction.AddClass("card-action")

		for _, a := range actions {
			caction.AddElement(*a)
		}
		card.AddElement(caction)
	}

	return card
}
