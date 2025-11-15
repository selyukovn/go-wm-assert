package assert

import (
	"fmt"
)

type mixinComparable[A assertInterface[T], T comparable] struct {
	assert A
}

func newMixinComparable[A assertInterface[T], T comparable](assert A) *mixinComparable[A, T] {
	return &mixinComparable[A, T]{assert: assert}
}

// ---------------------------------------------------------------------------------------------------------------------
// Eq
// ---------------------------------------------------------------------------------------------------------------------

// Eq
//
// Value expects to be equal to "eq".
func (m *mixinComparable[A, T]) Eq(eq T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if v == eq {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf("value expects to be equal to %s, got %s", fmtVal(eq), fmtVal(v)),
			customErrMsg,
		)
	})
	return m.assert
}

// NotEq
//
// Value expects to be not equal to "notEq".
func (m *mixinComparable[A, T]) NotEq(notEq T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if v == notEq {
			return mkCheckErr(
				fmt.Sprintf("value expects to be not equal to %s, got %s", fmtVal(notEq), fmtVal(v)),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
// In
// ---------------------------------------------------------------------------------------------------------------------

// In
//
// Value expects to be equal to any of the provided elements.
//
// Fails check, if no elements provided.
func (m *mixinComparable[A, T]) In(slice []T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if len(slice) > 0 {
			for _, sv := range slice {
				if sv == v {
					return nil
				}
			}
		}
		return mkCheckErr(
			fmt.Sprintf("value expects to be in %s, got %s", fmtVal(slice), fmtVal(v)),
			customErrMsg,
		)
	})
	return m.assert
}

// NotIn
//
// Value expects to be not equal to each of the provided elements.
//
// Passes check, if no elements provided.
func (m *mixinComparable[A, T]) NotIn(slice []T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if len(slice) == 0 {
			return nil
		}
		for _, sv := range slice {
			if sv == v {
				return mkCheckErr(
					fmt.Sprintf("value expects to be not in %s, got %s", fmtVal(slice), fmtVal(v)),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
