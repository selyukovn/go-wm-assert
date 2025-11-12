package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_MixinOrdered(t *testing.T) {
	type testAssert struct {
		*assert[int]
		*mixinOrdered[*testAssert, int]
	}
	fnNewAssert := func() *testAssert {
		a := new(testAssert)
		*a = testAssert{
			assert: newAssert[int](),
			mixinOrdered: newMixinOrdered[*testAssert, int](a, func(bigger, smaller int) bool {
				return bigger > smaller
			}),
		}
		return a
	}

	// Less
	// --------------------------------

	t.Run("Less", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().Less(10)
		tAssert.NoError(t, a.Check(5))
		tAssert.Error(t, a.Check(10))
		tAssert.Error(t, a.Check(15))

		a = fnNewAssert().Less(10, "e1")
		tAssert.NoError(t, a.Check(5))
		tAssert.Equal(t, "e1", a.Check(10).Error())
		tAssert.Equal(t, "e2", a.Check(15, "e2").Error())
	})

	t.Run("LessEq", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().LessEq(10)
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.Error(t, a.Check(15))

		a = fnNewAssert().LessEq(10, "e1")
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.Equal(t, "e1", a.Check(15).Error())
		tAssert.Equal(t, "e2", a.Check(15, "e2").Error())
	})

	t.Run("LessAny", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().LessAny([]int{5, 10})
		tAssert.NoError(t, a.Check(3))
		tAssert.NoError(t, a.Check(8))
		tAssert.Error(t, a.Check(12))

		a = fnNewAssert().LessAny([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(3))
		tAssert.NoError(t, a.Check(8))
		tAssert.Equal(t, "e1", a.Check(12).Error())
		tAssert.Equal(t, "e2", a.Check(12, "e2").Error())
	})

	t.Run("LessEqAny", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().LessEqAny([]int{5, 10})
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(3))
		tAssert.NoError(t, a.Check(8))
		tAssert.Error(t, a.Check(15))

		a = fnNewAssert().LessEqAny([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(3))
		tAssert.NoError(t, a.Check(8))
		tAssert.Equal(t, "e1", a.Check(15).Error())
		tAssert.Equal(t, "e2", a.Check(15, "e2").Error())
	})

	t.Run("LessEach", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().LessEach([]int{5, 10})
		tAssert.NoError(t, a.Check(4))
		tAssert.Error(t, a.Check(6))
		tAssert.Error(t, a.Check(12))

		a = fnNewAssert().LessEach([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(4))
		tAssert.Equal(t, "e1", a.Check(6).Error())
		tAssert.Equal(t, "e2", a.Check(12, "e2").Error())
	})

	t.Run("LessEqEach", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().LessEqEach([]int{5, 10})
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(3))
		tAssert.Error(t, a.Check(7))
		tAssert.Error(t, a.Check(12))

		a = fnNewAssert().LessEqEach([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(3))
		tAssert.Equal(t, "e1", a.Check(7).Error())
		tAssert.Equal(t, "e2", a.Check(12, "e2").Error())
	})

	// Greater
	// --------------------------------

	t.Run("Greater", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().Greater(10)
		tAssert.Error(t, a.Check(5))
		tAssert.Error(t, a.Check(10))
		tAssert.NoError(t, a.Check(15))

		a = fnNewAssert().Greater(10, "e1")
		tAssert.Equal(t, "e1", a.Check(5).Error())
		tAssert.Equal(t, "e2", a.Check(10, "e2").Error())
		tAssert.NoError(t, a.Check(15))
	})

	t.Run("GreaterEq", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().GreaterEq(10)
		tAssert.Error(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(15))

		a = fnNewAssert().GreaterEq(10, "e1")
		tAssert.Equal(t, "e1", a.Check(5).Error())
		tAssert.Equal(t, "e2", a.Check(5, "e2").Error())
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(15))
	})

	t.Run("GreaterAny", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().GreaterAny([]int{5, 10})
		tAssert.NoError(t, a.Check(7))
		tAssert.NoError(t, a.Check(12))
		tAssert.Error(t, a.Check(3))

		a = fnNewAssert().GreaterAny([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(7))
		tAssert.NoError(t, a.Check(12))
		tAssert.Equal(t, "e1", a.Check(3).Error())
		tAssert.Equal(t, "e2", a.Check(3, "e2").Error())
	})

	t.Run("GreaterEqAny", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().GreaterEqAny([]int{5, 10})
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(7))
		tAssert.NoError(t, a.Check(12))
		tAssert.Error(t, a.Check(3))

		a = fnNewAssert().GreaterEqAny([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(7))
		tAssert.NoError(t, a.Check(12))
		tAssert.Equal(t, "e1", a.Check(3).Error())
		tAssert.Equal(t, "e2", a.Check(3, "e2").Error())
	})

	t.Run("GreaterEach", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().GreaterEach([]int{5, 10})
		tAssert.NoError(t, a.Check(12))
		tAssert.Error(t, a.Check(8))
		tAssert.Error(t, a.Check(3))

		a = fnNewAssert().GreaterEach([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(12))
		tAssert.Equal(t, "e1", a.Check(8).Error())
		tAssert.Equal(t, "e2", a.Check(3, "e2").Error())
	})

	t.Run("GreaterEqEach", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().GreaterEqEach([]int{5, 10})
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(12))
		tAssert.Error(t, a.Check(8))
		tAssert.Error(t, a.Check(5))
		tAssert.Error(t, a.Check(3))

		a = fnNewAssert().GreaterEqEach([]int{5, 10}, "e1")
		tAssert.NoError(t, a.Check(10))
		tAssert.NoError(t, a.Check(12))
		tAssert.Equal(t, "e1", a.Check(8).Error())
		tAssert.Equal(t, "e1", a.Check(5).Error())
		tAssert.Equal(t, "e2", a.Check(3, "e2").Error())
	})

	// In Range
	// --------------------------------

	t.Run("InRange", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().InRange(5, 10)
		tAssert.Error(t, a.Check(4))
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(7))
		tAssert.NoError(t, a.Check(10))
		tAssert.Error(t, a.Check(11))

		a = fnNewAssert().InRange(5, 10, "e1")
		tAssert.Equal(t, "e1", a.Check(4).Error())
		tAssert.NoError(t, a.Check(5))
		tAssert.NoError(t, a.Check(7))
		tAssert.NoError(t, a.Check(10))
		tAssert.Equal(t, "e2", a.Check(11, "e2").Error())
	})

	t.Run("NotInRange", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().NotInRange(5, 10)
		tAssert.NoError(t, a.Check(4))
		tAssert.Error(t, a.Check(5))
		tAssert.Error(t, a.Check(7))
		tAssert.Error(t, a.Check(10))
		tAssert.NoError(t, a.Check(11))

		a = fnNewAssert().NotInRange(5, 10, "e1")
		tAssert.NoError(t, a.Check(4))
		tAssert.Equal(t, "e1", a.Check(5).Error())
		tAssert.Equal(t, "e1", a.Check(7).Error())
		tAssert.Equal(t, "e2", a.Check(10, "e2").Error())
		tAssert.NoError(t, a.Check(11))
	})
}
