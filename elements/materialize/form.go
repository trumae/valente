package elements

import "github.com/trumae/valente/elements"

//Form element produces an HTML form element (<form>)
type Form struct {
	elements.Base
	elements.Container
}

//String emit html text for Heading1
func (f Form) String() string {
	ret := "<form" + f.Attrs() + ">"

	for _, el := range f.Body {
		ret += el.String()
	}

	ret += "</form>"
	return ret
}

//Label element produces an HTML label element (<label>)
type Label struct {
	elements.Base
	For  string
	Text string
}

//String emit html text for Heading1
func (l Label) String() string {
	ret := "<label" + l.Attrs() + " for='" + l.For + "'>"
	ret += l.Text
	ret += "</label>"
	return ret
}

//NewInputTextRow return an input text element to use on form
func NewInputTextRow(id, label, value, itype, iclass string) *Row {
	row := NewRow()
	div := elements.Panel{}
	div.AddClass("input-field col s12")

	input := elements.InputText{}
	input.ID = id
	input.AddClass("validate")
	input.AddClass(iclass)
	input.SetData("placeholder", label)
	input.SetData("value", value)
	input.SetData("type", itype)
	input.SetData("name", id)
	lab := Label{For: input.ID,
		Text: label}

	div.AddElement(input)
	div.AddElement(lab)
	row.AddElement(div)

	return row
}

//NewInputCodBarRow return an input text element to use on form
func NewInputCodBarRow(id, label, value string) *Row {
	row := NewRow()
	div := elements.Panel{}
	div.AddClass("input-field col s8")

	input := elements.InputText{}
	input.ID = id
	input.AddClass("validate")
	input.SetData("placeholder", label)
	input.SetData("value", value)
	input.SetData("name", id)
	lab := Label{For: input.ID,
		Text: label}

	btn := elements.Link{Text: "CodBar",
		URL: "javascript:Coletor.startLeitorCodigoBarras()"}
	btn.AddClass("waves-effect")
	btn.AddClass("waves-light")
	btn.AddClass("btn")

	div.AddElement(input)
	div.AddElement(lab)
	row.AddElement(div)
	row.AddElement(btn)

	return row
}

//NewInputPasswordRow return an input text element to use on form
func NewInputPasswordRow(id, label, value string) *Row {
	row := NewRow()
	div := elements.Panel{}
	div.AddClass("input-field col s12")

	input := elements.InputPassword{}
	input.ID = id
	input.AddClass("validate")
	input.SetData("placeholder", label)
	input.SetData("value", value)
	input.SetData("name", id)
	lab := Label{For: input.ID,
		Text: label}

	div.AddElement(input)
	div.AddElement(lab)
	row.AddElement(div)

	return row
}
