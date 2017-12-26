package elements

import (
	"log"
	"testing"
)

func TestRow(t *testing.T) {
	row := NewRow()
	s := row.String()

	if s != "<div class='row'></div>" {
		t.Fatal("Value not expected")
	}
}

func TestCol(t *testing.T) {
	col := NewCol()

	s := col.String()
	if s != "<div class='col'></div>" {
		t.Fatal("Value not expected")
	}

	col.S = 6
	s = col.String()
	if s != "<div class='col s6'></div>" {
		t.Fatal("Value not expected")
	}

}

func TestImageCard(t *testing.T) {
	c := NewImageCard("Title", "Image", "arrow_forward", "Text Test")

	s := c.String()
	log.Println(s)
}
