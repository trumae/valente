package elements

//HorizontalRule produce the hr tag
type HorizontalRule struct {
	Base
}

//String implements element interface into HorizontalRule
func (hr HorizontalRule) String() string {
	return "<hr" + hr.Attrs() + "/>"
}
