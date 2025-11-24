package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

// ---------------------------------------------------------------------------------------------------------------------
// Substrings
// ---------------------------------------------------------------------------------------------------------------------

func Test_AString_Substrings(t *testing.T) {
	// Prefix
	// --------------------------------

	t.Run("PrefixEq", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().PrefixEq("").Check(""))
		tAssert.NoError(t, Str().PrefixEq("").Check("any-string"))

		// not empty
		a := Str().PrefixEq("hel")

		tAssert.NoError(t, a.Check("hel"))
		tAssert.NoError(t, a.Check("hell"))
		tAssert.NoError(t, a.Check("hello"))
		tAssert.NoError(t, a.Check("hello, world"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("Hel"))
		tAssert.Error(t, a.Check("Hell"))
		tAssert.Error(t, a.Check("Hello, world"))
		tAssert.Error(t, a.Check("world"))

		tAssert.Equal(t, "e1", Str().PrefixEq("hel", "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().PrefixEq("hel").Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().PrefixEq("hel", "e1").Check("", "e2").Error())
	})

	t.Run("PrefixNotEq", func(t *testing.T) {
		// empty
		tAssert.Error(t, Str().PrefixNotEq("").Check(""))
		tAssert.Error(t, Str().PrefixNotEq("").Check("any-string"))

		// not empty
		a := Str().PrefixNotEq("hel")

		tAssert.NoError(t, a.Check(""))
		tAssert.NoError(t, a.Check("Hel"))
		tAssert.NoError(t, a.Check("Hell"))
		tAssert.NoError(t, a.Check("Hello, world"))
		tAssert.NoError(t, a.Check("world"))

		tAssert.Error(t, a.Check("hel"))
		tAssert.Error(t, a.Check("hell"))
		tAssert.Error(t, a.Check("hello"))
		tAssert.Error(t, a.Check("hello, world"))

		tAssert.Equal(t, "e1", Str().PrefixNotEq("hel", "e1").Check("hel").Error())
		tAssert.Equal(t, "e2", Str().PrefixNotEq("hel").Check("hel", "e2").Error())
		tAssert.Equal(t, "e2", Str().PrefixNotEq("hel", "e1").Check("hel", "e2").Error())
	})

	t.Run("PrefixIn", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().PrefixIn([]string{""}).Check(""))
		tAssert.NoError(t, Str().PrefixIn([]string{""}).Check("any-string"))
		tAssert.Error(t, Str().PrefixIn([]string{}).Check(""))
		tAssert.Error(t, Str().PrefixIn([]string{}).Check("any-string"))

		// not empty
		a := Str().PrefixIn([]string{"hel", "Hel"})

		tAssert.NoError(t, a.Check("hel"))
		tAssert.NoError(t, a.Check("hell"))
		tAssert.NoError(t, a.Check("hello"))
		tAssert.NoError(t, a.Check("hello, world"))
		tAssert.NoError(t, a.Check("Hel"))
		tAssert.NoError(t, a.Check("Hell"))
		tAssert.NoError(t, a.Check("Hello"))
		tAssert.NoError(t, a.Check("Hello, world"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("world"))

		tAssert.Equal(t, "e1", Str().PrefixIn([]string{"hel", "Hel"}, "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().PrefixIn([]string{"hel", "Hel"}).Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().PrefixIn([]string{"hel", "Hel"}, "e1").Check("", "e2").Error())
	})

	t.Run("PrefixNotIn", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().PrefixNotIn([]string{}).Check(""))
		tAssert.NoError(t, Str().PrefixNotIn([]string{}).Check("any-string"))
		tAssert.Error(t, Str().PrefixNotIn([]string{""}).Check(""))
		tAssert.Error(t, Str().PrefixNotIn([]string{""}).Check("any-string"))

		// not empty
		a := Str().PrefixNotIn([]string{"hel", "Hel"})

		tAssert.NoError(t, a.Check(""))
		tAssert.NoError(t, a.Check("world"))

		tAssert.Error(t, a.Check("hel"))
		tAssert.Error(t, a.Check("hell"))
		tAssert.Error(t, a.Check("hello"))
		tAssert.Error(t, a.Check("hello, world"))
		tAssert.Error(t, a.Check("Hel"))
		tAssert.Error(t, a.Check("Hell"))
		tAssert.Error(t, a.Check("Hello"))
		tAssert.Error(t, a.Check("Hello, world"))

		tAssert.Equal(t, "e1", Str().PrefixNotIn([]string{"hel", "Hel"}, "e1").Check("hel").Error())
		tAssert.Equal(t, "e2", Str().PrefixNotIn([]string{"hel", "Hel"}).Check("hel", "e2").Error())
		tAssert.Equal(t, "e2", Str().PrefixNotIn([]string{"hel", "Hel"}, "e1").Check("hel", "e2").Error())
	})

	// Suffix
	// --------------------------------

	t.Run("SuffixEq", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().SuffixEq("").Check(""))
		tAssert.NoError(t, Str().SuffixEq("").Check("any-string"))

		// not empty
		a := Str().SuffixEq("able")

		tAssert.NoError(t, a.Check("able"))
		tAssert.NoError(t, a.Check("table"))
		tAssert.NoError(t, a.Check("applicable"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("Able"))
		tAssert.Error(t, a.Check("tablet"))
		tAssert.Error(t, a.Check("applicable application"))

		tAssert.Equal(t, "e1", Str().SuffixEq("able", "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().SuffixEq("able").Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().SuffixEq("able", "e1").Check("", "e2").Error())
	})

	t.Run("SuffixNotEq", func(t *testing.T) {
		// empty
		tAssert.Error(t, Str().SuffixNotEq("").Check(""))
		tAssert.Error(t, Str().SuffixNotEq("").Check("any-string"))

		// not empty
		a := Str().SuffixNotEq("able")

		tAssert.NoError(t, a.Check(""))
		tAssert.NoError(t, a.Check("Able"))
		tAssert.NoError(t, a.Check("tablet"))
		tAssert.NoError(t, a.Check("applicable application"))

		tAssert.Error(t, a.Check("able"))
		tAssert.Error(t, a.Check("table"))
		tAssert.Error(t, a.Check("applicable"))

		tAssert.Equal(t, "e1", Str().SuffixNotEq("able", "e1").Check("able").Error())
		tAssert.Equal(t, "e2", Str().SuffixNotEq("able").Check("able", "e2").Error())
		tAssert.Equal(t, "e2", Str().SuffixNotEq("able", "e1").Check("able", "e2").Error())
	})

	t.Run("SuffixIn", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().SuffixIn([]string{""}).Check(""))
		tAssert.NoError(t, Str().SuffixIn([]string{""}).Check("any-string"))
		tAssert.Error(t, Str().SuffixIn([]string{}).Check(""))
		tAssert.Error(t, Str().SuffixIn([]string{}).Check("any-string"))

		// not empty
		a := Str().SuffixIn([]string{"able", "cable"})

		tAssert.NoError(t, a.Check("able"))
		tAssert.NoError(t, a.Check("table"))
		tAssert.NoError(t, a.Check("applicable"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("Able"))
		tAssert.Error(t, a.Check("tablet"))
		tAssert.Error(t, a.Check("applicable application"))

		tAssert.Equal(t, "e1", Str().SuffixIn([]string{"able", "cable"}, "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().SuffixIn([]string{"able", "cable"}).Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().SuffixIn([]string{"able", "cable"}, "e1").Check("", "e2").Error())
	})

	t.Run("SuffixNotIn", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().SuffixNotIn([]string{}).Check(""))
		tAssert.NoError(t, Str().SuffixNotIn([]string{}).Check("any-string"))
		tAssert.Error(t, Str().SuffixNotIn([]string{""}).Check(""))
		tAssert.Error(t, Str().SuffixNotIn([]string{""}).Check("any-string"))

		// not empty
		a := Str().SuffixNotIn([]string{"able", "cable"})

		tAssert.NoError(t, a.Check(""))
		tAssert.NoError(t, a.Check("Able"))
		tAssert.NoError(t, a.Check("tablet"))
		tAssert.NoError(t, a.Check("applicable application"))

		tAssert.Error(t, a.Check("able"))
		tAssert.Error(t, a.Check("table"))
		tAssert.Error(t, a.Check("applicable"))

		tAssert.Equal(t, "e1", Str().SuffixNotIn([]string{"able", "cable"}, "e1").Check("able").Error())
		tAssert.Equal(t, "e2", Str().SuffixNotIn([]string{"able", "cable"}).Check("able", "e2").Error())
		tAssert.Equal(t, "e2", Str().SuffixNotIn([]string{"able", "cable"}, "e1").Check("able", "e2").Error())
	})

	// ContainsStr
	// --------------------------------

	t.Run("ContainsStr", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().ContainsStr("").Check(""))
		tAssert.NoError(t, Str().ContainsStr("").Check("any-string"))

		// not empty
		a := Str().ContainsStr("able")

		tAssert.NoError(t, a.Check("able to"))
		tAssert.NoError(t, a.Check("cable"))
		tAssert.NoError(t, a.Check("tablet"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("Able"))
		tAssert.Error(t, a.Check("a-ble"))

		tAssert.Equal(t, "e1", Str().ContainsStr("able", "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().ContainsStr("able").Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().ContainsStr("able", "e1").Check("", "e2").Error())
	})

	t.Run("NotContainsStr", func(t *testing.T) {
		// empty
		tAssert.Error(t, Str().NotContainsStr("").Check(""))
		tAssert.Error(t, Str().NotContainsStr("").Check("any-string"))

		// not empty
		a := Str().NotContainsStr("able")

		tAssert.NoError(t, a.Check(""))
		tAssert.NoError(t, a.Check("Able"))
		tAssert.NoError(t, a.Check("a-ble"))

		tAssert.Error(t, a.Check("able to"))
		tAssert.Error(t, a.Check("cable"))
		tAssert.Error(t, a.Check("tablet"))

		tAssert.Equal(t, "e1", Str().NotContainsStr("able", "e1").Check("cable").Error())
		tAssert.Equal(t, "e2", Str().NotContainsStr("able").Check("cable", "e2").Error())
		tAssert.Equal(t, "e2", Str().NotContainsStr("able", "e1").Check("cable", "e2").Error())
	})

	t.Run("ContainsStrAny", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().ContainsStrAny([]string{}).Check(""))
		tAssert.NoError(t, Str().ContainsStrAny([]string{}).Check("any-string"))
		tAssert.NoError(t, Str().ContainsStrAny([]string{""}).Check(""))
		tAssert.NoError(t, Str().ContainsStrAny([]string{""}).Check("any-string"))

		// not empty
		a := Str().ContainsStrAny([]string{"able", "Able"})

		tAssert.NoError(t, a.Check(". Able to"))
		tAssert.NoError(t, a.Check("able to"))
		tAssert.NoError(t, a.Check("cable"))
		tAssert.NoError(t, a.Check("tablet"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("a-ble"))

		tAssert.Equal(t, "e1", Str().ContainsStrAny([]string{"able", "Able"}, "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().ContainsStrAny([]string{"able", "Able"}).Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().ContainsStrAny([]string{"able", "Able"}, "e1").Check("", "e2").Error())
	})

	t.Run("ContainsStrEach", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().ContainsStrEach([]string{}).Check(""))
		tAssert.NoError(t, Str().ContainsStrEach([]string{}).Check("any-string"))
		tAssert.NoError(t, Str().ContainsStrEach([]string{""}).Check(""))
		tAssert.NoError(t, Str().ContainsStrEach([]string{""}).Check("any-string"))

		// not empty
		a := Str().ContainsStrEach([]string{"able", "Able"})

		tAssert.NoError(t, a.Check("Able to be unable"))

		tAssert.Error(t, a.Check(""))
		tAssert.Error(t, a.Check("Able"))
		tAssert.Error(t, a.Check("Unable"))

		tAssert.Equal(t, "e1", Str().ContainsStrEach([]string{"able", "Able"}, "e1").Check("").Error())
		tAssert.Equal(t, "e2", Str().ContainsStrEach([]string{"able", "Able"}).Check("", "e2").Error())
		tAssert.Equal(t, "e2", Str().ContainsStrEach([]string{"able", "Able"}, "e1").Check("", "e2").Error())
	})

	t.Run("ContainsStrNone", func(t *testing.T) {
		// empty
		tAssert.NoError(t, Str().ContainsStrNone([]string{}).Check(""))
		tAssert.NoError(t, Str().ContainsStrNone([]string{}).Check("any-string"))
		tAssert.Error(t, Str().ContainsStrNone([]string{""}).Check(""))
		tAssert.Error(t, Str().ContainsStrNone([]string{""}).Check("any-string"))

		// not empty
		a := Str().ContainsStrNone([]string{"able", "Able"})

		tAssert.NoError(t, a.Check(""))
		tAssert.NoError(t, a.Check("a-ble"))

		tAssert.Error(t, a.Check(". Able to"))
		tAssert.Error(t, a.Check("able to"))
		tAssert.Error(t, a.Check("cable"))
		tAssert.Error(t, a.Check("tablet"))

		tAssert.Equal(t, "e1", Str().ContainsStrNone([]string{"able", "Able"}, "e1").Check("able").Error())
		tAssert.Equal(t, "e2", Str().ContainsStrNone([]string{"able", "Able"}).Check("able", "e2").Error())
		tAssert.Equal(t, "e2", Str().ContainsStrNone([]string{"able", "Able"}, "e1").Check("able", "e2").Error())
	})
}

