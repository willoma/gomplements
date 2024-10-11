package gomplements

import (
	"bytes"
	"io"

	"maragu.dev/gomponents"
)

// Prepare pre-renders a node in memory for future uses.
func Prepare(e Element) gomponents.Node {
	var buf bytes.Buffer
	e.Render(&buf)

	if modifier, ok := e.(ParentModifier); ok {
		return &preparedElementWithParentModifier{
			preparedElement: preparedElement{buf.Bytes()},
			modifyFn:        modifier.ModifyParent,
		}
	}

	return &preparedElement{content: buf.Bytes()}
}

type preparedElement struct {
	content []byte
}

type preparedElementWithParentModifier struct {
	preparedElement
	modifyFn func(Element)
}

func (e *preparedElement) Render(w io.Writer) error {
	_, err := w.Write(e.content)
	return err
}

func (e *preparedElementWithParentModifier) ModifyParent(parent Element) {
	e.modifyFn(parent)
}
