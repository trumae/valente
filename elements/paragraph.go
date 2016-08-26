package elements

import "html"

//Paragraph element produces an HTML p
type Paragraph struct {
	Base
	Text       string
	HTMLEncode bool
}

//String return the string to p tag
func (p Paragraph) String() string {
	ret := "<p" + p.Attrs() + ">"
	if p.HTMLEncode {
		ret += html.EscapeString(p.Text)
	} else {
		ret += p.Text
	}
	ret += "</p>"
	return ret
}
