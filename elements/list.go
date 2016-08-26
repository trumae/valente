package elements

//ListElement produces an HTML listitem element (<li>).
type ListElement struct {
	Base
	Body       []Element
	Text       string
	HTMLEncode bool
}

//List element produces an HTML list element (<ol> and <ul>).
type List struct {
	Base
	Body     []ListElement
	Numbered bool
	Role     string
}
