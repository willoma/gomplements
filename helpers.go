package gomplements

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// AHref creates an <a> element with the provided href.
func AHref(href string, children ...any) Element {
	return Elem(html.A, html.Href(href), children)
}

// ID adds an "id" attribute to an element.
//
// It is useful for some components which need to know that an argument is an ID
// in order to apply it to the correct element.
type ID string

// ImgSrc creates a <img> element, with the provided src.
func ImgSrc(src string, children ...any) Element {
	return Elem(html.Img, html.Src(src), children)
}

// RawScript creates a <script> element, with the provided content as a script, unmodified.
func RawScript(script string) Element {
	return Script(gomponents.Raw(script))
}
