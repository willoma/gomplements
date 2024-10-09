package gomplements

import "maragu.dev/gomponents"

// ParentModifier is an interface for modifying the parent of an element.
type ParentModifier interface {
	ModifyParent(parent Element)
}

// ParentModifierAndNode is an interface for modifying the parent of an element and adding a child.
type ParentModifierAndNode interface {
	gomponents.Node
	ParentModifier
}
