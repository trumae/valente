package elements

//Paragraph element produces an HTML p
type Paragraph struct {
	Base
	Text       string
	HTMLEncode bool
}
