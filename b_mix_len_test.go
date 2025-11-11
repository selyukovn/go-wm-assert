package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_MixinLen(t *testing.T) {
	type testAssert struct {
		*assert[string]
		*mixinLen[*testAssert, string]
	}
	fnNewAssert := func() *testAssert {
		a := new(testAssert)

		*a = testAssert{
			assert:   newAssert[string](),
			mixinLen: newMixinLen[*testAssert, string](a),
		}

		return a
	}

	t.Run("LenEq", func(t *testing.T) {
		v := "hello"

		tAssert.NoError(t, fnNewAssert().LenEq(5).Check(v))

		tAssert.Error(t, fnNewAssert().LenEq(3).Check(v))

		tAssert.Equal(t, "e1", fnNewAssert().LenEq(3, "e1").Check(v).Error())
		tAssert.Equal(t, "e2", fnNewAssert().LenEq(3, "e1").Check(v, "e2").Error())
		tAssert.Equal(t, "e3", fnNewAssert().LenEq(3).Check(v, "e3").Error())
	})

	t.Run("LenNotEq", func(t *testing.T) {
		v := "world"

		tAssert.NoError(t, fnNewAssert().LenNotEq(3).Check(v))

		tAssert.Error(t, fnNewAssert().LenNotEq(5).Check(v))
		tAssert.Equal(t, "e1", fnNewAssert().LenNotEq(5, "e1").Check(v).Error())
		tAssert.Equal(t, "e2", fnNewAssert().LenNotEq(5, "e1").Check(v, "e2").Error())
		tAssert.Equal(t, "e3", fnNewAssert().LenNotEq(5).Check(v, "e3").Error())
	})

	t.Run("LenMin", func(t *testing.T) {
		for _, sucCase := range []string{"puppy", "elephant"} {
			tAssert.NoError(t, fnNewAssert().LenMin(5).Check(sucCase))
		}
		for _, errCase := range []string{"", "cat", "bird"} {
			tAssert.Error(t, fnNewAssert().LenMin(5).Check(errCase))
			tAssert.Equal(t, "e1", fnNewAssert().LenMin(5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", fnNewAssert().LenMin(5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", fnNewAssert().LenMin(5).Check(errCase, "e3").Error())
		}
	})

	t.Run("LenMax", func(t *testing.T) {
		for _, sucCase := range []string{"cat", "bird", "puppy"} {
			tAssert.NoError(t, fnNewAssert().LenMax(5).Check(sucCase))
		}
		for _, errCase := range []string{"puppies", "elephant"} {
			tAssert.Error(t, fnNewAssert().LenMax(5).Check(errCase))
			tAssert.Equal(t, "e1", fnNewAssert().LenMax(5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", fnNewAssert().LenMax(5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", fnNewAssert().LenMax(5).Check(errCase, "e3").Error())
		}
	})

	t.Run("LenInRange", func(t *testing.T) {
		for _, sucCase := range []string{"cat", "bird", "puppy"} {
			tAssert.NoError(t, fnNewAssert().LenInRange(3, 5).Check(sucCase))
		}
		for _, errCase := range []string{"", "assert", "elephant"} {
			tAssert.Error(t, fnNewAssert().LenInRange(3, 5).Check(errCase))
			tAssert.Equal(t, "e1", fnNewAssert().LenInRange(3, 5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", fnNewAssert().LenInRange(3, 5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", fnNewAssert().LenInRange(3, 5).Check(errCase, "e3").Error())
		}
	})

	t.Run("LenNotInRange", func(t *testing.T) {
		for _, sucCase := range []string{"", "assert", "elephant"} {
			tAssert.NoError(t, fnNewAssert().LenNotInRange(3, 5).Check(sucCase))
		}
		for _, errCase := range []string{"cat", "bird", "puppy"} {
			tAssert.Error(t, fnNewAssert().LenNotInRange(3, 5).Check(errCase))
			tAssert.Equal(t, "e1", fnNewAssert().LenNotInRange(3, 5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", fnNewAssert().LenNotInRange(3, 5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", fnNewAssert().LenNotInRange(3, 5).Check(errCase, "e3").Error())
		}
	})
}
