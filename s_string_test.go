package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_AString_Regexp(t *testing.T) {
	t.Run("Regexp", func(t *testing.T) {
		var err error
		r := regexp.MustCompile(`^[a-z]+$`)

		tAssert.NoError(t, String().Regexp(r).Check("hello"))

		v := "12345"
		err = String().Regexp(r).Check(v)
		tAssert.Error(t, err)
		err = String().Regexp(r, "e1").Check(v)
		tAssert.Equal(t, "e1", err.Error())
		err = String().Regexp(r, "e1").Check(v, "e2")
		tAssert.Equal(t, "e2", err.Error())
		err = String().Regexp(r).Check(v, "e3")
		tAssert.Equal(t, "e3", err.Error())
	})

	t.Run("Word", func(t *testing.T) {
		for _, sucCase := range []string{
			"Hello", "Go-World", "this", "is", "assert", "YEAH-great-assert-package-YEAH",
		} {
			tAssert.NoError(t, String().Word().Check(sucCase))
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
			err = String().Word().Check(errCase)
			tAssert.Error(t, err)
			err = String().Word("e1").Check(errCase)
			tAssert.Equal(t, "e1", err.Error())
			err = String().Word("e1").Check(errCase, "e2")
			tAssert.Equal(t, "e2", err.Error())
			err = String().Word().Check(errCase, "e3")
			tAssert.Equal(t, "e3", err.Error())
		}
	})

	t.Run("Numeric", func(t *testing.T) {
		for _, sucCase := range []string{
			"-999", "-3.14", "-1",
			"-0", "0.0", "0.00", "00.00", "00.0", "0000000", "0",
			"1", "3.14", "999",
		} {
			tAssert.NoError(t, String().Numeric().Check(sucCase))
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
			err = String().Numeric().Check(errCase)
			tAssert.Error(t, err)
			err = String().Numeric("e1").Check(errCase)
			tAssert.Equal(t, "e1", err.Error())
			err = String().Numeric("e1").Check(errCase, "e2")
			tAssert.Equal(t, "e2", err.Error())
			err = String().Numeric().Check(errCase, "e3")
			tAssert.Equal(t, "e3", err.Error())
		}
	})
}
