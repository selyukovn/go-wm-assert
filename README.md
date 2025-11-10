# Assertions

### TL;DR

Simple, type‑specific fluent assertions for Go --
initially inspired by [php webmozart/assert](https://github.com/webmozart/assert).

### Here's the thing

If you’ve ever worked with php, you probably used [webmozart/assert](https://github.com/webmozart/assert) package.
It made it easy to validate method arguments with supplementary criteria in addition to type checking --
for example, verifying that integers were natural or strings were non‑empty.
Moreover, combinations of such checks could be used to build more complex validations
instead of using heavy validation frameworks that might impose compromises on project design.

In Go, there are already well‑known validation packages, but most of them are designed for broad or complex cases,
such as struct validation by tags, validation with network requests, etc.
This also can be overkill, so something else is needed...

Something simpler...
Something more type-specific...
Something more friendly...
Something like this package! :)

So yes -- this package ~~un autre package de validation, mais avec les cartes et les femmes fatales~~
provides a type-specific fluent interface to build validations in a way like
`assert.String().Word().LenMax(5).Check(value)`.

As is usual in Go, it supports two types of results:
panic (via the [`Must()`](b_assert.go) or [`MustAll()`]((b_assert.go)) methods)
and returning errors (via the [`Check()`](b_assert.go) or [`CheckAll()`](b_assert.go) methods).

The package provides popular type‑specific assertions --
such as [`String()`](s_string.go), [`Numeric[T NumericTypes]()`](s_numeric.go), [`Time()`](s_time.go), etc. --
as well as the ability to work with more abstract assertions like [`Comparable[T comparable]()`](s_comparable.go) 
for broader type support.

Each assertion also supports a [`Custom()`](b_mix_custom.go) check for specific cases.

### [Examples](readme_test.go)

```go
package example

import "github.com/selyukovn/go-wm-assert"

func NameFromString(value string) (Name, error) {
	err := assert.String().
		Word().
		LenMax(100, "Too long, isn’t it?" /* <-- custom error message only for this fail */).
		NotIn([]string{"some", "bad", "words"}).
		Custom(func(value string) error {
			// if ... { return err }
			return nil
		}).
		Check(value, "Name is incorrect" /* <-- custom error message that overrides any other in the chain */)

	if err != nil {
		return Name{}, err
	}

	return Name{value: value}, nil
}
```

```go
package example

import (
	"github.com/selyukovn/go-wm-assert"
	"time"
)

type EventCollection struct{ /* ... */ }

func (a *Account) Deactivate(now time.Time, evs *EventCollection) error {
	assert.Time().NotZero().Must(now)
	assert.Comparable[*EventCollection]().NotEq(nil).Must(evs)

	if a.IsDeactivated() {
		return common.NewErrorAlreadyDone()
	}

	a.deactivatedAt = now

	evs.Add(NewEventDeactivated(now, a.id))

	return nil
}
```
