package elements

//Break produce the element br
type Break struct {
	Base
}

func (b Break) String() string {
	return "<br/>"
}
