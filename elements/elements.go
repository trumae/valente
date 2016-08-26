package elements

import "strings"

//Element is an interface for html components
type Element interface {
	String() string
}

//Base is an type base for all elements. This means that all Elements can use the attributes listed below.
type Base struct {
	ID      string
	Classes []string
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
	if base.Style == nil {
		base.Style = make(map[string]string)
	}
	base.Style[key] = value
}

//RemoveStyle delete a value in Style map
func (base *Base) RemoveStyle(key string) {
	delete(base.Style, key)
}

func (base Base) Attrs() string {
	var ret string

	//Classes
	if len(base.Classes) != 0 {
		ret += " class='"
		ret += strings.Join(base.Classes, " ")
		ret += "'"
	}

	//Style
	if len(base.Style) != 0 {
		ret += " style='"
		for key, val := range base.Style {
			ret += key + ":" + val + ";"
		}
		ret += "'"
	}
	return ret
}
