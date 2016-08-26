package elements

//Image produce the img html element
type Image struct {
	Base
	Source string
}

//String return img tag
func (image Image) String() string {
	ret := ""
	ret = "<img" + image.Attrs()
	if image.Source != "" {
		ret += " src='" + image.Source + "'"
	}
	ret += "/>"

	return ret
}
