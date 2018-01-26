package elements

import "html"

//ListItem produces an HTML listitem element (<li>).
type ListItem struct {
	Base
	Container
	Text       string
	HTMLEncode bool
}

//String return string to li tag
func (li ListItem) String() string {
	ret := "<li" + li.Attrs() + ">"
	if li.Text != "" {
		if li.HTMLEncode {
			ret += html.EscapeString(li.Text)
		} else {
			ret += li.Text
		}
	}
	for _, el := range li.Body {
		ret += el.String()
	}
	ret += "</li>"
	return ret
}

//List element produces an HTML list element (<ol> and <ul>).
type List struct {
	Base
	Container
	Numbered bool
	Role     string
}

//String return string to ul or ol tag
func (l List) String() string {
	ret := ""
	if l.Numbered {
		ret += "<ol" + l.Attrs() + ">"
	} else {
		ret += "<ul" + l.Attrs() + ">"
	}

	for _, el := range l.Body {
		ret += el.String()
	}

	if l.Numbered {
		ret += "</ol>"
	} else {
		ret += "</ul>"
	}
	return ret
}
