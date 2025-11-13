package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"testing"
)

// #####################################################################################################################
// ANY
// #####################################################################################################################

func Test_MixinSliceAny(t *testing.T) {
	type anyType = map[int]struct{}

	type testAssert struct {
		*assert[[]anyType]
		*mixinSliceAny[*testAssert, []anyType, anyType]
	}
	fnNewAssert := func() *testAssert {
		a := new(testAssert)
		*a = testAssert{
			assert:        newAssert[[]anyType](),
			mixinSliceAny: newMixinSliceAny[*testAssert, []anyType, anyType](a),
		}
		return a
	}

	e1 := anyType{1: {}, 3: {}, 5: {}}
	e2 := anyType{7: {}, 9: {}, 11: {}}
	e3 := anyType{13: {}, 15: {}, 17: {}, 19: {}}
	e1FnMatch := func(e anyType) bool {
		_, ok := e[3]
		return ok
	}
	e1e2FnMatch := func(e anyType) bool {
		return len(e) == 3
	}

	// Empty
	// --------------------------------

	t.Run("Empty", func(t *testing.T) {
		a := fnNewAssert().Empty()

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]anyType{}))
		tAssert.NoError(t, a.Check(make([]anyType, 0)))
		tAssert.NoError(t, a.Check(make([]anyType, 0, 1)))
		tAssert.NoError(t, a.Check(make([]anyType, 0, 2)))

		tAssert.Error(t, a.Check([]anyType{e1}))
		tAssert.Error(t, a.Check(make([]anyType, 1)))
		tAssert.Error(t, a.Check(append(make([]anyType, 0), e1)))
		tAssert.Error(t, a.Check(append(make([]anyType, 0, 1), e1)))
		tAssert.Error(t, a.Check(append(make([]anyType, 0, 2), e1, e2)))

		tAssert.Equal(t, "e1", fnNewAssert().Empty("e1").Check(make([]anyType, 1)).Error())
		tAssert.Equal(t, "e2", fnNewAssert().Empty().Check(make([]anyType, 1), "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().Empty("e1").Check(make([]anyType, 1), "e2").Error())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		a := fnNewAssert().NotEmpty()

		tAssert.NoError(t, a.Check([]anyType{e1}))
		tAssert.NoError(t, a.Check([]anyType{e1, e2}))
		tAssert.NoError(t, a.Check(make([]anyType, 1)))
		tAssert.NoError(t, a.Check(append(make([]anyType, 0), e1)))
		tAssert.NoError(t, a.Check(append(make([]anyType, 0, 1), e1)))
		tAssert.NoError(t, a.Check(append(make([]anyType, 0, 2), e1, e2)))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]anyType{}))
		tAssert.Error(t, a.Check(make([]anyType, 0)))
		tAssert.Error(t, a.Check(make([]anyType, 0, 1)))
		tAssert.Error(t, a.Check(make([]anyType, 0, 2)))

		tAssert.Equal(t, "e1", fnNewAssert().NotEmpty("e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().NotEmpty().Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().NotEmpty("e1").Check(nil, "e2").Error())
	})

	// Match
	// --------------------------------

	t.Run("CustomElementAny", func(t *testing.T) {
		a := fnNewAssert().CustomElementAny("e1FnMatch", e1FnMatch)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]anyType{}))
		tAssert.NoError(t, a.Check([]anyType{e1}))
		tAssert.NoError(t, a.Check([]anyType{e1, e1}))
		tAssert.NoError(t, a.Check([]anyType{e1, e2}))
		tAssert.NoError(t, a.Check([]anyType{e2, e2, e1}))

		tAssert.Error(t, a.Check([]anyType{e2}))
		tAssert.Error(t, a.Check([]anyType{e2, e2}))
		tAssert.Error(t, a.Check([]anyType{e2, e2}))

		tAssert.Equal(t, "e1", fnNewAssert().CustomElementAny("e1FnMatch", e1FnMatch, "e1").Check([]anyType{e2}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().CustomElementAny("e1FnMatch", e1FnMatch).Check([]anyType{e2}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().CustomElementAny("e1FnMatch", e1FnMatch, "e1").Check([]anyType{e2}, "e2").Error())
	})

	t.Run("CustomElementEach", func(t *testing.T) {
		a := fnNewAssert().CustomElementEach("e1e2FnMatch", e1e2FnMatch)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]anyType{}))
		tAssert.NoError(t, a.Check([]anyType{e1}))
		tAssert.NoError(t, a.Check([]anyType{e2}))
		tAssert.NoError(t, a.Check([]anyType{e1, e2}))
		tAssert.NoError(t, a.Check([]anyType{e1, e1, e2}))
		tAssert.NoError(t, a.Check([]anyType{e1, e2, e2}))
		tAssert.NoError(t, a.Check([]anyType{e2, e2, e1, e1}))

		tAssert.Error(t, a.Check([]anyType{e3}))
		tAssert.Error(t, a.Check([]anyType{e1, e3}))
		tAssert.Error(t, a.Check([]anyType{e1, e2, e3}))

		tAssert.Equal(t, "e1", fnNewAssert().CustomElementEach("e1e2FnMatch", e1e2FnMatch, "e1").Check([]anyType{e3}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().CustomElementEach("e1e2FnMatch", e1e2FnMatch).Check([]anyType{e3}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().CustomElementEach("e1e2FnMatch", e1e2FnMatch, "e1").Check([]anyType{e3}, "e2").Error())
	})

	t.Run("CustomElementNone", func(t *testing.T) {
		a := fnNewAssert().CustomElementNone("e1FnMatch", e1FnMatch)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]anyType{}))
		tAssert.NoError(t, a.Check([]anyType{e2}))
		tAssert.NoError(t, a.Check([]anyType{e2, e2}))
		tAssert.NoError(t, a.Check([]anyType{e2, e2}))

		tAssert.Error(t, a.Check([]anyType{e1}))
		tAssert.Error(t, a.Check([]anyType{e1, e1}))
		tAssert.Error(t, a.Check([]anyType{e1, e2}))
		tAssert.Error(t, a.Check([]anyType{e2, e2, e1}))

		tAssert.Equal(t, "e1", fnNewAssert().CustomElementNone("e1FnMatch", e1FnMatch, "e1").Check([]anyType{e1}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().CustomElementNone("e1FnMatch", e1FnMatch).Check([]anyType{e1}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().CustomElementNone("e1FnMatch", e1FnMatch, "e1").Check([]anyType{e1}, "e2").Error())
	})
}

