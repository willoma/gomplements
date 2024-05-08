package gomplements

// Styles allows to define CSS styles to apply to an Element.
type Styles map[string]string

// If allows to apply CSS styles to an Element, only if a condition is met.
func (s Styles) If(cond bool) ParentModifier {
	return &conditionalStyles{styles: s, cond: cond}
}

type conditionalStyles struct {
	styles Styles
	cond   bool
}

func (c *conditionalStyles) ModifyParent(p Element) {
	if c.cond {
		p.With(c.styles)
	}
}
