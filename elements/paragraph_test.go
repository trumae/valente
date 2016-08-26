package elements

import "testing"

func TestParagraph(t *testing.T) {
	p := Paragraph{}
	s := p.String()
	exp := "<p></p>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	p = Paragraph{Text: "Text & Text Text"}
	s = p.String()
	exp = "<p>Text & Text Text</p>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	p = Paragraph{Text: "Text & Text Text", HTMLEncode: true}
	s = p.String()
	exp = "<p>Text &amp; Text Text</p>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}
}
