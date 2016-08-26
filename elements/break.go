package elements

//Break produce the element br
type Break struct {
	Base
}

//String return string tag for Break
func (b Break) String() string {
	return "<br/>"
}
