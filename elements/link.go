package elements

import "html"

//Link produce the html anchor
type Link struct {
	Base
	Container
	Text       string
	HTMLEncode bool
	URL        string
	NewWindow  bool
	PostBack   string
}

//String return a string with tag anchor
func (link Link) String() string {
	ret := "<a" + link.Attrs()
	if link.URL != "" {
		ret += " href='" + link.URL + "'"
	}
	if link.NewWindow {
		ret += " target='_blank'"
	}
	if link.PostBack != "" {
		if link.URL != "" {
			ret += " onclick=\"javascript:sendEvent('" + link.PostBack + "')\""
		} else {
			ret += " href=\"javascript:sendEvent('" + link.PostBack + "')\""
		}
	}
	ret += ">"
	if link.HTMLEncode {
		ret += html.EscapeString(link.Text)
	} else {
		ret += link.Text
	}
	for _, el := range link.Body {
		ret += el.String()
	}
	ret += "</a>"
	return ret
}
