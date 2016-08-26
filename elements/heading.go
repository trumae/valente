package elements

//Heading1 element produces an HTML heading element (<h1>)
type Heading1 struct {
	Base
	Text string
}

//String emit html text for Heading1
func (h1 Heading1) String() string {
	return "<h1>" + h1.Text + "</h1>"
}

//Heading2 element produces an HTML heading element (<h1>)
type Heading2 struct {
	Base
	Text string
}

//String emit html text for Heading2
func (h2 Heading2) String() string {
	return "<h2>" + h2.Text + "</h2>"
}

//Heading3 element produces an HTML heading element (<h1>)
type Heading3 struct {
	Base
	Text string
}

//String emit html text for Heading3
func (h3 Heading3) String() string {
	return "<h3>" + h3.Text + "</h3>"
}

//Heading4 element produces an HTML heading element (<h1>)
type Heading4 struct {
	Base
	Text string
}

//String emit html text for Heading4
func (h4 Heading4) String() string {
	return "<h4>" + h4.Text + "</h4>"
}
