package gomplements

import (
	"fmt"
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Element is an element, which implements the gomponents.Node interface with
type Element interface {
	gomponents.Node

	// With adds childs to the element.
	With(...any) Element

	// Clone returns an independent cloned element.
	Clone() Element
}

type nodeTypeDescriber interface {
	Type() gomponents.NodeType
}

// IsAttribute returns true if the provided node is an attribute according to
// its gomponents.NodeType.
func IsAttribute(node any) bool {
	desc, ok := node.(nodeTypeDescriber)
	return ok && desc.Type() == gomponents.AttributeType
}

type elem struct {
	elemFn func(...gomponents.Node) gomponents.Node

	classes map[Class]bool
	styles  map[string]string

	attributes []gomponents.Node
	elements   []gomponents.Node
}

// Elem creates a base element. It takes a function that generates a
// gomponents.Node (these functions are found in the gomponents/html package),
// and optionally some children.
//
// The children are passed as arguments to the With method.
func Elem(elemFn func(...gomponents.Node) gomponents.Node, children ...any) *elem {
	e := &elem{
		elemFn:  elemFn,
		classes: map[Class]bool{},
		styles:  map[string]string{},
	}
	e.With(children...)
	return e
}

// With adds childs to the element. It accepts the following arguments types:
//
//   - func(...gomponents.Node) gomponents.Node: change the gomponents.Node builder function
//   - Class, Classer, Classeser: add a class or multiple classes to the element
//   - Styles: add one or multiple CSS styles to the element
//   - ParentModifierAndNode: execute the ModifyParent method with the element as its argument, and add as a child
//   - ParentModifier: execute the ModifyParent method with the element as its argument
//   - Element: add the provided element as a child
//   - gomponents.Node with type gomponents.AttributeType: add the provided attribute to the element
//   - gomponents.Node with another time: add the provided element as a child
//   - ID: define the id attribute for the element
//   - string, fmt.Stringer: add a string to the element (using gomponents.Text)
//   - []any: add the slice items by executing the With method recursively
//
// Any other type is ignored.
func (e *elem) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case func(...gomponents.Node) gomponents.Node:
			if c == nil {
				continue
			}
			e.elemFn = c
		case Class:
			if c != "" {
				e.classes[c] = true
			}
		case Classer:
			cl := c.Class()
			if cl != "" {
				e.classes[cl] = true
			}
		case Classeser:
			for _, cl := range c.Classes() {
				if cl != "" {
					e.classes[cl] = true
				}
			}
		case Styles:
			for prop, val := range c {
				e.styles[prop] = val
			}
		case ParentModifierAndNode:
			if c == nil {
				continue
			}
			c.ModifyParent(e)
			e.elements = append(e.elements, c)
		case ParentModifier:
			if c == nil {
				continue
			}
			c.ModifyParent(e)
		case Element:
			if c == nil {
				continue
			}
			e.elements = append(e.elements, c)
		case gomponents.Node:
			if c == nil {
				continue
			}
			if IsAttribute(c) {
				e.attributes = append(e.attributes, c)
			} else {
				e.elements = append(e.elements, c)
			}
		case ID:
			if c != "" {
				e.attributes = append(e.attributes, html.ID(string(c)))
			}
		case string:
			if c != "" {
				e.elements = append(e.elements, gomponents.Text(c))
			}
		case fmt.Stringer:
			s := c.String()
			if s != "" {
				e.elements = append(e.elements, gomponents.Text(s))
			}
		case []any:
			e.With(c...)
		}
	}

	return e
}

// Clone returns an independent cloned element.
func (e *elem) Clone() Element {
	if e == nil {
		return nil
	}

	classes := make(map[Class]bool, len(e.classes))
	for class, ok := range e.classes {
		classes[class] = ok
	}

	stylesCollection := make(map[string]string, len(e.styles))
	for style, value := range e.styles {
		stylesCollection[style] = value
	}

	attributes := make([]gomponents.Node, len(e.attributes))
	for i, attr := range e.attributes {
		attributes[i] = attr
	}

	elements := make([]gomponents.Node, len(e.elements))
	for i, elem := range e.elements {
		switch elem := elem.(type) {
		case Element:
			elements[i] = elem.Clone()
		default:
			elements[i] = elem
		}
	}

	return &elem{
		elemFn: e.elemFn,

		classes: classes,
		styles:  stylesCollection,

		attributes: attributes,
		elements:   elements,
	}
}

// GetNodes returns the element's children as gomponents nodes.
func (e *elem) GetNodes() []gomponents.Node {
	children := []gomponents.Node{}

	classes := Class("")
	for cl := range e.classes {
		classes += " " + cl
	}
	if classes != "" {
		children = append(children, html.Class(string(classes[1:])))
	}

	styles := ""
	for prop, val := range e.styles {
		styles += prop + ":" + val + ";"
	}
	if styles != "" {
		children = append(children, html.StyleAttr(styles))
	}

	for _, attr := range e.attributes {
		children = append(children, attr)
	}

	children = append(children, e.elements...)

	return children
}

// Render renders the element to a writer.
func (e *elem) Render(w io.Writer) error {
	if e == nil {
		return nil
	}
	return e.elemFn(e.GetNodes()...).Render(w)
}
