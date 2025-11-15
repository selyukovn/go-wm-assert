package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_MixinComparable(t *testing.T) {
	type testAssert struct {
		*assert[int]
		*mixinComparable[*testAssert, int]
	}
	fnNewAssert := func() *testAssert {
		a := new(testAssert)
		*a = testAssert{
			assert:          newAssert[int](),
			mixinComparable: newMixinComparable[*testAssert, int](a),
		}
		return a
	}

	t.Run("Eq", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().Eq(42)
		tAssert.Error(t, a.Check(0))
		tAssert.NoError(t, a.Check(42))
		tAssert.Error(t, a.Check(11))

		a = fnNewAssert().Eq(11, "e11").Eq(22, "e22").Eq(33, "e33")
		tAssert.Equal(t, "e22", a.Check(11).Error())
		tAssert.Equal(t, "e11", a.Check(22).Error())
		tAssert.Equal(t, "e11", a.Check(33).Error())
		tAssert.Equal(t, "e44", a.Check(44, "e44").Error())
	})

	t.Run("NotEq", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().NotEq(42)
		tAssert.NoError(t, a.Check(0))
		tAssert.Error(t, a.Check(42))
		tAssert.NoError(t, a.Check(11))

		a = fnNewAssert().NotEq(11, "e11").NotEq(22, "e22").NotEq(33, "e33")
		tAssert.Equal(t, "e11", a.Check(11).Error())
		tAssert.Equal(t, "e22", a.Check(22).Error())
		tAssert.Equal(t, "e33", a.Check(33).Error())
		tAssert.NoError(t, a.Check(44, "e44"))
	})

	t.Run("In", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().In([]int{})
		tAssert.Error(t, a.Check(0))
		tAssert.Error(t, a.Check(11))
		tAssert.Error(t, a.Check(22))
		tAssert.Error(t, a.Check(33))
		tAssert.Error(t, a.Check(44))

		a = fnNewAssert().In([]int{11, 22, 33})
		tAssert.Error(t, a.Check(0))
		tAssert.NoError(t, a.Check(11))
		tAssert.NoError(t, a.Check(22))
		tAssert.NoError(t, a.Check(33))
		tAssert.Error(t, a.Check(44))

		a = fnNewAssert().In([]int{11, 22}, "e1122").In([]int{33}, "e33")
		tAssert.Equal(t, "e1122", a.Check(0).Error())
		tAssert.Equal(t, "e33", a.Check(11).Error())
		tAssert.Equal(t, "e33", a.Check(22).Error())
		tAssert.Equal(t, "e1122", a.Check(33).Error())
		tAssert.Equal(t, "e44", a.Check(44, "e44").Error())
	})

	t.Run("NotIn", func(t *testing.T) {
		var a *testAssert

		a = fnNewAssert().NotIn([]int{})
		tAssert.NoError(t, a.Check(0))
		tAssert.NoError(t, a.Check(11))
		tAssert.NoError(t, a.Check(22))
		tAssert.NoError(t, a.Check(33))
		tAssert.NoError(t, a.Check(44))

		a = fnNewAssert().NotIn([]int{11, 22, 33})
		tAssert.NoError(t, a.Check(0))
		tAssert.Error(t, a.Check(11))
		tAssert.Error(t, a.Check(22))
		tAssert.Error(t, a.Check(33))
		tAssert.NoError(t, a.Check(44))

		a = fnNewAssert().NotIn([]int{11, 22}, "e1122").NotIn([]int{33, 44}, "e3344")
		tAssert.NoError(t, a.Check(0))
		tAssert.Equal(t, "e1122", a.Check(11).Error())
		tAssert.Equal(t, "e1122", a.Check(22).Error())
		tAssert.Equal(t, "e3344", a.Check(33).Error())
		tAssert.Equal(t, "e44", a.Check(44, "e44").Error())
	})
}
