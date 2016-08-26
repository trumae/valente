package elements

import "testing"

func TestPanel(t *testing.T) {
	panel := Panel{}
	s := panel.String()
	if s != "<div></div>" {
		t.Error("Expected '<div></div>', got", s)
	}

	panel = Panel{Text: "text&"}
	s = panel.String()
	if s != "<div>text&</div>" {
		t.Error("Expected '<div>text&</div>', got", s)
	}

	panel = Panel{Text: "text&", HTMLEncode: true}
	s = panel.String()
	if s != "<div>text&amp;</div>" {
		t.Error("Expected '<div>text&amp;</div>', got", s)
	}

	panel = Panel{Text: "text&"}
	panel.AddElement(Break{})
	panel.AddElement(HorizontalRule{})
	s = panel.String()
	if s != "<div>text&<br/><hr/></div>" {
		t.Error("Expected '<div>text&<br/><hr/></div>', got", s)
	}
}
