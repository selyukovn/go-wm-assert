package assert

import (
	"fmt"
	"reflect"
)

type mixinLen[A assertInterface[T], T any] struct {
	assert A
}

func newMixinLen[A assertInterface[T], T any](assert A) *mixinLen[A, T] {
	return &mixinLen[A, T]{assert: assert}
}

// ---------------------------------------------------------------------------------------------------------------------
// Len
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinLen[A, T]) lenComparable() *mixinComparable[*assert[int], int] {
	return newMixinComparable[*assert[int], int](newAssert[int]())
}

func mixinLenFnCmp(bigger, smaller int) bool {
	return bigger > smaller
}

func (m *mixinLen[A, T]) lenOrdered() *mixinOrdered[*assert[int], int] {
	return newMixinOrdered[*assert[int], int](newAssert[int](), mixinLenFnCmp)
}

func (m *mixinLen[A, T]) lenVal(v T) int {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return rv.Len()
	default:
		panic(fmt.Errorf("%T : unsupportable type %T : unable to get len of value %v", m.assert, v, v))
	}
}

// LenEq
//
// Length expects to be equal to "eq".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinLen[A, T]) LenEq(eq int, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		l := m.lenVal(v)
		if m.lenComparable().Eq(eq).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of %s expects to be equal to %s, got %s",
				fmtVal(v),
				fmtVal(eq),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// LenNotEq
//
// Length expects to be not equal to "notEq".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinLen[A, T]) LenNotEq(notEq int, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		l := m.lenVal(v)
		if m.lenComparable().NotEq(notEq).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of %s expects to be not equal to %s, got %s",
				fmtVal(v),
				fmtVal(notEq),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// LenMin
//
// Length expects to be greater or equal to "min".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinLen[A, T]) LenMin(min int, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		l := m.lenVal(v)
		if m.lenOrdered().GreaterEq(min).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of %s expects to be greater or equal to %s, got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// LenMax
//
// Length expects to be less or equal to "max".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinLen[A, T]) LenMax(max int, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		l := m.lenVal(v)
		if m.lenOrdered().LessEq(max).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of %s expects to be less or equal to %s, got %s",
				fmtVal(v),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// LenInRange
//
// Length expects to be in range [min, max].
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinLen[A, T]) LenInRange(min, max int, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		l := m.lenVal(v)
		if m.lenOrdered().InRange(min, max).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of %s expects to be in range [%s, %s], got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// LenNotInRange
//
// Length expects to be not in range [min, max] -- i.e. to be in ranges [0, min) or (max, MaxInt].
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinLen[A, T]) LenNotInRange(min, max int, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		l := m.lenVal(v)
		if m.lenOrdered().NotInRange(min, max).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of %s expects to be not in range [%s, %s], got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
