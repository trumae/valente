package elements

import (
	valente "github.com/trumae/valente/elements"
)

//Row is a element of layout
type Row struct {
	valente.Panel
}

//NewRow create a new Row
func NewRow() *Row {
	row := &Row{}

	row.AddClass("row")

	return row
}
