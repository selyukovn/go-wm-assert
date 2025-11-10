package assert

import (
	"fmt"
	"regexp"
)

type AString struct {
	*assert[string]
	*mixinComparable[*AString, string]
	*mixinCustom[*AString, string]
}

func String() *AString {
	a := new(AString)

	*a = AString{
		assert:          newAssert[string](),
		mixinComparable: newMixinComparable[*AString, string](a),
		mixinCustom:     newMixinCustom[*AString, string](a),
	}

	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Empty
// ---------------------------------------------------------------------------------------------------------------------

// Empty -- alias to Eq("")
func (a *AString) Empty(customErrMsg ...string) *AString {
	a.Eq("", customErrMsg...)
	return a
}

// NotEmpty -- alias to NotEq(time.Time{})
func (a *AString) NotEmpty(customErrMsg ...string) *AString {
	a.NotEq("", customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Len
// ---------------------------------------------------------------------------------------------------------------------

func (a *AString) lenComparable() *mixinComparable[*assert[int], int] {
	return newMixinComparable[*assert[int], int](newAssert[int]())
}

func stringFnCmpLen(bigger, smaller int) bool {
	return bigger > smaller
}

func (a *AString) lenOrdered() *mixinOrdered[*assert[int], int] {
	return newMixinOrdered[*assert[int], int](newAssert[int](), stringFnCmpLen)
}

func (a *AString) LenEq(eq int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := len(v)
		if nil != a.lenComparable().Eq(eq).Check(l) {
			return mkCheckErr(
				fmt.Sprintf(
					"length of %s expects to be equal to %s, got %s",
					fmtVal(v),
					fmtVal(eq),
					fmtVal(l),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

func (a *AString) LenNotEq(notEq int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := len(v)
		if nil != a.lenComparable().NotEq(notEq).Check(l) {
			return mkCheckErr(
				fmt.Sprintf(
					"length of %s expects to be not equal to %s, got %s",
					fmtVal(v),
					fmtVal(notEq),
					fmtVal(l),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

func (a *AString) LenMin(min int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := len(v)
		if nil != a.lenOrdered().GreaterEq(min).Check(l) {
			return mkCheckErr(
				fmt.Sprintf(
					"length of %s expects to be greater or equal than %s, got %s",
					fmtVal(v),
					fmtVal(min),
					fmtVal(l),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

func (a *AString) LenMax(max int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := len(v)
		if nil != a.lenOrdered().LessEq(max).Check(l) {
			return mkCheckErr(
				fmt.Sprintf(
					"length of %s expects to be less or equal than %s, got %s",
					fmtVal(v),
					fmtVal(max),
					fmtVal(l),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

func (a *AString) LenInRange(min, max int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := len(v)
		if nil != a.lenOrdered().InRange(min, max).Check(l) {
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
		}
		return nil
	})
	return a
}

func (a *AString) LenNotInRange(min, max int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := len(v)
		if nil != a.lenOrdered().NotInRange(min, max).Check(l) {
			return mkCheckErr(
				fmt.Sprintf(
					"length of %s expects not to be in range [%s, %s], got %s",
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
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Regexp
// ---------------------------------------------------------------------------------------------------------------------

func (a *AString) regexp(r *regexp.Regexp, customErrMsg []string) *AString {
	a.addCheck(func(v string) error {
		if !r.MatchString(v) {
			return mkCheckErr(
				fmt.Sprintf("value expects to be matched to regexp %s, got %s", fmtVal(r), fmtVal(v)),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

func (a *AString) Regexp(r *regexp.Regexp, customErrMsg ...string) *AString {
	return a.regexp(r, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
// Regexp - Word
// ---------------------------------------------------------------------------------------------------------------------

// StringRegexpWord
//
// Public variable to allow re-define globally.
var StringRegexpWord = regexp.MustCompile("^[A-Za-z](-?[A-Za-z]+)*$")

// Word
//
// See StringRegexpWord
func (a *AString) Word(customErrMsg ...string) *AString {
	return a.regexp(StringRegexpWord, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
// Regexp - Numeric
// ---------------------------------------------------------------------------------------------------------------------

// StringRegexpNumeric
//
// Public variable to allow re-define globally.
var StringRegexpNumeric = regexp.MustCompile("^-?\\d+(\\.\\d+)?$")

// Numeric
//
// See StringRegexpNumeric
func (a *AString) Numeric(customErrMsg ...string) *AString {
	return a.regexp(StringRegexpNumeric, customErrMsg)
}

// ---------------------------------------------------------------------------------------------------------------------
