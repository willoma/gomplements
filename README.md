# Gomplements

[![GoDoc](https://pkg.go.dev/badge/github.com/willoma/gomplements)](https://pkg.go.dev/github.com/willoma/gomplements)

_Gomplements_ extends [Gomponents](https://www.gomponents.com/), a collection of HTML 5 components in pure Go, by providing an `Element` interface and its base implementation, as well as specific `Class` and `Styles` types.

It is suggested to import _Gomplements_ with the `e` (meaning "element") alias:

```go
import e "github.com/willoma/gomplements"
```

## Example

Simple example with _Gomplements_:

```go
component := e.Div(
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
component := e.Div(
	Class("my-component"),
	"Hello world",
)

component.With(Class("blue"), e.ImgSrc("http://www.example.com/example.png"))
```

This makes it easier to build a component while reading some data in a loop, or by adding some children conditionally.

## Class and Styles

Option 1: use the `Class` type in order to add a single CSS class to an element:

```go
errorColor := e.Class("red")
successColor := e.Class("green")
```

Option 2: use the `Classes` type in order to add multiple CSS classes to an element:

```go
errorStatus := e.Classes{"red", "bold"}
successStatus := e.Classes{"green", "italic"}
```

Option 3: implement `Classer` in order to add a single CSS class to an element:

```go
type color struct {
	success bool
}

func (c *color) Class() e.Class {
	if c.success {
		return e.Class("green")
	} else {
		return e.Class("red")
	}
}
```

Option 4: implement `Classeser` in order to add multiple CSS classes to an element:

```go
type status struct {
	success bool
}

func (c *color) Classes() []e.Class {
	if c.success {
		return []e.Class{"green", "italic"}
	} else {
		return []e.Class{"red", "bold"}
	}
}
```

Option 5: define specific CSS styles with the `Styles` map:

```go
e.Styles{
	"color": "red",
	"border": "1px solid blue",
}
```

Please note the _Gomplements_ classes and styles do not work on traditional `gomponents.Node` elements, only on _Gomplements_' `Element`.

## Parent modifier

Sometimes, you need to modify a parent element depending on a child. You may
implement the `ParentModifier` interface:

```go
type subcomponent struct {
	success bool
}

func (s *subcomponent) ModifyParent(parent e.Element) {
	parent.With(Class("has-my-subcomponent"))
	if s.success {
		parent.With(successColor)
	}
}
```

## HTML elements

HTML elements are already provided as functions that return `Element` instances, which means the following are equivalent:

```go
component := e.Elem(html.Span, e.Class("my-component"))

component := e.Span(e.Class("my-component"))
```

## Conditional attributes

In addition to regular attributes from `github.com/maragudk/gomponents/html`, you may use their equivalent in `e`, which allow appending a boolean indicating if the attribute must be included. It helps avoiding `gomponents.If`, the following are equivalent:

```go
component := e.Span(
	e.Class("my-component"),
	html.Disabled(),
	gomponents.If(condition1, html.Selected()),
	gomponents.If(condition2, html.Value("42")),
)

component := e.Span(
	e.Class("my-component"),
	e.Disabled(),
	e.Selected(condition1),
	e.Value("42", condition2),
)
```

If the condition is false, the attribute is not built and nil is returned.

## Conditional element

If you need to add an element conditionally, you may use `e.If`. If its condition argument is false, the element is not built and nil is returned.

For example:

```go
component := e.Span(
	e.Class("my-component"),
	e.If(condition3, e.Div, e.Class("my-content"), "Something"),
)
```

Here, `e.Div` is executed (built) only if `condition3` is true.

However, there is a limitation: this function only works if the element builder takes only one `any` variadic as its argument(s). It would not work with the `e.AHref` helper, for example.

## Conditional classes and styles

If yoyu need to add a class to an element conditionally, call the `.If(bool)` method on that class. If the condition is false, the class is not added to the element. This method is also available on the `Classes` and `Styles` types.

For example:

```go
func myComponent(alert bool) e.Element {
	return e.Div(
		e.Class("my-component"),
		e.Class("alert").If(alert),
	)
}
```

## Helpers

_Gomplements_ also provides some helpers to simplify some common situations,
for instance `AHref` or `ImgSrc`. Check out the package reference!

## ARIA

_Gomplements_ includes a collection of ARIA helpers, directly usable as children
on elements in order to define ARIA roles or attributes.

For example:

```go
component := e.Span(
	e.Class("my-component"),
	e.AriaToolbar,
	e.AriaLabel("main tools"),
)
```
