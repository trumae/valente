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

func params(postback []string) string {
	ret := "'" + postback[0] + "'"
	for idx, val := range postback {
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
			ret += " onclick=\"javascript:sendEvent(" + params(link.PostBack) + ")\""
		} else {
			ret += " href=\"javascript:sendEvent(" + params(link.PostBack) + ")\""
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
