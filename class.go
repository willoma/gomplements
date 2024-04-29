package gomplements

// Class represents a CSS class, which is applied in the `class` attribute of an HTML element.
type Class string

// Classes contains a list of CSS classes to be applied in the `class` attribute of an HTML element.
type Classes []Class

// Classes implements Classeser.
func (c Classes) Classes() []Class {
	return c
}

// Classer is an interface that allows to add a CSS class to an HTML element.
type Classer interface {
	Class() Class
}

// Classeser is an interface that allows to add multiple CSS classes to an HTML element.
type Classeser interface {
	Classes() []Class
}
