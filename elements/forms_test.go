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
