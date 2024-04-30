package gomplements

import "github.com/maragudk/gomponents"

type ariaBusy string

// aria-busy attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-busy
const (
	AriaBusyTrue  ariaBusy = "true"
	AriaBusyFalse ariaBusy = "false"
)

func (a ariaBusy) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-busy", string(a)))
}

type ariaLive string

// aria-live attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-live
const (
	AriaLiveOff       ariaLive = "off"
	AriaLiveAssertive ariaLive = "assertive"
	AriaLivePolite    ariaLive = "polite"
)

func (a ariaLive) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-live", string(a)))
}

type ariaRelevant string

// aria-relevant attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-relevant
const (
	AriaRelevantAdditions     ariaRelevant = "additions"
	AriaRelevantAll           ariaRelevant = "all"
	AriaRelevantRemovals      ariaRelevant = "removals"
	AriaRelevantText          ariaRelevant = "text"
	AriaRelevantAdditionsText ariaRelevant = "additions text"
)

func (a ariaRelevant) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-relevant", string(a)))
}

type ariaAtomic string

// aria-atomic attribute
// https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-atomic
const (
	AriaAtomicTrue  ariaAtomic = "true"
	AriaAtomicFalse ariaAtomic = "false"
)

func (a ariaAtomic) ModifyParent(p Element) {
	p.With(gomponents.Attr("aria-atomic", string(a)))
}
