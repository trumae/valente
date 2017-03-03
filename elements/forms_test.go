package elements

import "testing"

func TestButton(t *testing.T) {
	btn := Button{}
	s := btn.String()
	if s != "<button></button>" {
		t.Error("Expected '<button></button>', got", s)
	}

	btn = Button{Text: "valente"}
	s = btn.String()
	if s != "<button>valente</button>" {
		t.Error("Expected '<button>valente</button>', got", s)
	}

	btn = Button{Text: "valente", PostBack: []string{"back"}}
	s = btn.String()
	if s != "<button onclick=\"sendEvent('back')\">valente</button>" {
		t.Error("Expected '<button onclick=\"sendEvent('back')\">valente</button>', got", s)
	}

	btn = Button{PostBack: []string{"back"}}
	btn.AddClass("class1")
	s = btn.String()
	if s != "<button class='class1' onclick=\"sendEvent('back')\"></button>" {
		t.Error("Expected '<button class='class1' onclick=\"sendEvent('back')\"></button>', got", s)
	}

	btn = Button{Text: "valente", PostBack: []string{"back", "two"}}
	s = btn.String()
	if s != "<button onclick=\"sendEvent('back','two')\">valente</button>" {
		t.Error("Expected '<button onclick=\"sendEvent('back','two')\">valente</button>', got", s)
	}

}

func TestInputText(t *testing.T) {
	ta := InputText{}
	s := ta.String()
	if s != "<input type='text' value=''/>" {
		t.Error("Expected '<input type='text' value=''/>', got", s)
	}

	ta = InputText{Value: "Valente"}
	s = ta.String()
	if s != "<input type='text' value='Valente'/>" {
		t.Error("Expected '<input type='text' value='Valente'/>', got", s)
	}

	ta = InputText{Value: "Valente"}
	ta.SetData("data1", "10")
	s = ta.String()
	if s != "<input type='text' value='Valente' data1='10'/>" {
		t.Error("Expected '<input type='text' value='Valente' data1='10'/>', got", s)
	}

	ta = InputText{Value: "Valente"}
	ta.SetData("data1", "10")
	ta.SetData("data2", "20")
	ta.AddClass("class1")
	s = ta.String()
	if s != "<input type='text' value='Valente' class='class1' data1='10' data2='20'/>" {
		t.Error("Expected '<input type='text' value='Valente' class='class1' data1='10' data2='20'/>', got", s)
	}

	ta = InputText{Value: "Valente"}
	ta.SetData("data1", "10")
	ta.SetData("data2", "20")
	s = ta.String()
	if s != "<input type='text' value='Valente' data1='10' data2='20'/>" {
		t.Error("Expected '<input type='text' value='Valente' data1='10' data2='20'/>', got", s)
	}
}

func TestTextArea(t *testing.T) {
	ta := TextArea{}
	s := ta.String()
	if s != "<textarea></textarea>" {
		t.Error("Expected '<textarea></textares>', got", s)
	}

	ta = TextArea{Text: "Valente"}
	s = ta.String()
	if s != "<textarea>Valente</textarea>" {
		t.Error("Expected '<textarea>Valente</textares>', got", s)
	}

	ta = TextArea{Text: "Valente"}
	ta.SetData("rows", "10")
	s = ta.String()
	if s != "<textarea rows='10'>Valente</textarea>" {
		t.Error("Expected '<textarea rows='10'>Valente</textarea>', got", s)
	}

	ta = TextArea{Text: "Valente"}
	ta.SetData("rows", "10")
	ta.SetData("cols", "40")
	s = ta.String()
	if s != "<textarea rows='10' cols='40'>Valente</textarea>" {
		t.Error("Expected '<textarea rows='10' cols='40'>Valente</textarea>', got", s)
	}
}
