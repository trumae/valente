package elements

import "testing"

func TestHorizontalRule(t *testing.T) {
	hr := HorizontalRule{}
	s := hr.String()
	if s != "<hr/>" {
		t.Error("Expected '<hr/>'")
	}
}
