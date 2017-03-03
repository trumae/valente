package elements

import "html"

//Button element produces an HTML button
type Button struct {
	Base
	Text       string
	HTMLEncode bool
	PostBack   []string
}

//String return a string with button
func (btn Button) String() string {
	ret := "<button" + btn.Attrs()
	if len(btn.PostBack) > 0 {
		ret += " onclick=\"sendEvent(" + params(btn.PostBack) + ")\""
	}
	ret += ">"
	if btn.HTMLEncode {
		ret += html.EscapeString(btn.Text)
	} else {
		ret += btn.Text
	}
	ret += "</button>"
	return ret
}

//InputText element produces an HTML text input
type InputText struct {
	Base
	Value string
}

//String return a string with input text
func (text InputText) String() string {
	ret := "<input type='text' "
	ret += "value='" + html.EscapeString(text.Value) + "'"
	ret += text.Attrs() + "/>"
	return ret
}

//InputPassword element produces an HTML password input
type InputPassword struct {
	Base
}

//String return a string with input password
func (pass InputPassword) String() string {
	ret := "<input type='password'"
	ret += pass.Attrs() + "/>"
	return ret
}

//TextArea element produces an HTML textarea
type TextArea struct {
	Base
	Text string
}

//String return the string to textarea tag
func (ta TextArea) String() string {
	ret := "<textarea" + ta.Attrs() + ">"
	ret += html.EscapeString(ta.Text)
	ret += "</textarea>"
	return ret
}
