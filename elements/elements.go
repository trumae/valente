package elements

import (
	"sort"
	"strings"
)

//Element is an interface for html components
type Element interface {
	String() string
}

//Base is an type base for all elements. This means that all Elements can use the attributes listed below.
type Base struct {
	ID      string
	Classes []string
	Style   map[string]string
	Data    map[string]string
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

//SetData set a data value for a key
func (base *Base) SetData(key string, value string) {
	if base.Data == nil {
		base.Data = make(map[string]string)
	}
	base.Data[key] = value
}

//RemoveData delete a value in Data map
func (base *Base) RemoveData(key string) {
	delete(base.Data, key)
}

//Attrs return string for ID, class info and style
func (base Base) Attrs() string {
	var ret string

	if base.ID != "" {
		ret += " id='" + base.ID + "'"
	}

	//Classes
	if len(base.Classes) != 0 {
		ret += " class='"
		ret += strings.Join(base.Classes, " ")
		ret += "'"
	}

	//Style
	if len(base.Style) != 0 {
		ret += " style='"
		var keys []string
		for key, _ := range base.Style {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, val := range keys {
			ret += val + ":" + base.Style[val] + ";"
		}
		ret += "'"
	}

	for dat, dval := range base.Data {
		ret += " " + dat + "='" + dval + "'"
	}

	return ret
}

//Container is an base to elements with body
type Container struct {
	Body []Element
}

//AddElement put a new element on Body
func (c *Container) AddElement(el Element) {
	c.Body = append(c.Body, el)
}
