package elements

import "html"

//Span element produces an HTML span
type Span struct {
	Base
	Container
	Text       string
	HTMLEncode bool
}

//String emit text for span
func (span Span) String() string {
	ret := "<span" + span.Attrs() + ">"
	if span.HTMLEncode {
		ret += html.EscapeString(span.Text)
	} else {
		ret += span.Text
	}
	for _, el := range span.Body {
		ret += el.String()
	}
	ret += "</span>"
	return ret
}
