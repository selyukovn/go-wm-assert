# Assertions

## TL;DR

Simple, type‑specific fluent assertions for Go
(initially inspired by [php webmozart/assert](https://github.com/webmozart/assert)).

Requires Go (1.18) or later.

```
assert.Str().Word().LenMax(5).Check(value)
```

## Reasons

If you’ve ever worked with php, you probably used [webmozart/assert](https://github.com/webmozart/assert) package.
It made it easy to validate method arguments with supplementary criteria in addition to type checking --
for example, verifying that integers were natural or strings were non‑empty.
Moreover, combinations of such checks could be used to build more complex validations
instead of using heavy validation frameworks that might impose compromises on project design.

In Go, there are already well‑known validation packages, but most of them are designed for broad or complex cases,
such as struct validation by tags, validation with network requests, etc. -- this also can be overkill.

Some Go-packages also have not enough friendly interface and weak typing.
For example, it can be possible to add length checks to an integer value validation flow, etc.
Such things confuse, make interface more complex and increase chances to make an accidental mistake.

So something else was needed...

Something _simpler_...
Something more _type-specific_...
Something more _friendly_...
Something like this package! :)

## Description

This package ~~un autre package de validation, mais avec les cartes et les femmes fatales~~
provides a simple type-specific fluent interface to build validations in a way like

```
assert.Str().Word().LenMax(5).Check(value)
```

The package provides assertions for specific types:

- [`Bool`](s_bool.go)
- [`Num`](s_num.go) -- for all _int_, _uint_ and _float_ based types (see [`NumericTypes`](s_num.go))
- [`Str`](s_str.go)
- [`Time`](s_time.go) -- for `time.Time` type
- [`TimeDur`](s_time_dur.go) -- for `time.Duration` type

as well as more abstract assertions for broader type support:

- [`Cmp`](s_cmp.go) -- for any comparable type
- [`SliceAny`](s_slice_any.go) -- for slice-based types with any type of elements
- [`SliceCmp`](s_slice_cmp.go) -- for slice-based types with comparable type of elements

Each assertion supports only specific methods related to its value type -- this protects you from accidental mistakes.
For example, it is impossible to validate integer value via some length-based rule,
because [`Num`](s_num.go) assertion does not have methods to add such validation to the chain.

Each assertion supports a [`Custom()`](b_mix_custom.go) check for cases not covered by built-ins.

Assertions support a few types of results:
- panic via the [`Must()`](b_assert.go) or [`MustAll()`]((b_assert.go)) methods
- panic or returning value via the [`MustGet()`](b_assert.go) or [`MustAllGet()`]((b_assert.go)) methods
- returning errors via the [`Check()`](b_assert.go) or [`CheckAll()`](b_assert.go) methods

Error messages can be customized for any rule as well as for the whole chain.

## [Examples](readme_test.go)

### Method arguments assertion (primary use-case)

This is the initial use-case of this package -- ensuring without ugly boilerplate code
that method works with valid values of the argument types, not with nil-pointers, etc.

```go
package example

import (
	"github.com/selyukovn/go-wm-assert"
	"time"
)

type Account struct{ /* ... */ }
type EventCollection struct{ /* ... */ }

func (a *Account) Deactivate(deactivatedAt time.Time, evs *EventCollection) error {
	assert.Time().NotZero().LessEq(time.Now()).Must(deactivatedAt)
	assert.Cmp[*EventCollection]().NotEq(nil).Must(evs)

	// ...

	return nil
}
```

### Simple validation (secondary use-case)

Looks the same argument assertion,
but customizing messages makes possible to use this package in outer layers, closer to end-user interaction code, etc.

```go
package example

import (
	"fmt"
	"github.com/selyukovn/go-wm-assert"
)

type Name struct{ value string }

func NameFromString(value string) (Name, error) {
	// custom error message that overrides any other in the chain */
	err := assert.Str().Word().Check(value, fmt.Sprintf("Name %q is incorrect!", value))

	if err != nil {
		return Name{}, err
	}

	return Name{value: value}, nil
}
```

### Form validation (possible use-case)

This package is not about a form validation,
but it can be used as a "brick" to build things you need without heavy validation libraries.

```go
package example

import "github.com/selyukovn/go-wm-assert"

type SignUpForm struct {
	email     string
	name      string // let it be optional
	age       uint
	agreement bool

	errors map[string][]error
}

// ...

func (f *SignUpForm) Validate() bool {
	f.errors = map[string][]error{
		"email":     {},
		"name":      {},
		"age":       {},
		"agreement": {},
	}

	f.errors["email"] = assert.Str().
		NotEmpty("Email is required!").
		Regexp(
			emailRegexpCompiled,

			// custom error message only for this rule
			// instead of technical "value ... regexp ..."
			"Email is incorrect!",
		).
		Custom(func(v string) error {
			// e.g. check that it is not registered previously
			return nil
		}).
		CheckAll(f.email)

	// optional field, remember?
	if f.name != "" {
		f.errors["name"] = assert.Str().
			Word("Only letters and '-' allowed!").
			RunesMin(2, "Too short, isn't it?").
			RunesMax(255, "Too long, isn't it?").
			NotIn(
				[]string{ /* e.g. some set of bad words or so */ },
				"Is that your real name, friend?",
			).
			CheckAll(f.name)
	}

	f.errors["age"] = assert.Num[uint]().
		GreaterEq(18, "Things are serious -- come back later!").
		Less(65, "Take a rest, friend!").
		CheckAll(f.age)

	f.errors["agreement"] = assert.Bool().
		True("This flag is required!").
		CheckAll(f.agreement)

	return len(f.errors["email"]) == 0 &&
		len(f.errors["name"]) == 0 &&
		len(f.errors["age"]) == 0 &&
		len(f.errors["agreement"]) == 0
}

func (f *SignUpForm) NameErrors() []error {
	return f.errors["name"]
}

// ...

```

## Package Structure

- `b_*.go` — basic components
- `s_*.go` — specific assertions (`Str`, `Num`, etc.)
- `readme_test.go` — examples from the Readme to be sure they really work :)
