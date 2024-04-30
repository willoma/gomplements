package gomplements

import "github.com/maragudk/gomponents"

type ariaRole string

// Document structure ARIA roles
const (
	AriaToolbar      ariaRole = "toolbar"
	AriaTooltip      ariaRole = "tooltip"
	AriaFeed         ariaRole = "feed"
	AriaMath         ariaRole = "math"
	AriaPresentation ariaRole = "presentation"
	AriaNone         ariaRole = "none"
	AriaNote         ariaRole = "note"
)

// Document structure ARIA roles to avoid
const (
	AriaApplication  ariaRole = "application"
	AriaArticle      ariaRole = "article"
	AriaCell         ariaRole = "cell"
	AriaColumnHeader ariaRole = "columnheader"
	AriaDefinition   ariaRole = "definition"
	AriaDocument     ariaRole = "document"
	AriaFigure       ariaRole = "figure"
	AriaGroup        ariaRole = "group"
	AriaHeading      ariaRole = "heading"
	AriaImg          ariaRole = "img"
	AriaList         ariaRole = "list"
	AriaListItem     ariaRole = "listitem"
	AriaMeter        ariaRole = "meter"
	AriaRow          ariaRole = "row"
	AriaRowGroup     ariaRole = "rowgroup"
	AriaRowHeader    ariaRole = "rowheader"
	AriaTable        ariaRole = "table"
	AriaTerm         ariaRole = "term"
)

// Document structure ARIA roles rarely, if ever, useful
const (
	AriaAssociationList          ariaRole = "associationlist"
	AriaAssociationListItemKey   ariaRole = "associationlistitemkey"
	AriaAssociationListItemValue ariaRole = "associationlistitemvalue"
	AriaBlockquote               ariaRole = "blockquote"
	AriaCaption                  ariaRole = "caption"
	AriaCode                     ariaRole = "code"
	AriaDeletion                 ariaRole = "deletion"
	AriaEmphasis                 ariaRole = "emphasis"
	AriaInsertion                ariaRole = "insertion"
	AriaParagraph                ariaRole = "paragraph"
	AriaStrong                   ariaRole = "strong"
	AriaSubscription             ariaRole = "subscript"
	AriaSuperscript              ariaRole = "superscript"
	AriaTime                     ariaRole = "time"
)

// Widget ARIA roles
const (
	AriaScrollbar  ariaRole = "scrollbar"
	AriaSearchbox  ariaRole = "searchbox"
	AriaSeparator  ariaRole = "separator"
	AriaSlider     ariaRole = "slider"
	AriaSpinButton ariaRole = "spinbutton"
	AriaSwitch     ariaRole = "switch"
	AriaTab        ariaRole = "tab"
	AriaTabPanel   ariaRole = "tabpanel"
	AriaTreeItem   ariaRole = "treeitem"
	AriaComboBox   ariaRole = "combobox"
	AriaMenu       ariaRole = "menu"
	AriaMenuBar    ariaRole = "menubar"
	AriaTabList    ariaRole = "tablist"
	AriaTree       ariaRole = "tree"
	AriaTreeGrid   ariaRole = "treegrid"
)

// Widget ARIA roles to avoid
const (
	AriaButton           ariaRole = "button"
	AriaCheckbox         ariaRole = "checkbox"
	AriaGridCell         ariaRole = "gridcell"
	AriaLink             ariaRole = "link"
	AriaMenuItem         ariaRole = "menuitem"
	AriaMenuItemCheckbox ariaRole = "menuitemcheckbox"
	AriaMenuItemRadio    ariaRole = "menuitemradio"
	AriaOption           ariaRole = "option"
	AriaProgressbar      ariaRole = "progressbar"
	AriaRadio            ariaRole = "radio"
	AriaTextbox          ariaRole = "textbox"
	AriaGrid             ariaRole = "grid"
	AriaListbox          ariaRole = "listbox"
	AriaRadioGroup       ariaRole = "radiogroup"
)

// Landmark ARIA roles
const (
	AriaBanner        ariaRole = "banner"
	AriaComplementary ariaRole = "complementary"
	AriaContentInfo   ariaRole = "contentinfo"
	AriaForm          ariaRole = "form"
	AriaMain          ariaRole = "main"
	AriaNavigation    ariaRole = "navigation"
	AriaRegion        ariaRole = "region"
	AriaSearch        ariaRole = "search"
)

// Live region ARIA roles
const (
	AriaAlert   ariaRole = "alert"
	AriaLog     ariaRole = "log"
	AriaMarquee ariaRole = "marquee"
	AriaStatus  ariaRole = "status"
	AriaTimer   ariaRole = "timer"
)

// Window ARIA roles
const (
	AriaAlertDialog ariaRole = "alertdialog"
	AriaDialog      ariaRole = "dialog"
)

func (a ariaRole) ModifyParent(p Element) {
	p.With(gomponents.Attr("role", string(a)))
}
