package gomplements

// Styles allows to define CSS styles to apply to an Element.
type Styles map[string]string

// If allows to apply CSS styles to an Element, only if a condition is met.
func (s Styles) If(cond bool) Styles {
	if cond {
		return s
	}
	return Styles{}
}
