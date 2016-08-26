package elements

//Element is an interface for html components
type Element interface {
	String() string
}

//Base is an type base for all elements. This means that all Elements can use the attributes listed below.
type Base struct {
	ID      string
	Classes []string
	Visible bool
	Style   map[string]string
}

//AddClass add a new class string for element
func (base *Base) AddClass(class string) {
	base.Classes = append(base.Classes, class)
}

//RemoveClass delete an class string in Classes slice
func (base *Base) RemoveClass(class string) {
	newCls := []string{}
	for _, c := range base.Classes {
		if class != c {
			newCls = append(newCls, c)
		}
	}
	base.Classes = newCls
}

//SetStyle set a style value for a key
func (base *Base) SetStyle(key string, value string) {
	base.Style[key] = value
}

//RemoveStyle delete a value in Style map
func (base *Base) RemoveStyle(key string) {
	delete(base.Style, key)
}

//Types
//Break
//Image
//Link
//EmailLink
//HorizontalRule
//
