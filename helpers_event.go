package gomplements

import (
	"io"
	"strings"

	"maragu.dev/gomponents"
)

// On adds a "on<event>" attribute to an element.
func On(event, script string) gomponents.Node {
	return &eventAttr{event: event, script: script}
}

// OnClick adds a "onclick" attribute to an element.
func OnClick(script string) gomponents.Node {
	return &eventAttr{event: "click", script: script}
}

type eventAttr struct {
	event  string
	script string
}

func (a *eventAttr) Render(w io.Writer) error {
	_, err := w.Write([]byte(" on" + a.event + `="` + strings.ReplaceAll(a.script, `"`, "&#34;") + `"`))
	return err
}

func (a *eventAttr) Type() gomponents.NodeType {
	return gomponents.AttributeType
}
