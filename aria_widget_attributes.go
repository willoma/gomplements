package gomplements

import (
	"strconv"

	"maragu.dev/gomponents"
)

type ariaAutocomplete string

// aria-autocomplete attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-autocomplete
const (
	AriaAutocompleteNone ariaAutocomplete = "none"
	AriaAutocompleteList ariaAutocomplete = "list"
	AriaAutocompleteBoth ariaAutocomplete = "both"
)

func (a ariaAutocomplete) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-autocomplete", string(a)))
}

type ariaChecked string

// aria-checked attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-checked
const (
	AriaCheckedUndefined ariaChecked = "undefined"
	AriaCheckedTrue      ariaChecked = "true"
	AriaCheckedFalse     ariaChecked = "false"
	AriaCheckedMixed     ariaChecked = "mixed"
)

func (a ariaChecked) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-checked", string(a)))
}

type ariaDisabled string

// aria-disabled attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-disabled
const (
	AriaDisabledTrue  ariaDisabled = "true"
	AriaDisabledFalse ariaDisabled = "false"
)

func (a ariaDisabled) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-disabled", string(a)))
}

// aria-errormessage attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-errormessage
type AriaErrorMessageID string

func (a AriaErrorMessageID) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-errormessage", string(a)))
}

type ariaExpanded string

// aria-expanded attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-expanded
const (
	AriaExpandedUndefined ariaExpanded = "undefined"
	AriaExpandedTrue      ariaExpanded = "true"
	AriaExpandedFalse     ariaExpanded = "false"
)

func (a ariaExpanded) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-expanded", string(a)))
}

type ariaHasPopup string

// aria-haspopup attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-haspopup
const (
	AriaHasPopupTrue    ariaHasPopup = "true"
	AriaHasPopupFalse   ariaHasPopup = "false"
	AriaHasPopupMenu    ariaHasPopup = "menu"
	AriaHasPopupListbox ariaHasPopup = "listbox"
	AriaHasPopupTree    ariaHasPopup = "tree"
	AriaHasPopupGrid    ariaHasPopup = "grid"
	AriaHasPopupDialog  ariaHasPopup = "dialog"
)

func (a ariaHasPopup) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-haspopup", string(a)))
}

type ariaHidden string

// aria-hidden attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-hidden
const (
	AriaHiddenUndefined ariaHidden = "undefined"
	AriaHiddenTrue      ariaHidden = "true"
	AriaHiddenFalse     ariaHidden = "false"
)

func (a ariaHidden) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-hidden", string(a)))
}

type ariaInvalid string

// aria-invalid attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-invalid
const (
	AriaInvalidTrue     ariaInvalid = "true"
	AriaInvalidFalse    ariaInvalid = "false"
	AriaInvalidGrammar  ariaInvalid = "grammar"
	AriaInvalidSpelling ariaInvalid = "spelling"
)

func (a ariaInvalid) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-invalid", string(a)))
}

// aria-label attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-label
type AriaLabel string

func (a AriaLabel) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-label", string(a)))
}

// aria-level attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-level
type AriaLevel int

func (a AriaLevel) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-level", strconv.Itoa(int(a))))
}

type ariaModal string

// aria-modal attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-modal
const (
	AriaModalTrue  ariaModal = "true"
	AriaModalFalse ariaModal = "false"
)

func (a ariaModal) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-modal", string(a)))
}

type ariaMultiline string

// aria-multiline attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-multiline
const (
	AdiaMultilineTrue  ariaMultiline = "true"
	AdiaMultilineFalse ariaMultiline = "false"
)

func (a ariaMultiline) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-multiline", string(a)))
}

type ariaMultiselectable string

// aria-multiselectable attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-multiselectable
const (
	AdiaMultiselectableTrue  ariaMultiselectable = "true"
	AdiaMultiselectableFalse ariaMultiselectable = "false"
)

func (a ariaMultiselectable) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-multiselectable", string(a)))
}

type ariaOrientation string

// aria-orientation attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-orientation
const (
	AriaOrientationUndefined  ariaOrientation = "undefined"
	AriaOrientationHorizontal ariaOrientation = "horizontal"
	AriaOrientationVertical   ariaOrientation = "vertical"
)

func (a ariaOrientation) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-orientation", string(a)))
}

// aria-placeholder attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-placeholder
type AriaPlaceholder string

func (a AriaPlaceholder) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-placeholder", string(a)))
}

type ariaPressed string

// aria-pressed attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-pressed
const (
	AriaPressedUndefined ariaPressed = "undefined"
	AriaPressedTrue      ariaPressed = "true"
	AriaPressedFalse     ariaPressed = "false"
	AriaPressedMixed     ariaPressed = "mixed"
)

func (a ariaPressed) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-pressed", string(a)))
}

type ariaReadonly string

// aria-readonly attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-readonly
const (
	AriaReadonlyTrue  ariaReadonly = "true"
	AriaReadonlyFalse ariaReadonly = "false"
)

func (a ariaReadonly) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-readonly", string(a)))
}

type ariaRequired string

// aria-required attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-required
const (
	AriaRequiredTrue  ariaRequired = "true"
	AriaRequiredFalse ariaRequired = "false"
)

func (a ariaRequired) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-required", string(a)))
}

type ariaSelected string

// aria-selected attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-selected
const (
	AriaSelectedUndefined ariaSelected = "undefined"
	AriaSelectedTrue      ariaSelected = "true"
	AriaSelectedFalse     ariaSelected = "false"
)

func (a ariaSelected) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-selected", string(a)))
}

type ariaSort string

// aria-sort attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-sort
const (
	AriaSortNone       ariaSort = "none"
	AriaSortAscending  ariaSort = "ascending"
	AriaSortDescending ariaSort = "descending"
	AriaSortOther      ariaSort = "other"
)

func (a ariaSort) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-sort", string(a)))
}

// aria-valuemax attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-valuemax
type AriaValueMax float64

func (a AriaValueMax) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-valuemax", strconv.FormatFloat(float64(a), 'f', -1, 64)))
}

// aria-valuemin attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-valuemin
type AriaValueMin float64

func (a AriaValueMin) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-valuemin", strconv.FormatFloat(float64(a), 'f', -1, 64)))
}

// aria-valuenow attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-valuenow
type AriaValueNow float64

func (a AriaValueNow) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-valuenow", strconv.FormatFloat(float64(a), 'f', -1, 64)))
}

// aria-valuetext attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-valuetext
type AriaValueText string

func (a AriaValueText) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-valuetext", string(a)))
}
