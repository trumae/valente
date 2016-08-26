package elements

import "html"

//Panel element produces an HTML div
type Panel struct {
	Base
	Container
	Text       string
	HTMLEncode bool
}

//String emit text for Panel
func (panel Panel) String() string {
	ret := "<div" + panel.Attrs() + ">"
	if panel.HTMLEncode {
		ret += html.EscapeString(panel.Text)
	} else {
		ret += panel.Text
	}
	for _, el := range panel.Body {
		ret += el.String()
	}
	ret += "</div>"
	return ret
}
