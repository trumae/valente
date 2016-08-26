package elements

import "testing"

func TestBreak(t *testing.T) {
	br := Break{}
	s := br.String()
	if s != "<br/>" {
		t.Error("Expected '<br/>', got ", s)
	}

	br = Break{}
	br.AddClass("class1")
	s = br.String()
	if s != "<br class='class1'/>" {
		t.Error("Expected '<br class='class1'/>', got", s)
	}

	br = Break{}
	br.AddClass("class1")
	br.AddClass("class2")
	s = br.String()
	if s != "<br class='class1 class2'/>" {
		t.Error("Expected '<br class='class1,class2'/>', got", s)
	}

	br = Break{}
	br.SetStyle("key", "value")
	s = br.String()
	if s != "<br style='key:value;'/>" {
		t.Error("Expected '<br style='key:value;'/>', got", s)
	}

	br = Break{}
	br.SetStyle("key", "value")
	br.SetStyle("key2", "value2")
	s = br.String()
	if s != "<br style='key:value;key2:value2;'/>" {
		t.Error("Expected '<br style='key:value;key2:value2;'/>', got", s)
	}

	br = Break{}
	br.SetStyle("key", "value")
	br.AddClass("class1")
	s = br.String()
	if s != "<br class='class1' style='key:value;'/>" {
		t.Error("Expected '<br class='class1' style='key:value;'/>', got", s)
	}
}

func TestBreakElement(t *testing.T) {
	var br Element

	br = Break{}
	s := br.String()
	if s != "<br/>" {
		t.Error("Expected '<br/>', got ", s)
	}
}
