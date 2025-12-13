package assert

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type AString struct {
	*assert[string]
	*mixinComparable[*AString, string]
	*mixinCustom[*AString, string]
	*mixinLen[*AString, string]
}

// String
//
// Deprecated: use Str() instead.
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

func Str() *AString {
	return String()
}

// ---------------------------------------------------------------------------------------------------------------------
// Empty
// ---------------------------------------------------------------------------------------------------------------------

// Empty -- alias to Eq("")
func (a *AString) Empty(customErrMsg ...string) *AString {
	a.Eq("", customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Not Empty
// ---------------------------------------------------------------------------------------------------------------------

// NotEmpty -- alias to NotEq(time.Time{})
func (a *AString) NotEmpty(customErrMsg ...string) *AString {
	a.NotEq("", customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Substrings - Prefix
// ---------------------------------------------------------------------------------------------------------------------

// Equal
// ---------------------------------------------------------------------------------------------------------------------

// PrefixEq
//
// Value expects to have prefix equal to "eq".
//
// Passes check, if empty prefix provided.
//
// See strings.HasPrefix.
func (a *AString) PrefixEq(eq string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if strings.HasPrefix(v, eq) {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to have prefix equal to %s, got %s",
				fmtVal(eq),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return a
}

// Not Equal
// ---------------------------------------------------------------------------------------------------------------------

// PrefixNotEq
//
// Value expects to have prefix not equal to "notEq".
//
// Fails check, if empty prefix provided.
//
// See strings.HasPrefix.
func (a *AString) PrefixNotEq(notEq string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if strings.HasPrefix(v, notEq) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to have prefix not equal to %s, got %s",
					fmtVal(notEq),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

// In
// ---------------------------------------------------------------------------------------------------------------------

// PrefixIn
//
// Value expects to have any of provided prefixes.
//
// Fails check, if no prefixes provided.
//
// See strings.HasPrefix.
func (a *AString) PrefixIn(in []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(in) > 0 {
			for _, p := range in {
				if strings.HasPrefix(v, p) {
					return nil
				}
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to have any of %s prefixes, got %s",
				fmtVal(in),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return a
}

// Not In
// ---------------------------------------------------------------------------------------------------------------------

// PrefixNotIn
//
// Value expects to have none of provided prefixes.
//
// Passes check, if no prefixes provided.
//
// See strings.HasPrefix.
func (a *AString) PrefixNotIn(notIn []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(notIn) == 0 {
			return nil
		}
		for _, p := range notIn {
			if strings.HasPrefix(v, p) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to have none of %s prefixes, got %s",
						fmtVal(notIn),
						fmtVal(v),
					),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Substrings - Suffix
// ---------------------------------------------------------------------------------------------------------------------

// Equal
// ---------------------------------------------------------------------------------------------------------------------

// SuffixEq
//
// Value expects to have suffix equal to "eq".
//
// Passes check, if empty suffix provided.
//
// See strings.HasSuffix.
func (a *AString) SuffixEq(eq string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if strings.HasSuffix(v, eq) {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to have suffix equal to %s, got %s",
				fmtVal(eq),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return a
}

// Not Eq
// ---------------------------------------------------------------------------------------------------------------------

// SuffixNotEq
//
// Value expects to have suffix not equal to "notEq".
//
// Fails check, if empty suffix provided.
//
// See strings.HasSuffix.
func (a *AString) SuffixNotEq(notEq string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if strings.HasSuffix(v, notEq) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to have suffix not equal to %s, got %s",
					fmtVal(notEq),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

// In
// ---------------------------------------------------------------------------------------------------------------------

// SuffixIn
//
// Value expects to have any of provided suffixes.
//
// Fails check, if no suffixes provided.
//
// See strings.HasSuffix.
func (a *AString) SuffixIn(in []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(in) > 0 {
			for _, s := range in {
				if strings.HasSuffix(v, s) {
					return nil
				}
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to have any of %s suffixes, got %s",
				fmtVal(in),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return a
}

// Not In
// ---------------------------------------------------------------------------------------------------------------------

// SuffixNotIn
//
// Value expects to have none of provided suffixes.
//
// Passes check, if no suffixes provided.
//
// See strings.HasSuffix.
func (a *AString) SuffixNotIn(notIn []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(notIn) == 0 {
			return nil
		}
		for _, s := range notIn {
			if strings.HasSuffix(v, s) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to have none of %s suffixes, got %s",
						fmtVal(notIn),
						fmtVal(v),
					),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Substrings - Contains
// ---------------------------------------------------------------------------------------------------------------------

// Contains
// ---------------------------------------------------------------------------------------------------------------------

// ContainsStr
//
// Value expects to contain provided substring.
//
// Passes check, if empty substring provided.
//
// See strings.Contains.
func (a *AString) ContainsStr(s string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if strings.Contains(v, s) {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to contain %s substring, got %s",
				fmtVal(s),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return a
}

// Not Contains
// ---------------------------------------------------------------------------------------------------------------------

// NotContainsStr
//
// Value expects to not contain provided substring.
//
// Fails check, if empty substring provided.
//
// See strings.Contains.
func (a *AString) NotContainsStr(s string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if strings.Contains(v, s) {
			return mkCheckErr(
				fmt.Sprintf(
					"value expects to not contain %s substring, got %s",
					fmtVal(s),
					fmtVal(v),
				),
				customErrMsg,
			)
		}
		return nil
	})
	return a
}

// Contains Any
// ---------------------------------------------------------------------------------------------------------------------

// ContainsStrAny
//
// Value expects to contain any of provided substrings.
//
// Passes check, if provided set of substrings is empty.
//
// See strings.Contains.
func (a *AString) ContainsStrAny(ss []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(ss) == 0 {
			return nil
		}
		for _, s := range ss {
			if strings.Contains(v, s) {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"value expects to contain any of %s substrings, got %s",
				fmtVal(ss),
				fmtVal(v),
			),
			customErrMsg,
		)
	})
	return a
}

// Contains Each
// ---------------------------------------------------------------------------------------------------------------------

// ContainsStrEach
//
// Value expects to contain each of provided substrings.
//
// Passes check, if provided set of substrings is empty.
//
// See strings.Contains.
func (a *AString) ContainsStrEach(ss []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(ss) == 0 {
			return nil
		}
		for _, s := range ss {
			if !strings.Contains(v, s) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to contain each of %s substrings, got %s",
						fmtVal(ss),
						fmtVal(v),
					),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return a
}

// Contains None
// ---------------------------------------------------------------------------------------------------------------------

// ContainsStrNone
//
// Value expects to contain none of provided substrings.
//
// Passes check, if provided set of substrings is empty.
//
// See strings.Contains.
func (a *AString) ContainsStrNone(ss []string, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if len(ss) == 0 {
			return nil
		}
		for _, s := range ss {
			if strings.Contains(v, s) {
				return mkCheckErr(
					fmt.Sprintf(
						"value expects to contain none of %s substrings, got %s",
						fmtVal(ss),
						fmtVal(v),
					),
					customErrMsg,
				)
			}
		}
		return nil
	})
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Runes
// ---------------------------------------------------------------------------------------------------------------------

// Equal
// ---------------------------------------------------------------------------------------------------------------------

// RunesEq
//
// Runes count of the value expects to be equal to "eq".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (a *AString) RunesEq(eq int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := utf8.RuneCountInString(v)
		if newMixinComparable[*assert[int], int](newAssert[int]()).Eq(eq).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"runes count of %s expects to be equal to %s, got %s",
				fmtVal(v),
				fmtVal(eq),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return a
}

// Not Equal
// ---------------------------------------------------------------------------------------------------------------------

// RunesNotEq
//
// Runes count of the value expects to be not equal to "notEq".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (a *AString) RunesNotEq(notEq int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := utf8.RuneCountInString(v)
		if newMixinComparable[*assert[int], int](newAssert[int]()).NotEq(notEq).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"runes count of %s expects to be not equal to %s, got %s",
				fmtVal(v),
				fmtVal(notEq),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return a
}

// ---------------------------------------------------------------------------------------------------------------------

func (a *AString) runesOrdCmp(b, s int) bool {
	return b > s
}

// Min
// ---------------------------------------------------------------------------------------------------------------------

// RunesMin
//
// Runes count of the value expects to be greater or equal to "min".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (a *AString) RunesMin(min int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := utf8.RuneCountInString(v)
		if newMixinOrdered[*assert[int], int](newAssert[int](), a.runesOrdCmp).GreaterEq(min).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"runes count of %s expects to be greater or equal to %s, got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return a
}

// Max
// ---------------------------------------------------------------------------------------------------------------------

// RunesMax
//
// Runes count of the value expects to be less or equal to "max".
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (a *AString) RunesMax(max int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := utf8.RuneCountInString(v)
		if newMixinOrdered[*assert[int], int](newAssert[int](), a.runesOrdCmp).LessEq(max).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"runes count of %s expects to be less or equal to %s, got %s",
				fmtVal(v),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return a
}

// In Range
// ---------------------------------------------------------------------------------------------------------------------

// RunesInRange
//
// Runes count of the value expects to be in range [min, max].
//
// Fails check, if min > max -- it works like empty range.
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (a *AString) RunesInRange(min, max int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		l := utf8.RuneCountInString(v)
		if min <= max {
			if newMixinOrdered[*assert[int], int](newAssert[int](), a.runesOrdCmp).InRange(min, max).Check(l) == nil {
				return nil
			}
		}
		return mkCheckErr(
			fmt.Sprintf(
				"runes count of %s expects to be in range [%s, %s], got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
	})
	return a
}

// Not In Range
// ---------------------------------------------------------------------------------------------------------------------

// RunesNotInRange
//
// Runes count of the value expects to be not in range [min, max] -- i.e. to be in ranges [0, min) or (max, MaxInt].
//
// Passes check, if min > max -- it works like empty range.
//
// Logically incorrect params (e.g. negative values, etc.) are processed as usual.
func (a *AString) RunesNotInRange(min, max int, customErrMsg ...string) *AString {
	a.addCheck(func(v string) error {
		if min > max {
			return nil
		}
		l := utf8.RuneCountInString(v)
		if newMixinOrdered[*assert[int], int](newAssert[int](), a.runesOrdCmp).NotInRange(min, max).Check(l) == nil {
			return nil
		}
		return mkCheckErr(
			fmt.Sprintf(
				"runes count of %s expects to be not in range [%s, %s], got %s",
				fmtVal(v),
				fmtVal(min),
				fmtVal(max),
				fmtVal(l),
			),
			customErrMsg,
		)
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
