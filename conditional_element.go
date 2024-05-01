package gomplements

type ElementBuilder func(...any) Element

// If builds the provided element with the provided children, only if the condition is true.
//
// If condition is false, the element is not built and nil is returned.
func If(condition bool, builder ElementBuilder, children ...any) Element {
	if !condition {
		return nil
	}

	return builder(children...)
}
