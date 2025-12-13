## [0.2.1] - 2025-11-15

*Note: v0.2.0 has been retracted.*

### BUGFIX

- Fixed `CRITICAL` issue on Go `1.18-1.20` with generic type inferring.

---

## [0.2.0] - 2025-11-15

### BUGFIX

- Fixed edge cases with empty slices and incorrect ranges (kind of empty sets) according to the Set theory:
    - `Comparable` mixin and related specific types:
        - In -- fails check, if no elements provided
        - NotIn -- passes check, if no elements provided
    - `Ordered` mixin:
        - LessAny / LessEqAny / LessEach / LessEqEach -- pass check, if no elements provided
        - GreaterAny / GreaterEqAny / GreaterEach / GreaterEqEach -- pass check, if no elements provided
        - InRange -- fails check, if min > max
        - NotInRange -- passes check, if min > max
    - `Len` mixin:
        - LenInRange -- fails check, if min > max
        - LenNotInRange -- passes check, if min > max

- Fixed default error messages

### FEATURES

- Added [`Len`](b_mix_len.go) mixin (from `String` specific Len...-methods)
    - LenEq / LenNotEq / LenMin / LenMax / LenInRange / LenNotInRange
    - [`String`](s_str.go) from now includes all methods from `Len` mixin

- Added [`SliceAny`](b_mix_slice.go) mixin
    - extends `Len` mixin
    - Empty / NotEmpty
    - CustomElementAny / CustomElementEach / CustomElementNone

- Added [`SliceCmp`](b_mix_slice.go) mixin
    - extends `SliceAny` mixin
    - Contains / NotContains / ContainsAny / ContainsEach / ContainsNone
    - Uniques
    - UniquesLenEq / UniquesLenNotEq / UniquesLenMin / UniquesLenMax / UniquesLenInRange / UniquesLenNotInRange

- Added [`SliceAny`](s_slice_any.go) specific type
    - `Custom` mixin
    - `SliceAny` mixin

- Added [`SliceCmp`](s_slice_cmp.go) specific type
    - `Custom` mixin
    - `SliceCmp` mixin

- Added new rules to [`String`](s_str.go) specific type
    - to check substrings
        - PrefixEq / PrefixNotEq / PrefixIn / PrefixNotIn
        - SuffixEq / SuffixNotEq / SuffixIn / SuffixNotIn
        - ContainsStr / NotContainsStr / ContainsStrAny / ContainsStrEach / ContainsStrNone
    - to check number of runes
        - RunesEq / RunesNotEq / RunesMin / RunesMax / RunesInRange / RunesNotInRange

- Improved [`README.md`](README.md)

---

## [0.1.0] - 2025-11-11

First implementation of the package.

### REQUIREMENTS

- Go 1.18+ (due to use of generics).

### FEATURES

- Fluent chaining (e.g., `assert.String().Word().LenMax(5).Check(value)`).

- `Must`/`MustAll` methods panic; `Check`/`CheckAll` return errors.

- Custom error messages for each chain part or for the whole chain to override any result error.

- Inner logic is based on mixins:
    - [`Comparable`](b_mix_comparable.go)
        - Eq / NotEq
        - In / NotIn

    - [`Custom`](b_mix_custom.go)
        - Custom

    - [`Ordered`](b_mix_ordered.go)
        - Less / LessEq / LessAny / LessEqAny / LessEach / LessEqEach
        - Greater / GreaterEq / GreaterAny / GreaterEqAny / GreaterEach / GreaterEqEach
        - InRange / NotInRange

- Type‑specific assertions:
    - [`Bool`](s_bool.go)
        - `Comparable` mixin
        - `Custom` mixin
        - True / False aliases

    - [`Comparable[T comparable]`](s_cmp.go)
        - `Comparable` mixin
        - `Custom` mixin

    - [`Numeric[T NumericTypes]`](s_num.go)
        - `Comparable` mixin
        - `Custom` mixin
        - `Ordered` mixin
        - sign aliases: Negative / Zero / NotZero / Positive

    - [`String`](s_str.go)
        - `Comparable` mixin
        - `Custom` mixin
        - Empty / NotEmpty aliases
        - LenEq / LenNotEq / LenMin / LenMax / LenInRange / LenNotInRange
        - Regexp with Word and Numeric aliases

    - [`Time`](s_time.go)
        - `Comparable` mixin
        - `Custom` mixin
        - `Ordered` mixin
        - Zero / NotZero aliases

    - [`TimeDuration`](s_time_dur.go)
        - `Comparable` mixin
        - `Custom` mixin
        - `Ordered` mixin
        - Zero / NotZero aliases

- Tests for
    - core components (basic assertion; mixins, except `Custom`; some helper functions)
    - complex type‑specific assertions (`String`)
    - examples from the Readme file
