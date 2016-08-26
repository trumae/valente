package elements

import "testing"

func TestBreak(t *testing.T) {
	br := Break{}
	s := br.String()
	if s != "<br/>" {
		t.Error("Expected '<br/>'")
	}
}

func TestBreakElement(t *testing.T) {
	var br Element

	br = Break{}
	s := br.String()
	if s != "<br/>" {
		t.Error("Expected '<br/>'")
	}
}
