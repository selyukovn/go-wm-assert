package assert

import "fmt"

// ---------------------------------------------------------------------------------------------------------------------

type mixinOrdered[A assertInterface[T], T comparable] struct {
	assert A
	fnCmp  func(bigger, smaller T) bool
}

func newMixinOrdered[A assertInterface[T], T comparable](
	assert A,
	fnCmp func(bigger, smaller T) bool,
) *mixinOrdered[A, T] {
	if fnCmp == nil {
		panic(fmt.Errorf("newMixinOrdered expects not nil fnCmp"))
	}

	return &mixinOrdered[A, T]{
		assert: assert,
		fnCmp:  fnCmp,
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// Less
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinOrdered[A, T]) less(orEq bool, than T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		if m.fnCmp(than, v) || (orEq && v == than) {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to be less %s %s, got %s",
				ternary[string](orEq, "or equal than", "than"),
				fmtVal(than),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// Less
//
// Value expects to be less than "than".
func (m *mixinOrdered[A, T]) Less(than T, customErrMsg ...string) A {
	return m.less(false, than, customErrMsg)
}

// LessEq
//
// Value expects to be less or equal to "than".
func (m *mixinOrdered[A, T]) LessEq(than T, customErrMsg ...string) A {
	return m.less(true, than, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) lessAny(orEq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		for _, t := range elems {
			if m.fnCmp(t, v) || (orEq && v == t) {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to be less %s any of %s, got %s",
				ternary[string](orEq, "or equal than", "than"),
				fmtVal(elems),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// LessAny
//
// Value expects to be less than any of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) LessAny(elems []T, customErrMsg ...string) A {
	return m.lessAny(false, elems, customErrMsg)
}

// LessEqAny
//
// Value expects to be less or equal to any of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) LessEqAny(elems []T, customErrMsg ...string) A {
	return m.lessAny(true, elems, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) lessEach(orEq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		for _, t := range elems {
			if !(m.fnCmp(t, v) || (orEq && v == t)) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to be less %s each of %s, got %s",
						ternary[string](orEq, "or equal than", "than"),
						fmtVal(elems),
						fmtVal(v),
					),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.assert
}

// LessEach
//
// Value expects to be less than each of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) LessEach(elems []T, customErrMsg ...string) A {
	return m.lessEach(false, elems, customErrMsg)
}

// LessEqEach
//
// Value expects to be less or equal to each of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) LessEqEach(elems []T, customErrMsg ...string) A {
	return m.lessEach(true, elems, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
// Greater
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinOrdered[A, T]) greater(orEq bool, than T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		if m.fnCmp(v, than) || (orEq && v == than) {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to be greater %s %s, got %s",
				ternary[string](orEq, "or equal than", "than"),
				fmtVal(than),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// Greater
//
// Value expects to be greater than "than".
func (m *mixinOrdered[A, T]) Greater(than T, customErrMsg ...string) A {
	return m.greater(false, than, customErrMsg)
}

// GreaterEq
//
// Value expects to be greater or equal to "than".
func (m *mixinOrdered[A, T]) GreaterEq(than T, customErrMsg ...string) A {
	return m.greater(true, than, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) greaterAny(orEq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		for _, t := range elems {
			if m.fnCmp(v, t) || (orEq && v == t) {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to be greater %s any of %s, got %s",
				ternary[string](orEq, "or equal than", "than"),
				fmtVal(elems),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// GreaterAny
//
// Value expects to be greater than any of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) GreaterAny(elems []T, customErrMsg ...string) A {
	return m.greaterAny(false, elems, customErrMsg)
}

// GreaterEqAny
//
// Value expects to be greater or equal to any of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) GreaterEqAny(elems []T, customErrMsg ...string) A {
	return m.greaterAny(true, elems, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) greaterEach(orEq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		for _, t := range elems {
			if !(m.fnCmp(v, t) || (orEq && v == t)) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to be greater %s each of %s, got %s",
						ternary[string](orEq, "or equal than", "than"),
						fmtVal(elems),
						fmtVal(v),
					),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.assert
}

// GreaterEach
//
// Value expects to be greater than each of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) GreaterEach(elems []T, customErrMsg ...string) A {
	return m.greaterEach(false, elems, customErrMsg)
}

// GreaterEqEach
//
// Value expects to be greater or equal to each of provided elements.
//
// Passes check, if no elements provided.
func (m *mixinOrdered[A, T]) GreaterEqEach(elems []T, customErrMsg ...string) A {
	return m.greaterEach(true, elems, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
// Range
//
// The "min <= max" rule check is unnecessary here, since it doesn't affect the consistency of the internal state.
// Panics or errors won't simplify the method interface,
// but won't provide any benefit to either the current library or to the client code.
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinOrdered[A, T]) inRangeCond(v, min, max T) bool {
	return (m.fnCmp(v, min) || min == v) && (m.fnCmp(max, v) || v == max)
}

// InRange
//
// Value expects to be in range [min, max].
//
// Fails check, if min > max.
func (m *mixinOrdered[A, T]) InRange(min T, max T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if m.inRangeCond(v, min, max) {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to be in range [%s, %s], got %s",
				fmtVal(min),
				fmtVal(max),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// NotInRange
//
// Value expects to be not in range [min, max] -- i.e. to be in ranges [PossibleMin, min) or (max, PossibleMax]
//
// Passes check, if min > max.
func (m *mixinOrdered[A, T]) NotInRange(min T, max T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if m.inRangeCond(v, min, max) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects not to be in range [%s, %s], got %s",
					fmtVal(min),
					fmtVal(max),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
