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
	PostBack   []string
}

func (l *Link) linkParams() string {
	ret := "'" + l.PostBack[0] + "'"
	for idx, val := range l.PostBack {
		if idx == 0 {
			continue
		}
		ret += ",'" + val + "'"
	}
	return ret
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
	if len(link.PostBack) > 0 {
		if link.URL != "" {
			ret += " onclick=\"javascript:sendEvent(" + link.linkParams() + ")\""
		} else {
			ret += " href=\"javascript:sendEvent(" + link.linkParams() + ")\""
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
