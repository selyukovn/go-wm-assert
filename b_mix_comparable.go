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

func (m *mixinComparable[A, T]) Eq(eq T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if v != eq {
			return mkCheckErr(
				fmt.Sprintf("value expects to be equal to %s, got %s", fmtVal(eq), fmtVal(v)),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

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

func (m *mixinComparable[A, T]) In(slice []T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		found := false
		for _, sv := range slice {
			if sv == v {
				found = true
				break
			}
		}
		if !found {
			return mkCheckErr(
				fmt.Sprintf("value expects to be in %s, got %s", fmtVal(slice), fmtVal(v)),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

func (m *mixinComparable[A, T]) NotIn(slice []T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		for _, sv := range slice {
			if sv == v {
				return mkCheckErr(
					fmt.Sprintf("value expects not to be in %s, got %s", fmtVal(slice), fmtVal(v)),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
