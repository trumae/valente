package elements

import "testing"

func TestImage(t *testing.T) {
	img := Image{}
	s := img.String()
	if s != "<img/>" {
		t.Error("Expected '<img/>', got", s)
	}

	img = Image{Source: "img.jpg"}
	s = img.String()
	if s != "<img src='img.jpg'/>" {
		t.Error("Expected '<img src='img.jpg'/>', got", s)
	}

	img = Image{Source: "img.jpg"}
	img.AddClass("big")
	s = img.String()
	if s != "<img class='big' src='img.jpg'/>" {
		t.Error("Expected '<img class='big' src='img.jpg'/>', got", s)
	}
}
