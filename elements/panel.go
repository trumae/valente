package elements

//Panel element produces an HTML div
type Panel struct {
	Base
	Body       []Element
	Text       string
	HTMLEncode bool
}