// #####################################################################################################################
// COMPARABLE
// #####################################################################################################################

func Test_MixinSliceCmp(t *testing.T) {
	type cmpType = int

	type testAssert struct {
		*assert[[]cmpType]
		*mixinSliceCmp[*testAssert, []cmpType, cmpType]
	}
	fnNewAssert := func() *testAssert {
		a := new(testAssert)
		*a = testAssert{
			assert:        newAssert[[]cmpType](),
			mixinSliceCmp: newMixinSliceCmp[*testAssert, []cmpType, cmpType](a),
		}
		return a
	}

	e1 := 1
	e2 := 2
	e3 := 3

	// Contains
	// --------------------------------

	t.Run("Contains", func(t *testing.T) {
		a := fnNewAssert().Contains(e1)

		tAssert.NoError(t, a.Check([]cmpType{e1}))
		tAssert.NoError(t, a.Check([]cmpType{e2, e1}))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]cmpType{}))
		tAssert.Error(t, a.Check([]cmpType{e2}))
		tAssert.Error(t, a.Check([]cmpType{e2, e3}))

		tAssert.Equal(t, "e1", fnNewAssert().Contains(e1, "e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().Contains(e1).Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().Contains(e1, "e1").Check(nil, "e2").Error())
	})

	t.Run("ContainsAny", func(t *testing.T) {
		a := fnNewAssert().ContainsAny([]cmpType{e2, e3})

		tAssert.NoError(t, a.Check([]cmpType{e2}))
		tAssert.NoError(t, a.Check([]cmpType{e3}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e2, e3}))
		tAssert.NoError(t, a.Check([]cmpType{e2, e3, e3, e2}))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]cmpType{}))
		tAssert.Error(t, a.Check([]cmpType{e1}))

		tAssert.Equal(t, "e1", fnNewAssert().ContainsAny([]cmpType{e2, e3}, "e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().ContainsAny([]cmpType{e2, e3}).Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().ContainsAny([]cmpType{e2, e3}, "e1").Check(nil, "e2").Error())
	})

	t.Run("ContainsEach", func(t *testing.T) {
		a := fnNewAssert().ContainsEach([]cmpType{e2, e3})

		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e3}))
		tAssert.NoError(t, a.Check([]cmpType{e2, e3}))
		tAssert.NoError(t, a.Check([]cmpType{e3, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e3, e1, e2, e1, e3}))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]cmpType{}))
		tAssert.Error(t, a.Check([]cmpType{e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2}))
		tAssert.Error(t, a.Check([]cmpType{e1, e3}))
		tAssert.Error(t, a.Check([]cmpType{e1, e3, e3, e1}))
		tAssert.Error(t, a.Check([]cmpType{e2}))
		tAssert.Error(t, a.Check([]cmpType{e3}))

		tAssert.Equal(t, "e1", fnNewAssert().ContainsEach([]cmpType{e2, e3}, "e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().ContainsEach([]cmpType{e2, e3}).Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().ContainsEach([]cmpType{e2, e3}, "e1").Check(nil, "e2").Error())
	})

	t.Run("NotContains", func(t *testing.T) {
		a := fnNewAssert().NotContains(e1)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]cmpType{}))
		tAssert.NoError(t, a.Check([]cmpType{e2}))
		tAssert.NoError(t, a.Check([]cmpType{e2, e3}))

		tAssert.Error(t, a.Check([]cmpType{e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2}))
		tAssert.Error(t, a.Check([]cmpType{e2, e1}))

		tAssert.Equal(t, "e1", fnNewAssert().NotContains(e1, "e1").Check([]cmpType{e1}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().NotContains(e1).Check([]cmpType{e1}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().NotContains(e1, "e1").Check([]cmpType{e1}, "e2").Error())
	})

	t.Run("ContainsNone", func(t *testing.T) {
		a := fnNewAssert().ContainsNone([]cmpType{e2, e3})

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]cmpType{}))
		tAssert.NoError(t, a.Check([]cmpType{e1}))

		tAssert.Error(t, a.Check([]cmpType{e2}))
		tAssert.Error(t, a.Check([]cmpType{e3}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2}))
		tAssert.Error(t, a.Check([]cmpType{e2, e3}))
		tAssert.Error(t, a.Check([]cmpType{e2, e3, e3, e2}))

		tAssert.Equal(t, "e1", fnNewAssert().ContainsNone([]cmpType{e2, e3}, "e1").Check([]cmpType{e2}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().ContainsNone([]cmpType{e2, e3}).Check([]cmpType{e2}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().ContainsNone([]cmpType{e2, e3}, "e1").Check([]cmpType{e2}, "e2").Error())
	})

	// Uniques
	// --------------------------------

	t.Run("Uniques", func(t *testing.T) {
		a := fnNewAssert().Uniques()

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]cmpType{}))
		tAssert.NoError(t, a.Check([]cmpType{e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e3}))

		tAssert.Error(t, a.Check([]cmpType{e1, e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e2}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e3, e1}))

		tAssert.Equal(t, "e1", fnNewAssert().Uniques("e1").Check([]cmpType{e1, e1}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().Uniques().Check([]cmpType{e1, e1}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().Uniques("e1").Check([]cmpType{e1, e1}, "e2").Error())
	})

	t.Run("UniquesLenEq", func(t *testing.T) {
		tAssert.Error(t, fnNewAssert().UniquesLenEq(-1).Check(nil))

		a := fnNewAssert().UniquesLenEq(2)

		tAssert.NoError(t, a.Check([]cmpType{e1, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e1, e1, e1}))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]cmpType{}))
		tAssert.Error(t, a.Check([]cmpType{e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e1, e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e3}))

		tAssert.Equal(t, "e1", fnNewAssert().UniquesLenEq(2, "e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenEq(2).Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenEq(2, "e1").Check(nil, "e2").Error())
	})

	t.Run("UniquesLenNotEq", func(t *testing.T) {
		tAssert.NoError(t, fnNewAssert().UniquesLenNotEq(-2).Check(nil))

		a := fnNewAssert().UniquesLenNotEq(2)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]cmpType{}))
		tAssert.NoError(t, a.Check([]cmpType{e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e3}))

		tAssert.Error(t, a.Check([]cmpType{e1, e2}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e1, e1, e1}))

		tAssert.Equal(t, "e1", fnNewAssert().UniquesLenNotEq(2, "e1").Check([]cmpType{e1, e2}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenNotEq(2).Check([]cmpType{e1, e2}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenNotEq(2, "e1").Check([]cmpType{e1, e2}, "e2").Error())
	})

	t.Run("UniquesLenMin", func(t *testing.T) {
		tAssert.NoError(t, fnNewAssert().UniquesLenMin(-2).Check(nil))

		a := fnNewAssert().UniquesLenMin(2)

		tAssert.NoError(t, a.Check([]cmpType{e1, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e1, e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e3}))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]cmpType{}))
		tAssert.Error(t, a.Check([]cmpType{e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e1, e1}))

		tAssert.Equal(t, "e1", fnNewAssert().UniquesLenMin(2, "e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenMin(2).Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenMin(2, "e1").Check(nil, "e2").Error())
	})

	t.Run("UniquesLenMax", func(t *testing.T) {
		tAssert.Error(t, fnNewAssert().UniquesLenMax(-2).Check(nil))

		a := fnNewAssert().UniquesLenMax(2)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]cmpType{}))
		tAssert.NoError(t, a.Check([]cmpType{e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e1, e1, e1}))

		tAssert.Error(t, a.Check([]cmpType{e1, e2, e3}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e3, e1, e2, e3, e3, e3}))

		tAssert.Equal(t, "e1", fnNewAssert().UniquesLenMax(2, "e1").Check([]cmpType{e1, e2, e3}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenMax(2).Check([]cmpType{e1, e2, e3}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenMax(2, "e1").Check([]cmpType{e1, e2, e3}, "e2").Error())
	})

	t.Run("UniquesLenInRange", func(t *testing.T) {
		tAssert.NoError(t, fnNewAssert().UniquesLenInRange(-2, 1).Check([]cmpType{e1}))
		tAssert.Error(t, fnNewAssert().UniquesLenInRange(-2, -1).Check([]cmpType{e1}))
		tAssert.Error(t, fnNewAssert().UniquesLenInRange(1, -2).Check([]cmpType{e1}))
		tAssert.Error(t, fnNewAssert().UniquesLenInRange(2, 1).Check([]cmpType{e1, e2}))

		a := fnNewAssert().UniquesLenInRange(1, 2)

		tAssert.NoError(t, a.Check([]cmpType{e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e1, e1}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e1, e1, e1}))

		tAssert.Error(t, a.Check(nil))
		tAssert.Error(t, a.Check([]cmpType{}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e3}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e3, e1, e2, e3, e3, e3}))

		tAssert.Equal(t, "e1", fnNewAssert().UniquesLenInRange(1, 2, "e1").Check(nil).Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenInRange(1, 2).Check(nil, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenInRange(1, 2, "e1").Check(nil, "e2").Error())
	})

	t.Run("UniquesLenNotInRange", func(t *testing.T) {
		tAssert.Error(t, fnNewAssert().UniquesLenNotInRange(-2, 1).Check([]cmpType{e1}))
		tAssert.NoError(t, fnNewAssert().UniquesLenNotInRange(-2, -1).Check([]cmpType{e1}))
		tAssert.NoError(t, fnNewAssert().UniquesLenNotInRange(-1, -2).Check([]cmpType{e1}))

		a := fnNewAssert().UniquesLenNotInRange(1, 2)

		tAssert.NoError(t, a.Check(nil))
		tAssert.NoError(t, a.Check([]cmpType{}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e3}))
		tAssert.NoError(t, a.Check([]cmpType{e1, e2, e3, e1, e2, e3, e3, e3}))

		tAssert.Error(t, a.Check([]cmpType{e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e1, e1}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2}))
		tAssert.Error(t, a.Check([]cmpType{e1, e2, e1, e1, e1}))

		tAssert.Equal(t, "e1", fnNewAssert().UniquesLenNotInRange(1, 2, "e1").Check([]cmpType{e1}).Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenNotInRange(1, 2).Check([]cmpType{e1}, "e2").Error())
		tAssert.Equal(t, "e2", fnNewAssert().UniquesLenNotInRange(1, 2, "e1").Check([]cmpType{e1}, "e2").Error())
	})
}

// #####################################################################################################################
