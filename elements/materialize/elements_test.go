package elements

import (
	"testing"
)

func TestRow(t *testing.T) {
	row := NewRow()
	s := row.String()

	if s != "<div class='row'></div>" {
		t.Fatal("Value not expected")
	}
}

func TestCol(t *testing.T) {
	col := NewCol()

	s := col.String()
	if s != "<div class='col'></div>" {
		t.Fatal("Value not expected")
	}

	col.S = 6
	s = col.String()
	if s != "<div class='col s6'></div>" {
		t.Fatal("Value not expected")
	}

}

func TestImageCard(t *testing.T) {
	c := NewImageCard("Title", "Image", "arrow_forward", "Text Test")

	s := c.String()
	if s != "<div class='card'><div class='card-image'><img src='Image'/><span class='card-title'>Title</span><a class='btn-floating halfway-fab waves-effect waves-light red'><i class='material-icons'>arrow_forward</i></a></div><div class='card-content'><p>Text Test</p></div></div>" {
		t.Fatal("Value not expected")
	}
}

func TestInputSelectRow(t *testing.T) {
	val := `<div class="row">
	      <div class="input-field col s12">
	        <select id="id">
                  <option id='opvalue' value='value'>Value</option><option id='opotheroption' value='otheroption'>Other Option</option>
	        </select>
	        <label for="id">label</label>
	      </div>
	    </div>`

	s := NewInputSelectRow("id", "label", "value",
		[]string{"value", "otheroption"}, []string{"Value", "Other Option"})

	if s.String() != val {
		t.Fatal("Value not expected", len(s.String()), len(val))
	}

}
