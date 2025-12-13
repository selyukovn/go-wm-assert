package assert

import "fmt"

// #####################################################################################################################
// SLICE TYPE
// #####################################################################################################################

type sliceType[E any] interface {
	~[]E
}

// #####################################################################################################################
// ANY
// #####################################################################################################################

type mixinSliceAny[A assertInterface[S], S sliceType[E], E any] struct {
	*mixinLen[A, S]
}

func newMixinSliceAny[A assertInterface[S], S sliceType[E], E any](assert A) *mixinSliceAny[A, S, E] {
	return &mixinSliceAny[A, S, E]{
		mixinLen: newMixinLen[A, S](assert),
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// Empty
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceAny[A, S, E]) Empty(customErrMsg ...string) A {
	m.assert.addCheck(func(v S) error {
		if len(v) == 0 {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf("value expects to be empty, got %s", fmtVal(v)),
			customErrMsg,
		)
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
// Not Empty
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceAny[A, S, E]) NotEmpty(customErrMsg ...string) A {
	m.assert.addCheck(func(v S) error {
		if len(v) == 0 {
			return mkCheckErr(
				fmt.Sprintf("value expects to be not empty, got %s", fmtVal(v)),
				customErrMsg,
			)
		}
		return nil
	})
	return m.assert
}

// ---------------------------------------------------------------------------------------------------------------------
// Match
// ---------------------------------------------------------------------------------------------------------------------

// Any
// ---------------------------------------------------------------------------------------------------------------------

// CustomElementAny
//
// Expects the slice with any element matched to custom condition.
//
// Passes check, if the slice is empty.
// If it confuses, just add the NotEmpty() rule to the chain.
func (m *mixinSliceAny[A, S, E]) CustomElementAny(
	conditionName string,
	conditionFn func(e E) bool,
	customErrMsg ...string,
) A {
	m.assert.addCheck(func(v S) error {
		if len(v) == 0 {
			return nil
		}
		for _, e := range v {
			if conditionFn(e) {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects any element to match %s condition, got none matched in %s",
				fmtVal(conditionName),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return m.assert
}

// Each
// ---------------------------------------------------------------------------------------------------------------------

// CustomElementEach
//
// Expects the slice with each element matched to custom condition.
//
// Passes check, if the slice is empty.
// If it confuses, just add the NotEmpty() rule to the chain.
func (m *mixinSliceAny[A, S, E]) CustomElementEach(
	conditionName string,
	conditionFn func(e E) bool,
	customErrMsg ...string,
) A {
	m.assert.addCheck(func(v S) error {
		if len(v) == 0 {
			return nil
		}
		for _, e := range v {
			if !conditionFn(e) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects each element to match %s condition, got at least %s not matched in %s",
						fmtVal(conditionName),
						fmtVal(e),
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

// None
// ---------------------------------------------------------------------------------------------------------------------

// CustomElementNone
//
// Expects the slice with none element matched to custom condition.
//
// Passes check, if the slice is empty.
// If it confuses, just add the NotEmpty() rule to the chain.
func (m *mixinSliceAny[A, S, E]) CustomElementNone(
	conditionName string,
	conditionFn func(e E) bool,
	customErrMsg ...string,
) A {
	m.assert.addCheck(func(v S) error {
		if len(v) == 0 {
			return nil
		}
		for _, e := range v {
			if conditionFn(e) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects none element to match %s condition, got at least %s matched in %s",
						fmtVal(conditionName),
						fmtVal(e),
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

// #####################################################################################################################
// COMPARABLE
// #####################################################################################################################

type mixinSliceCmp[A assertInterface[S], S sliceType[E], E comparable] struct {
	*mixinSliceAny[A, S, E]
}

func newMixinSliceCmp[A assertInterface[S], S sliceType[E], E comparable](assert A) *mixinSliceCmp[A, S, E] {
	return &mixinSliceCmp[A, S, E]{
		mixinSliceAny: newMixinSliceAny[A, S, E](assert),
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// Contains
// ---------------------------------------------------------------------------------------------------------------------

// Contains
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceCmp[A, S, E]) Contains(e E, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		for _, ev := range v {
			if ev == e {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf("value expects to contain %s, got %s", fmtVal(e), fmtVal(v)),
			customErrMsg,
		)
	})
	return m.mixinSliceAny.assert
}

// Not Contains
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceCmp[A, S, E]) NotContains(e E, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		for _, ev := range v {
			if ev == e {
				return mkCheckErr(
					fmt.Sprintf("value expects to not contain %s, got %s", fmtVal(e), fmtVal(v)),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.mixinSliceAny.assert
}

// Any
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceCmp[A, S, E]) ContainsAny(s S, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		mv := make(map[E]struct{}, len(v))
		for _, ev := range v {
			mv[ev] = struct{}{}
		}
		for _, es := range s {
			if _, ok := mv[es]; ok {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf("value expects to contain any of %s, got %s", fmtVal(s), fmtVal(v)),
			customErrMsg,
		)
	})
	return m.mixinSliceAny.assert
}

// Each
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceCmp[A, S, E]) ContainsEach(s S, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		mv := make(map[E]struct{}, len(v))
		for _, ev := range v {
			mv[ev] = struct{}{}
		}
		for _, es := range s {
			if _, ok := mv[es]; !ok {
				return mkCheckErr(
					fmt.Sprintf("value expects to contain each of %s, got %s", fmtVal(s), fmtVal(v)),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.mixinSliceAny.assert
}

// None
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceCmp[A, S, E]) ContainsNone(s S, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		mv := make(map[E]struct{}, len(v))
		for _, ev := range v {
			mv[ev] = struct{}{}
		}
		for _, es := range s {
			if _, ok := mv[es]; ok {
				return mkCheckErr(
					fmt.Sprintf("value expects to contain none of %s, got %s", fmtVal(s), fmtVal(v)),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return m.mixinSliceAny.assert
}

// ---------------------------------------------------------------------------------------------------------------------
// Uniques
// ---------------------------------------------------------------------------------------------------------------------

// Uniques
//
// Expects the slice to have no duplicated elements -- i.e. expects that all elements of the slice are unique.
//
// Passes check, if the slice is empty.
// If it confuses, just add the NotEmpty() rule to the chain.
func (m *mixinSliceCmp[A, S, E]) Uniques(customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		if len(v) == 0 {
			return nil
		}
		me := make(map[E]struct{})
		for _, e := range v {
			if _, ok := me[e]; ok {
				return mkCheckErr(
					fmt.Sprintf("value expects all elements to be unique, got %s", fmtVal(v)),
					customErrMsg,
				)
			}
			me[e] = struct{}{}
		}
		return nil
	})
	return m.mixinSliceAny.assert
}

// ---------------------------------------------------------------------------------------------------------------------
// Uniques Len
// ---------------------------------------------------------------------------------------------------------------------

func (m *mixinSliceCmp[A, S, E]) uniquesLen(v S) int {
	me := make(map[E]struct{})
	for _, e := range v {
		me[e] = struct{}{}
	}
	return len(me)
}

// Equal
// ---------------------------------------------------------------------------------------------------------------------

// UniquesLenEq
//
// Length of unique elements sub-slice expects to be equal to "eq".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinSliceCmp[A, S, E]) UniquesLenEq(eq int, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		l := m.uniquesLen(v)
		if eq == l {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of unique elements sub-slice of %s expects to be equal to %s, got %s",
				fmtVal(v),
				fmtVal(eq),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.mixinSliceAny.assert
}

// Not Equal
// ---------------------------------------------------------------------------------------------------------------------

// UniquesLenNotEq
//
// Length of unique elements sub-slice expects to be not equal to "notEq".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinSliceCmp[A, S, E]) UniquesLenNotEq(notEq int, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		l := m.uniquesLen(v)
		if notEq == l {
			return mkCheckErr(
				fmt.Sprintf(
					"length of unique elements sub-slice of %s expects to be not equal to %s, got %s",
					fmtVal(v),
					fmtVal(notEq),
					fmtVal(l),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.mixinSliceAny.assert
}

// Min
// ---------------------------------------------------------------------------------------------------------------------

// UniquesLenMin
//
// Length of unique elements sub-slice expects to be greater or equal to "min".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinSliceCmp[A, S, E]) UniquesLenMin(min int, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		l := m.uniquesLen(v)
		if min <= l {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of unique elements sub-slice of %s expects to be greater or equal to %s, got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.mixinSliceAny.assert
}

// Max
// ---------------------------------------------------------------------------------------------------------------------

// UniquesLenMax
//
// Length of unique elements sub-slice expects to be less or equal to "max".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinSliceCmp[A, S, E]) UniquesLenMax(max int, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		l := m.uniquesLen(v)
		if l <= max {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of unique elements sub-slice of %s expects to be less or equal to %s, got %s",
				fmtVal(v),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.mixinSliceAny.assert
}

// In Range
// ---------------------------------------------------------------------------------------------------------------------

// UniquesLenInRange
//
// Length of unique elements sub-slice expects to be in range [min, max].
//
// Fails check, if min > max -- it works like empty range.
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinSliceCmp[A, S, E]) UniquesLenInRange(min, max int, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		l := m.uniquesLen(v)
		if min <= max {
			if min <= l && l <= max {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"length of unique elements sub-slice of %s expects to be in range [%s, %s], got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return m.mixinSliceAny.assert
}

// Not In Range
// ---------------------------------------------------------------------------------------------------------------------

// UniquesLenNotInRange
//
// Length of unique elements sub-slice expects to be not in range [min, max] --
// i.e. to be in ranges [0, min) or (max, MaxInt].
//
// Passes check, if min > max -- it works like empty range.
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (m *mixinSliceCmp[A, S, E]) UniquesLenNotInRange(min, max int, customErrMsg ...string) A {
	m.mixinSliceAny.assert.addCheck(func(v S) error {
		if min > max {
			return nil
		}
		l := m.uniquesLen(v)
		if min <= l && l <= max {
			return mkCheckErr(
				fmt.Sprintf(
					"length of unique elements sub-slice of %s expects to be not in range [%s, %s], got %s",
					fmtVal(v),
					fmtVal(min),
					fmtVal(max),
					fmtVal(l),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return m.mixinSliceAny.assert
}

// #####################################################################################################################
