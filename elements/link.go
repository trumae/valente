package elements

//Link produce the html anchor
type Link struct {
	Base
	Text       string
	Body       []Element
	HTMLEncode bool
	URL        string
	NewWindow  bool
	Click      string
}
