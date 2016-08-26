package elements

import "testing"

func TestHeading1(t *testing.T) {
	h1 := Heading1{Text: "heading1"}
	s := h1.String()
	if s != "<h1>heading1</h1>" {
		t.Error("Expected '<h1>heading1</h1>', got ", s)
	}
}

func TestHeading2(t *testing.T) {
	h2 := Heading2{Text: "heading2"}
	s := h2.String()
	if s != "<h2>heading2</h2>" {
		t.Error("Expected '<h2>heading2</h2>', got ", s)
	}
}

func TestHeading3(t *testing.T) {
	h3 := Heading3{Text: "heading3"}
	s := h3.String()
	if s != "<h3>heading3</h3>" {
		t.Error("Expected '<h3>heading3</h3>', got ", s)
	}
}

func TestHeading4(t *testing.T) {
	h4 := Heading4{Text: "heading4"}
	s := h4.String()
	if s != "<h4>heading4</h4>" {
		t.Error("Expected '<h4>heading4</h4>', got ", s)
	}
}
