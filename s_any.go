package assert

import (
	"fmt"
	"reflect"
)

type AAny[T any] struct {
	*assert[T]
	*mixinCustom[*AAny[T], T]
}

func Any[T any]() *AAny[T] {
	a := new(AAny[T])

	*a = AAny[T]{
		assert:      newAssert[T](),
		mixinCustom: newMixinCustom[*AAny[T], T](a),
	}

	return a
}

// Attention!
//
// It is currently technically impossible to correctly implement a non-recursive `NotNil` function
// that works similarly to Go's default `v != nil` behavior (e.g., passing the `interface|(*int)(nil)` check).
//
// The reason is in the `any` type of arguments.
// E.g. `func(v any)` handles `interface|(*int)(nil)` value same as `(*int)(nil)`.
// It would confuse a lot, since `interface|(*int)(nil)` should pass the `v != nil` check by Go's default behavior.
//
// See tests for detailed example.

// ---------------------------------------------------------------------------------------------------------------------
// Not Zero
// ---------------------------------------------------------------------------------------------------------------------

func isZeroValue(v any) bool {
	switch v.(type) {
	case nil:
		return true
	case
		int, int8, int16, int32 /* rune */, int64,
		uint, uint8 /* byte */, uint16, uint32, uint64, uintptr:
		return v == 0
	case float32, float64:
		return v == 0.0
	case complex64, complex128:
		return v == 0i
	case bool:
		return v == false
	case string:
		return v == ""
	default:
		return reflect.ValueOf(v).IsZero()
	}
}

// NotZero
//
// Fails check, if value is the zero value for its type.
//
// ATTENTION! THIS IS NOT THE SAME AS `v != nil` FOR INTERFACES: e.g. `interface|(*int)(nil)` fails the check.
func (a *AAny[T]) NotZero(customErrMsg ...string) *AAny[T] {
	a.addCheck(func(v T) error {
		if isZeroValue(v) {
			return mkCheckErr(
				fmt.Sprintf("value expects to be non-zero, got %s", fmtVal(v)),
				customErrMsg,
			)
		}

		return nil
	})
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Not Nil
// ---------------------------------------------------------------------------------------------------------------------

func isNilInDepth(v any) bool {
	switch v.(type) {
	case nil:
		return true
	case
		int, int8, int16, int32 /* rune */, int64,
		uint, uint8 /* byte */, uint16, uint32, uint64, uintptr,
		float32, float64,
		complex64, complex128,
		bool,
		string:
		return false
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Struct, reflect.Array:
		return false
	case reflect.Pointer, reflect.Interface:
		rve := rv.Elem()
		if rve.IsValid() {
			return isNilInDepth(rve.Interface())
		}
		return true
	default:
		return rv.IsNil()
	}
}

// NotNilDeep
//
// ATTENTION! THIS IS NOT THE SAME AS `v != nil`.
//
// For non-nil types, the check succeeds because values of these types are never `nil`.
//
// For nil types except pointers and interfaces, the check is analogous to Go's default behavior `v != nil`.
//
// For pointers and interfaces, the check runs recursively -- for example, for `interface|(*int)(nil)` it fails.
func (a *AAny[T]) NotNilDeep(customErrMsg ...string) *AAny[T] {
	a.addCheck(func(v T) error {
		if isNilInDepth(v) {
			return mkCheckErr(
				fmt.Sprintf("value expects to be non-nil in depth, got %s", fmtVal(v)),
				customErrMsg,
			)
		}

		return nil
	})

	return a
}

// ---------------------------------------------------------------------------------------------------------------------
