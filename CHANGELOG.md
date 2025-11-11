
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
    - [`Custom`](b_mix_custom.go) for custom checks -- added to all type-specific asserters
    - [`Ordered`](b_mix_ordered.go)
      - Less / LessEq / LessAny / LessEqAny / LessEach / LessEqEach
      - Greater / GreaterEq / GreaterAny / GreaterEqAny / GreaterEach / GreaterEqEach
      - InRange/NotInRange
- Type‑specific asserters: 
    - [`Bool`](s_bool.go) 
      - `Comparable` mixin
      - True / False aliases
    - [`Comparable[T comparable]`](s_comparable.go) -- `Comparable` mixin for `comparable` types like pointers, etc.
    - [`Numeric[T NumericTypes]`](s_numeric.go)
      - all mixins 
      - sign aliases: Negative / Zero / NotZero / Positive
    - [`String`](s_string.go) 
      - `Comparable` mixin
      - Empty / NotEmpty aliases
      - LenEq / LenNotEq / LenMin / LenMax / LenInRange / LenNotInRange
      - Regexp with Word and Numeric aliases
    - [`Time`](s_time.go)
      - all mixins
      - Zero / NotZero aliases
    - [`TimeDuration`](s_time_duration.go) 
      - all mixins
      - Zero / NotZero aliases
- Tests for core components and complex type‑specific asserters.
