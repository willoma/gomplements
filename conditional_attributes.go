package gomplements

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type conditionalAttribute struct {
	attrFn func() gomponents.Node
	cond   bool
}

type conditionalAttributeWithValue struct {
	attrFn func(string) gomponents.Node
	value  string
	cond   bool
}

type conditionalAttributeWithNameAndValue struct {
	attrFn func(string, string) gomponents.Node
	name   string
	value  string
	cond   bool
}

func (b conditionalAttribute) ModifyParent(p Element) {
	if b.cond {
		p.With(b.attrFn())
	}
}

func (b conditionalAttributeWithValue) ModifyParent(p Element) {
	if b.cond {
		p.With(b.attrFn(b.value))
	}
}

func (b conditionalAttributeWithNameAndValue) ModifyParent(p Element) {
	if b.cond {
		p.With(b.attrFn(b.name, b.value))
	}
}

func Async(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Async,
		cond:   condition,
	}
}

func AutoFocus(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.AutoFocus,
		cond:   condition,
	}
}

func AutoPlay(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.AutoPlay,
		cond:   condition,
	}
}

func Checked(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Checked,
		cond:   condition,
	}
}

func Defer(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Defer,
		cond:   condition,
	}
}

func Disabled(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Disabled,
		cond:   condition,
	}
}

func Loop(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Loop,
		cond:   condition,
	}
}

func Multiple(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Multiple,
		cond:   condition,
	}
}

func Muted(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Muted,
		cond:   condition,
	}
}

func PlaysInline(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.PlaysInline,
		cond:   condition,
	}
}

func ReadOnly(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.ReadOnly,
		cond:   condition,
	}
}

func Required(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Required,
		cond:   condition,
	}
}

func Selected(condition bool) *conditionalAttribute {
	return &conditionalAttribute{
		attrFn: html.Selected,
		cond:   condition,
	}
}

func Accept(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Accept,
		value:  v,
		cond:   condition,
	}
}

func Action(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Action,
		value:  v,
		cond:   condition,
	}
}

func Alt(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Alt,
		value:  v,
		cond:   condition,
	}
}

func As(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.As,
		value:  v,
		cond:   condition,
	}
}

func AutoComplete(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.AutoComplete,
		value:  v,
		cond:   condition,
	}
}

func Charset(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Charset,
		value:  v,
		cond:   condition,
	}
}

func Cols(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Cols,
		value:  v,
		cond:   condition,
	}
}

func ColSpan(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.ColSpan,
		value:  v,
		cond:   condition,
	}
}

func Content(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Content,
		value:  v,
		cond:   condition,
	}
}

func DataAttr(name, v string, condition bool) *conditionalAttributeWithNameAndValue {
	return &conditionalAttributeWithNameAndValue{
		attrFn: html.DataAttr,
		name:   name,
		value:  v,
		cond:   condition,
	}
}

func For(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.For,
		value:  v,
		cond:   condition,
	}
}

func FormAttr(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.FormAttr,
		value:  v,
		cond:   condition,
	}
}

func Height(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Height,
		value:  v,
		cond:   condition,
	}
}

func Href(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Href,
		value:  v,
		cond:   condition,
	}
}

func Lang(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Lang,
		value:  v,
		cond:   condition,
	}
}

func Loading(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Loading,
		value:  v,
		cond:   condition,
	}
}

func Max(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Max,
		value:  v,
		cond:   condition,
	}
}

func MaxLength(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.MaxLength,
		value:  v,
		cond:   condition,
	}
}

func Method(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Method,
		value:  v,
		cond:   condition,
	}
}

func Min(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Min,
		value:  v,
		cond:   condition,
	}
}

func MinLength(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.MinLength,
		value:  v,
		cond:   condition,
	}
}

func Name(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Name,
		value:  v,
		cond:   condition,
	}
}

func Pattern(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Pattern,
		value:  v,
		cond:   condition,
	}
}

func Placeholder(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Placeholder,
		value:  v,
		cond:   condition,
	}
}

func Poster(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Poster,
		value:  v,
		cond:   condition,
	}
}

func Preload(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Preload,
		value:  v,
		cond:   condition,
	}
}

func Rel(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Rel,
		value:  v,
		cond:   condition,
	}
}

func Rows(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Rows,
		value:  v,
		cond:   condition,
	}
}

func RowSpan(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.RowSpan,
		value:  v,
		cond:   condition,
	}
}

func Src(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Src,
		value:  v,
		cond:   condition,
	}
}

func SrcSet(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.SrcSet,
		value:  v,
		cond:   condition,
	}
}

func Step(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Step,
		value:  v,
		cond:   condition,
	}
}

func TabIndex(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.TabIndex,
		value:  v,
		cond:   condition,
	}
}

func Target(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Target,
		value:  v,
		cond:   condition,
	}
}

func TitleAttr(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.TitleAttr,
		value:  v,
		cond:   condition,
	}
}

func Type(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Type,
		value:  v,
		cond:   condition,
	}
}

func Value(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Value,
		value:  v,
		cond:   condition,
	}
}

func Width(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.Width,
		value:  v,
		cond:   condition,
	}
}

func EncType(v string, condition bool) *conditionalAttributeWithValue {
	return &conditionalAttributeWithValue{
		attrFn: html.EncType,
		value:  v,
		cond:   condition,
	}
}
