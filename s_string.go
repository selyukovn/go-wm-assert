package assert

import (
	"fmt"
	"regexp"
)

type AString struct {
	*assert[string]
	*mixinComparable[*AString, string]
	*mixinCustom[*AString, string]
	*mixinLen[*AString, string]
}

func String() *AString {
	a := new(AString)

	*a = AString{
		assert:          newAssert[string](),
		mixinComparable: newMixinComparable[*AString, string](a),
		mixinCustom:     newMixinCustom[*AString, string](a),
		mixinLen:        newMixinLen[*AString, string](a),
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
