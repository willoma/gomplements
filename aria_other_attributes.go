package gomplements

import (
	"strings"

	"github.com/maragudk/gomponents"
)

type ariaCurrent string

// aria-current attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-current
const (
	AriaCurrentPage     ariaCurrent = "page"
	AriaCurrentStep     ariaCurrent = "step"
	AriaCurrentLocation ariaCurrent = "location"
	AriaCurrentDate     ariaCurrent = "date"
	AriaCurrentTime     ariaCurrent = "time"
	AriaCurrentTrue     ariaCurrent = "true"
	AriaCurrentFalse    ariaCurrent = "false"
)

func (a ariaCurrent) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-current", string(a)))
}

type ariaKeyShortcuts []string

// aria-keyshortcuts attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-keyshortcuts
func AriaKeyShortcuts(shortcuts ...string) ariaKeyShortcuts {
	return ariaKeyShortcuts(shortcuts)
}

func (a ariaKeyShortcuts) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-keyshortcuts", strings.Join(a, " ")))
}

// aria-roledescription attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-roledescription
type AriaRoleDescription string

func (a AriaRoleDescription) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-roledescription", string(a)))
}

// aria-braillelabel attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-braillelabel
type AriaBrailleLabel string

func (a AriaBrailleLabel) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-braillelabel", string(a)))
}

// aria-brailleroledescription attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-brailleroledescription
type AriaBrailleRoleDescription string

func (a AriaBrailleRoleDescription) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-brailleroledescription", string(a)))
}

// aria-colindextext attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-colindextext
type AriaColIndexText string

func (a AriaColIndexText) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-colindextext", string(a)))
}

// aria-rowindextext attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-rowindextext
type AriaRowIndexText string

func (a AriaRowIndexText) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-rowindextext", string(a)))
}
