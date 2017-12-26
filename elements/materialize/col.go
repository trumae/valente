package elements

import (
	"fmt"

	valente "github.com/trumae/valente/elements"
)

//Col element of layout
type Col struct {
	div valente.Panel
	S   int
	M   int
	L   int
	XL  int
}

//NewCol create a new column for layout
func NewCol() *Col {
	col := &Col{}

	col.div.AddClass("col")

	return col
}

func (col *Col) String() string {
	if col.S != 0 {
		c := fmt.Sprintf("s%d", col.S)
		col.div.AddClass(c)
	}

	if col.M != 0 {
		c := fmt.Sprintf("m%d", col.M)
		col.div.AddClass(c)
	}

	if col.L != 0 {
		c := fmt.Sprintf("l%d", col.L)
		col.div.AddClass(c)
	}
	if col.XL != 0 {
		c := fmt.Sprintf("xl%d", col.XL)
		col.div.AddClass(c)
	}

	return col.div.String()
}
