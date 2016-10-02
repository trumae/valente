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
