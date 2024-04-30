package gomplements

import (
	"strconv"
	"strings"

	"github.com/maragudk/gomponents"
)

// aria-activedescendant attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-activedescendant
type AriaActiveDescendantID string

func (a AriaActiveDescendantID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-activedescendant", string(a)))
}

// aria-colcount attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-colcount
type AriaColCount int

func (a AriaColCount) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-colcount", strconv.Itoa(int(a))))
}

// aria-colindex attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-colindex
type AriaColIndex int

func (a AriaColIndex) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-colindex", strconv.Itoa(int(a))))
}

// aria-colspan attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-colspan
type AriaColSpan int

func (a AriaColSpan) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-colspan", strconv.Itoa(int(a))))
}

type ariaControlsID []string

// aria-controls attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-controls
func AriaControlsID(ids ...string) ariaControlsID {
	return ariaControlsID(ids)
}

func (a ariaControlsID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-controls", strings.Join(a, " ")))
}

// aria-describedby attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-describedby
type AriaDescribedBy string

func (a AriaDescribedBy) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-describedby", string(a)))
}

// aria-description attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-description
type AriaDescription string

func (a AriaDescription) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-description", string(a)))
}

type ariaDetailsID []string

// aria-details attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-details
func AriaDetailsID(ids ...string) ariaDetailsID {
	return ariaDetailsID(ids)
}

func (a ariaDetailsID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-details", strings.Join(a, " ")))
}

type ariaFlowToID []string

// aria-flowto attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-flowto
func AriaFlowToID(ids ...string) ariaFlowToID {
	return ariaFlowToID(ids)
}

func (a ariaFlowToID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-flowto", strings.Join(a, " ")))
}

type ariaLabelledByID []string

// aria-labelledby attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-labelledby
func AriaLabelledByID(ids ...string) ariaLabelledByID {
	return ariaLabelledByID(ids)
}

func (a ariaLabelledByID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-labelledby", strings.Join(a, " ")))
}

type ariaOwnsID []string

// aria-owns attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-owns
func AriaOwnsID(ids ...string) ariaOwnsID {
	return ariaOwnsID(ids)
}

func (a ariaOwnsID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-owns", strings.Join(a, " ")))
}

// aria-posinset attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-posinset
type AriaPosInSet int

func (a AriaPosInSet) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-posinset", strconv.Itoa(int(a))))
}

// aria-rowcount
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-rowcount
type AriaRowCount int

func (a AriaRowCount) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-rowcount", strconv.Itoa(int(a))))
}

// aria-rowindex
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-rowindex
type AriaRowIndex int

func (a AriaRowIndex) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-rowindex", strconv.Itoa(int(a))))
}

// aria-rowspan
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-rowspan
type AriaRowSpan int

func (a AriaRowSpan) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-rowspan", strconv.Itoa(int(a))))
}

// aria-setsize
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-setsize
type AriaSetSize int

func (a AriaSetSize) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-setsize", strconv.Itoa(int(a))))
}
