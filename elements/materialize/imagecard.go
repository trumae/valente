package elements

import (
	valente "github.com/trumae/valente/elements"
)

//ImageCard is a materializecss component
type ImageCard struct {
	valente.Panel
	image   valente.Panel
	Content valente.Panel
	Actions valente.Panel
}

//NewImageCard create and initialize a new card - put icon blank to without icon
func NewImageCard(title, fimg, icon, textcontent string) ImageCard {
	card := ImageCard{}

	card.AddClass("card")

	card.image.AddClass("card-image")

	img := valente.Image{Source: fimg}
	card.image.AddElement(img)

	stitle := valente.Span{Text: title}
	stitle.AddClass("card-title")
	stitle.SetStyle("background-color", "#22222299")
	stitle.SetStyle("padding", "3px")

	card.image.AddElement(stitle)

	if len(icon) != 0 {
		aicon := valente.Link{Text: "<i class='material-icons'>" + icon + "</i>"}
		aicon.AddClass("btn-floating halfway-fab waves-effect waves-light red")
		card.image.AddElement(aicon)
	}

	card.Content.AddClass("card-content")
	card.Content.Text = "<p>" + textcontent + "</p>"

	card.AddElement(card.image)
	card.AddElement(card.Content)

	return card
}
