package elements

import "testing"

func TestLink(t *testing.T) {
	link := Link{}
	s := link.String()
	if s != "<a></a>" {
		t.Error("Expected '<a></a>', got", s)
	}

	link = Link{Text: "valente"}
	s = link.String()
	if s != "<a>valente</a>" {
		t.Error("Expected '<a>valente</a>', got", s)
	}

	link = Link{Text: "valente", URL: "http://github.com/trumae/valente"}
	s = link.String()
	if s != "<a href='http://github.com/trumae/valente'>valente</a>" {
		t.Error("Expected '<a href='http://github.com/trumae/valente'>text</a>', got", s)
	}

	link = Link{Text: "valente", URL: "http://github.com/trumae/valente"}
	link.AddClass("class1")
	s = link.String()
	if s != "<a class='class1' href='http://github.com/trumae/valente'>valente</a>" {
		t.Error("Expected '<a class='class1' href='http://github.com/trumae/valente'>valente</a>', got", s)
	}

	link = Link{Text: "valente", URL: "http://github.com/trumae/valente", NewWindow: true}
	link.AddClass("class1")
	s = link.String()
	if s != "<a class='class1' href='http://github.com/trumae/valente' target='_blank'>valente</a>" {
		t.Error("Expected '<a class='class1' href='http://github.com/trumae/valente' target='_blank'>valente</a>', got", s)
	}

	link = Link{Text: "valente", URL: "http://github.com/trumae/valente", NewWindow: true, PostBack: []string{"back"}}
	link.AddClass("class1")
	s = link.String()
	if s != "<a class='class1' href='http://github.com/trumae/valente' target='_blank' onclick=\"javascript:sendEvent('back')\">valente</a>" {
		t.Error("Expected '<a class='class1' href='http://github.com/trumae/valente' target='_blank' onclick=\"javascript:sendEvent('back')\">valente</a>', got", s)
	}

	link = Link{Text: "valente", PostBack: []string{"back"}}
	s = link.String()
	if s != "<a href=\"javascript:sendEvent('back')\">valente</a>" {
		t.Error("Expected '<a href=\"javascript:sendEvent('back')\">valente</a>', got", s)
	}

	link = Link{URL: "http://github.com/trumae/valente", NewWindow: true, PostBack: []string{"back"}}
	link.AddClass("class1")
	link.AddElement(Break{})
	link.AddElement(HorizontalRule{})
	s = link.String()
	if s != "<a class='class1' href='http://github.com/trumae/valente' target='_blank' onclick=\"javascript:sendEvent('back')\"><br/><hr/></a>" {
		t.Error("Expected '<a class='class1' href='http://github.com/trumae/valente' target='_blank' onclick=\"javascript:sendEvent('back')\"><br/><hr/></a>', got", s)
	}

	link = Link{Text: "valente", PostBack: []string{"back", "two"}}
	s = link.String()
	if s != "<a href=\"javascript:sendEvent('back','two')\">valente</a>" {
		t.Error("Expected '<a href=\"javascript:sendEvent('back','two')\">valente</a>', got", s)
	}
}
