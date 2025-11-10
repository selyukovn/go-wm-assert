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

func (m *mixinOrdered[A, T]) less(eq bool, than T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		if !(eq && (m.fnCmp(than, v) || v == than) || !eq && m.fnCmp(than, v)) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to be less %s %s, got %s",
					ternary[string](eq, "or equal than", "than"),
					fmtVal(than),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

func (m *mixinOrdered[A, T]) Less(than T, customErrMsg ...string) A {
	return m.less(false, than, customErrMsg)
}

func (m *mixinOrdered[A, T]) LessEq(than T, customErrMsg ...string) A {
	return m.less(true, than, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) lessAny(eq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		found := false
		for _, t := range elems {
			if (eq && v == t) || m.fnCmp(t, v) {
				found = true
				break
			}
		}
		if !found {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to be less %s any of %s, got %s",
					ternary[string](eq, "or equal than", "than"),
					fmtVal(elems),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

func (m *mixinOrdered[A, T]) LessAny(elems []T, customErrMsg ...string) A {
	return m.lessAny(false, elems, customErrMsg)
}

func (m *mixinOrdered[A, T]) LessEqAny(elems []T, customErrMsg ...string) A {
	return m.lessAny(true, elems, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) lessEach(eq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		for _, t := range elems {
			if !((eq && v == t) || m.fnCmp(t, v)) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to be less %s each of %s, got %s",
						ternary[string](eq, "or equal than", "than"),
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

func (m *mixinOrdered[A, T]) LessEach(elems []T, customErrMsg ...string) A {
	return m.lessEach(false, elems, customErrMsg)
}

func (m *mixinOrdered[A, T]) LessEqEach(elems []T, customErrMsg ...string) A {
	return m.lessEach(true, elems, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
// Greater
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinOrdered[A, T]) greater(eq bool, than T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		if !(eq && (m.fnCmp(v, than) || than == v) || !eq && m.fnCmp(v, than)) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to be greater %s %s, got %s",
					ternary[string](eq, "or equal than", "than"),
					fmtVal(than),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

func (m *mixinOrdered[A, T]) Greater(than T, customErrMsg ...string) A {
	return m.greater(false, than, customErrMsg)
}

func (m *mixinOrdered[A, T]) GreaterEq(than T, customErrMsg ...string) A {
	return m.greater(true, than, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) greaterAny(eq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		found := false
		for _, t := range elems {
			if (eq && v == t) || m.fnCmp(v, t) {
				found = true
				break
			}
		}
		if !found {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to be greater %s any of %s, got %s",
					ternary[string](eq, "or equal than", "than"),
					fmtVal(elems),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

func (m *mixinOrdered[A, T]) GreaterAny(elems []T, customErrMsg ...string) A {
	return m.greaterAny(false, elems, customErrMsg)
}

func (m *mixinOrdered[A, T]) GreaterEqAny(elems []T, customErrMsg ...string) A {
	return m.greaterAny(true, elems, customErrMsg)
}

// ----

func (m *mixinOrdered[A, T]) greaterEach(eq bool, elems []T, customErrMsg []string) A {
	m.assert.addCheck(func(v T) error {
		for _, t := range elems {
			if !((eq && v == t) || m.fnCmp(v, t)) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to be greater %s each of %s, got %s",
						ternary[string](eq, "or equal than", "than"),
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

func (m *mixinOrdered[A, T]) GreaterEach(elems []T, customErrMsg ...string) A {
	return m.greaterEach(false, elems, customErrMsg)
}

func (m *mixinOrdered[A, T]) GreaterEqEach(elems []T, customErrMsg ...string) A {
	return m.greaterEach(true, elems, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
// Range
// ---------------------------------------------------------------------------------------------------------------------

// The "min <= max" rule check is unnecessary here, since it doesn't affect the consistency of the internal state.
// Panics or errors won't simplify the method interface,
// but won't provide any benefit to either the current library or to the client code.

func (m *mixinOrdered[A, T]) inRangeCond(v, min, max T) bool {
	return (m.fnCmp(v, min) || min == v) && (m.fnCmp(max, v) || v == max)
}

func (m *mixinOrdered[A, T]) InRange(min T, max T, customErrMsg ...string) A {
	m.assert.addCheck(func(v T) error {
		if !m.inRangeCond(v, min, max) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to be in range [%s, %s], got %s",
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
