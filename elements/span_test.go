package elements

import "testing"

func TestSpan(t *testing.T) {
	span := Span{}
	s := span.String()
	if s != "<span></span>" {
		t.Error("Expected '<span></span>', got", s)
	}

	span = Span{Text: "text&"}
	s = span.String()
	if s != "<span>text&</span>" {
		t.Error("Expected '<span>text&</span>', got", s)
	}

	span = Span{Text: "text&", HTMLEncode: true}
	s = span.String()
	if s != "<span>text&amp;</span>" {
		t.Error("Expected '<span>text&amp;</span>', got", s)
	}

	span = Span{Text: "text&"}
	span.AddElement(Break{})
	span.AddElement(HorizontalRule{})
	s = span.String()
	if s != "<span>text&<br/><hr/></span>" {
		t.Error("Expected '<span>text&<br/><hr/></span>', got", s)
	}
}
