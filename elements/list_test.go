package elements

import "testing"

func TestListItem(t *testing.T) {
	li := ListItem{}
	s := li.String()
	exp := "<li></li>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	li = ListItem{Text: "item1 &"}
	s = li.String()
	exp = "<li>item1 &</li>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	li = ListItem{Text: "item1 &", HTMLEncode: true}
	s = li.String()
	exp = "<li>item1 &amp;</li>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	li = ListItem{Text: "item1"}
	li.AddClass("class1")
	s = li.String()
	exp = "<li class='class1'>item1</li>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	li = ListItem{}
	li.AddElement(Heading1{Text: "title"})
	s = li.String()
	exp = "<li><h1>title</h1></li>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}
}

func TestList(t *testing.T) {
	l := List{}
	s := l.String()
	exp := "<ul></ul>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	l = List{Numbered: true}
	s = l.String()
	exp = "<ol></ol>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}

	l = List{}
	l.AddElement(ListItem{Text: "item1"})
	l.AddElement(ListItem{Text: "item2"})
	l.AddElement(ListItem{Text: "item3"})
	s = l.String()
	exp = "<ul><li>item1</li><li>item2</li><li>item3</li></ul>"
	if s != exp {
		t.Error("Expected", exp, "got", s)
	}
}
