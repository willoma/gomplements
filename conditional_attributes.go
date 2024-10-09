package gomplements

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func newConditionalAttribute(attrFn func() gomponents.Node, condition ...bool) gomponents.Node {
	if len(condition) == 0 || condition[0] {
		return attrFn()
	}

	return nil
}

type conditionalAttribute struct {
	attrFn func() gomponents.Node
	cond   bool
}

func (b conditionalAttribute) ModifyParent(p Element) {
	if b.cond {
		p.With(b.attrFn())
	}
}

func newConditionalAttributeWithValue(attrFn func(string) gomponents.Node, value string, condition ...bool) gomponents.Node {
	if len(condition) == 0 || condition[0] {
		return attrFn(value)
	}

	return nil
}

type conditionalAttributeWithValue struct {
	attrFn func(string) gomponents.Node
	value  string
	cond   bool
}

func (b conditionalAttributeWithValue) ModifyParent(p Element) {
	if b.cond {
		p.With(b.attrFn(b.value))
	}
}

func newConditionalAttributeWithNameAndValue(attrFn func(string, string) gomponents.Node, name, value string, condition ...bool) gomponents.Node {
	if len(condition) == 0 || condition[0] {
		return attrFn(name, value)
	}

	return nil
}

type conditionalAttributeWithNameAndValue struct {
	attrFn func(string, string) gomponents.Node
	name   string
	value  string
	cond   bool
}

func (b conditionalAttributeWithNameAndValue) ModifyParent(p Element) {
	if b.cond {
		p.With(b.attrFn(b.name, b.value))
	}
}

func Async(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Async, condition...)
}

func AutoFocus(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.AutoFocus, condition...)
}

func AutoPlay(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.AutoPlay, condition...)
}

func Checked(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Checked, condition...)
}

func Defer(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Defer, condition...)
}

func Disabled(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Disabled, condition...)
}

func Loop(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Loop, condition...)
}

func Multiple(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Multiple, condition...)
}

func Muted(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Muted, condition...)
}

func PlaysInline(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.PlaysInline, condition...)
}

func ReadOnly(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.ReadOnly, condition...)
}

func Required(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Required, condition...)
}

func Selected(condition ...bool) gomponents.Node {
	return newConditionalAttribute(html.Selected, condition...)
}

func Accept(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Accept, value, condition...)
}

func Action(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Action, value, condition...)
}

func Alt(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Alt, value, condition...)
}

func As(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.As, value, condition...)
}

func AutoComplete(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.AutoComplete, value, condition...)
}

func Charset(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Charset, value, condition...)
}

func Cols(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Cols, value, condition...)
}

func ColSpan(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.ColSpan, value, condition...)
}

func Content(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Content, value, condition...)
}

func DataAttr(name, v string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithNameAndValue(html.DataAttr, name, v, condition...)
}

func For(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.For, value, condition...)
}

func FormAttr(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.FormAttr, value, condition...)
}

func Height(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Height, value, condition...)
}

func Href(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Href, value, condition...)
}

func Lang(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Lang, value, condition...)
}

func Loading(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Loading, value, condition...)
}

func Max(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Max, value, condition...)
}

func MaxLength(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.MaxLength, value, condition...)
}

func Method(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Method, value, condition...)
}

func Min(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Min, value, condition...)
}

func MinLength(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.MinLength, value, condition...)
}

func Name(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Name, value, condition...)
}

func Pattern(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Pattern, value, condition...)
}

func Placeholder(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Placeholder, value, condition...)
}

func Poster(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Poster, value, condition...)
}

func Preload(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Preload, value, condition...)
}

func Rel(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Rel, value, condition...)
}

func Rows(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Rows, value, condition...)
}

func RowSpan(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.RowSpan, value, condition...)
}

func Src(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Src, value, condition...)
}

func SrcSet(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.SrcSet, value, condition...)
}

func Step(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Step, value, condition...)
}

func TabIndex(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.TabIndex, value, condition...)
}

func Target(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Target, value, condition...)
}

func TitleAttr(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.TitleAttr, value, condition...)
}

func Type(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Type, value, condition...)
}

func Value(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Value, value, condition...)
}

func Width(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.Width, value, condition...)
}

func EncType(value string, condition ...bool) gomponents.Node {
	return newConditionalAttributeWithValue(html.EncType, value, condition...)
}
