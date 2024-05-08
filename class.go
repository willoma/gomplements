package gomplements

// Class represents a CSS class, which is applied in the class attribute of a Element.
type Class string

// If allows to add a CSS class to an Element, only if a condition is met.
func (c Class) If(condition bool) ParentModifier {
	return &conditionalClass{class: c, cond: condition}
}

type conditionalClass struct {
	class Class
	cond  bool
}

func (c *conditionalClass) ModifyParent(p Element) {
	if c.cond {
		p.With(c.class)
	}
}

// Classes contains multiple CSS classes to be applied in the class attribute of an Element.
type Classes []Class

// Classes implements Classeser.
func (c Classes) Classes() []Class {
	return c
}

// If allows to add multiple CSS classes to an Element, only if a condition is met.
func (c Classes) If(condition bool) ParentModifier {
	return &conditionalClasses{classes: c, cond: condition}
}

type conditionalClasses struct {
	classes Classes
	cond    bool
}

func (c *conditionalClasses) ModifyParent(p Element) {
	if c.cond {
		p.With(c.classes)
	}
}

// Classer is an interface that allows to add a CSS class to an HTML element.
type Classer interface {
	Class() Class
}

// Classeser is an interface that allows to add multiple CSS classes to an HTML element.
type Classeser interface {
	Classes() []Class
}
