# Gomplements

[![GoDoc](https://pkg.go.dev/badge/github.com/willoma/gomplements)](https://pkg.go.dev/github.com/willoma/gomplements)

_Gomplements_ extends [https://www.gomponents.com/](Gomponents), a collection of HTML 5 components in pure Go, by providing an `Element` interface and its base implementation, as well as specific `Class` and `Styles` types.

It is suggested to import _Gomplements_ with the `e` (meaning "element") alias:

```go
import e "github.com/willoma/gomplements"
```

## Example

Simple example with _Gomplements_:

```go
component := e.Elem(
	html.Div,
	Class("my-component"),
	Class("blue"),
	"Hello world",
	e.ImgSrc("http://www.example.com/example.png"),
)
```

The equivalent using only _Gomponents_:

```go
component := h.Div(
	Class("my-component blue"),
	g.Text("Hello world"),
	h.Img(
		h.Src("http://www.example.com/example.png")
	),
)
```

## With

The _Gomplements_ `Element` allows adding children afterwards, using the `With` method:

```go
component := e.Elem(
	html.Div,
	Class("my-component"),
	"Hello world",
)

component.With(Class("blue"), e.ImgSrc("http://www.example.com/example.png"))
```

This makes it easier to build a component while reading some data in a loop, or by adding some children conditionally.

## Class and Styles

Use the `Class` type or implement `Classer` or `Classeser` in order to add CSS classes to an element:

```go
errorColor := e.Class("red")
successColor := e.Class("green")

type color struct {
	success bool
}

func (c *color) Class() Class {
	if c.success {
		return successColor
	} else {
		return errorColor
	}
}

type component struct {
	success bool
}

func (c *component) Classes() []Class {
	if c.success {
		return []Class{"myComponent", successColor}
	} else {
		return []Class{"myComponent", errorColor}
	}
}
```

You may also define specific CSS styles with the `Styles` map:

```go
e.Styles{
	"color": "red",
	"border": "1px solid blue",
}
```

## Parent modifier

Sometimes, you need to modify a parent element depending on a child. You may
implement the `ParentModifier` interface:

```go
type subcomponent struct {
	success bool
}

func (s *subcomponent) ModifyParent(parent e.Element) {
	parent.With(Class("hasMySubcomponent"))
	if s.success {
		parent.With(successColor)
	}
}
```

## HTML elements

HTML elements are already provided as functions that return `Element` instances,
which means the following are equivalent:

```go
component := e.Elem(html.Span, e.Class("my-component"))

component := e.Span(e.Class("my-component"))
```

## Helpers

_Gomplements_ also provides some helpers to simplify some common situations,
for instance `AHref` or `ImgSrc`. Check out the package doc!
