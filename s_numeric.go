package assert

type NumericTypes interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type ANumeric[T NumericTypes] struct {
	*assert[T]
	*mixinComparable[*ANumeric[T], T]
	*mixinCustom[*ANumeric[T], T]
	*mixinOrdered[*ANumeric[T], T]
}

func numericFnCmp[T NumericTypes](bigger, smaller T) bool {
	return bigger > smaller
}

func Numeric[T NumericTypes]() *ANumeric[T] {
	a := new(ANumeric[T])

	*a = ANumeric[T]{
		assert:          newAssert[T](),
		mixinComparable: newMixinComparable[*ANumeric[T], T](a),
		mixinCustom:     newMixinCustom[*ANumeric[T], T](a),
		mixinOrdered:    newMixinOrdered[*ANumeric[T], T](a, numericFnCmp[T]),
	}

	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Sign
// ---------------------------------------------------------------------------------------------------------------------

// Negative -- alias to Less(0)
func (a *ANumeric[T]) Negative(customErrMsg ...string) *ANumeric[T] {
	a.Less(0, customErrMsg...)
	return a
}

// Zero -- alias to Eq(0)
func (a *ANumeric[T]) Zero(customErrMsg ...string) *ANumeric[T] {
	a.Eq(0, customErrMsg...)
	return a
}

// NotZero -- alias to NotEq(0)
func (a *ANumeric[T]) NotZero(customErrMsg ...string) *ANumeric[T] {
	a.NotEq(0, customErrMsg...)
	return a
}

// Positive -- alias to Greater(0)
func (a *ANumeric[T]) Positive(customErrMsg ...string) *ANumeric[T] {
	a.Greater(0, customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