// ---------------------------------------------------------------------------------------------------------------------
// Runes
// ---------------------------------------------------------------------------------------------------------------------

func Test_AString_Runes(t *testing.T) {
	t.Run("RunesEq", func(t *testing.T) {
		tAssert.Error(t, Str().RunesEq(-5).Check(""))

		for _, sucCase := range []string{"туман", "night"} {
			tAssert.NoError(t, Str().RunesEq(5).Check(sucCase))
		}
		for _, errCase := range []string{"", "蛇蝎美人", "детектив"} {
			tAssert.Error(t, Str().RunesEq(5).Check(errCase))

			tAssert.Equal(t, "e1", Str().RunesEq(5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", Str().RunesEq(5).Check(errCase, "e2").Error())
			tAssert.Equal(t, "e2", Str().RunesEq(5, "e1").Check(errCase, "e2").Error())
		}
	})
	t.Run("RunesNotEq", func(t *testing.T) {
		tAssert.NoError(t, Str().RunesNotEq(-5).Check(""))

		for _, sucCase := range []string{"", "蛇蝎美人", "детектив"} {
			tAssert.NoError(t, Str().RunesNotEq(5).Check(sucCase))
		}
		for _, errCase := range []string{"туман", "night"} {
			tAssert.Error(t, Str().RunesNotEq(5).Check(errCase))

			tAssert.Equal(t, "e1", Str().RunesNotEq(5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", Str().RunesNotEq(5).Check(errCase, "e2").Error())
			tAssert.Equal(t, "e2", Str().RunesNotEq(5, "e1").Check(errCase, "e2").Error())
		}
	})

	t.Run("RunesMin", func(t *testing.T) {
		tAssert.NoError(t, Str().RunesMin(-5).Check(""))

		for _, sucCase := range []string{"night", "детектив", "femme fatale"} {
			tAssert.NoError(t, Str().RunesMin(5).Check(sucCase))
		}
		for _, errCase := range []string{"", "夜晚", "jazz"} {
			tAssert.Error(t, Str().RunesMin(5).Check(errCase))

			tAssert.Equal(t, "e1", Str().RunesMin(5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", Str().RunesMin(5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", Str().RunesMin(5).Check(errCase, "e3").Error())
		}
	})

	t.Run("RunesMax", func(t *testing.T) {
		tAssert.Error(t, Str().RunesMax(-5).Check(""))

		for _, sucCase := range []string{"", "夜晚", "jazz", "туман"} {
			tAssert.NoError(t, Str().RunesMax(5).Check(sucCase))
		}
		for _, errCase := range []string{"детектив", "femme fatale"} {
			tAssert.Error(t, Str().RunesMax(5).Check(errCase))
			tAssert.Equal(t, "e1", Str().RunesMax(5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", Str().RunesMax(5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", Str().RunesMax(5).Check(errCase, "e3").Error())
		}
	})

	t.Run("RunesInRange", func(t *testing.T) {
		tAssert.NoError(t, Str().RunesInRange(-1, 1).Check(""))
		tAssert.Error(t, Str().RunesInRange(0, -1).Check(""))
		tAssert.Error(t, Str().RunesInRange(-2, -1).Check(""))
		tAssert.Error(t, Str().RunesInRange(-1, -2).Check(""))

		for _, sucCase := range []string{"fog", "蛇蝎美人", "туман"} {
			tAssert.NoError(t, Str().RunesInRange(3, 5).Check(sucCase))
			tAssert.Error(t, Str().RunesInRange(5, 3).Check(sucCase))
		}
		for _, errCase := range []string{"", "夜晚", "детектив", "femme fatale"} {
			tAssert.Error(t, Str().RunesInRange(5, 3).Check(errCase))
			tAssert.Error(t, Str().RunesInRange(3, 5).Check(errCase))
			tAssert.Equal(t, "e1", Str().RunesInRange(3, 5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", Str().RunesInRange(3, 5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", Str().RunesInRange(3, 5).Check(errCase, "e3").Error())
		}
	})

	t.Run("RunesNotInRange", func(t *testing.T) {
		tAssert.Error(t, Str().RunesNotInRange(-2, 1).Check(""))
		tAssert.NoError(t, Str().RunesNotInRange(1, -2).Check(""))
		tAssert.NoError(t, Str().RunesNotInRange(-1, -2).Check(""))
		tAssert.NoError(t, Str().RunesNotInRange(-2, -1).Check(""))

		for _, sucCase := range []string{"", "夜晚", "детектив", "femme fatale"} {
			tAssert.NoError(t, Str().RunesNotInRange(3, 5).Check(sucCase))
			tAssert.NoError(t, Str().RunesNotInRange(5, 3).Check(sucCase))
		}
		for _, errCase := range []string{"fog", "蛇蝎美人", "туман"} {
			tAssert.NoError(t, Str().RunesNotInRange(5, 3).Check(errCase))
			tAssert.Error(t, Str().RunesNotInRange(3, 5).Check(errCase))
			tAssert.Equal(t, "e1", Str().RunesNotInRange(3, 5, "e1").Check(errCase).Error())
			tAssert.Equal(t, "e2", Str().RunesNotInRange(3, 5, "e1").Check(errCase, "e2").Error())
			tAssert.Equal(t, "e3", Str().RunesNotInRange(3, 5).Check(errCase, "e3").Error())
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------
// Regexp
// ---------------------------------------------------------------------------------------------------------------------

func Test_AString_Regexp(t *testing.T) {
	t.Run("Regexp", func(t *testing.T) {
		var err error
		r := regexp.MustCompile(`^[a-z]+$`)

		tAssert.NoError(t, Str().Regexp(r).Check("hello"))

		v := "12345"
		err = Str().Regexp(r).Check(v)
		tAssert.Error(t, err)
		err = Str().Regexp(r, "e1").Check(v)
		tAssert.Equal(t, "e1", err.Error())
		err = Str().Regexp(r, "e1").Check(v, "e2")
		tAssert.Equal(t, "e2", err.Error())
		err = Str().Regexp(r).Check(v, "e3")
		tAssert.Equal(t, "e3", err.Error())
	})

	t.Run("Word", func(t *testing.T) {
		for _, sucCase := range []string{
			"Hello", "Go-World", "this", "is", "assert", "YEAH-great-assert-package-YEAH",
		} {
			tAssert.NoError(t, Str().Word().Check(sucCase))
		}
		for _, errCase := range []string{
			"",
			"0",
			"-invalid", "invalid-", "-invalid-", "in--valid",
			"_invalid", "invalid_", "in_va_lid", "in__valid",
			" invalid", "invalid ", "in va lid", "in  valid",
			"0invalid", "invalid0", "in0va0lid", "in00valid",
		} {
			var err error
			err = Str().Word().Check(errCase)
			tAssert.Error(t, err)
			err = Str().Word("e1").Check(errCase)
			tAssert.Equal(t, "e1", err.Error())
			err = Str().Word("e1").Check(errCase, "e2")
			tAssert.Equal(t, "e2", err.Error())
			err = Str().Word().Check(errCase, "e3")
			tAssert.Equal(t, "e3", err.Error())
		}
	})

	t.Run("Numeric", func(t *testing.T) {
		for _, sucCase := range []string{
			"-999", "-3.14", "-1",
			"-0", "0.0", "0.00", "00.00", "00.0", "0000000", "0",
			"1", "3.14", "999",
		} {
			tAssert.NoError(t, Str().Numeric().Check(sucCase))
		}
		for _, errCase := range []string{
			"",
			"-", "--", "--1", "-1-1", "1-1", "1-1-", "1-",
			"+0", "+1", "+3.14", "+999", // umm... no.
			" 0", "0 ", "1 000 000",
			"+", "++", "++1", "+1+1", "1+1", "1+1+", "1+",
			".", "..", ".1", "..1", ".1.1", "1.1.", "1.",
			"1,000,000", "1_000_000",
		} {
			var err error
			err = Str().Numeric().Check(errCase)
			tAssert.Error(t, err)
			err = Str().Numeric("e1").Check(errCase)
			tAssert.Equal(t, "e1", err.Error())
			err = Str().Numeric("e1").Check(errCase, "e2")
			tAssert.Equal(t, "e2", err.Error())
			err = Str().Numeric().Check(errCase, "e3")
			tAssert.Equal(t, "e3", err.Error())
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------
