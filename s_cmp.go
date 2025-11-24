package assert

type AComparable[T comparable] struct {
	*assert[T]
	*mixinComparable[*AComparable[T], T]
	*mixinCustom[*AComparable[T], T]
}

// Comparable
//
// Deprecated: use Cmp instead.
func Comparable[T comparable]() *AComparable[T] {
	a := new(AComparable[T])

	*a = AComparable[T]{
		assert:          newAssert[T](),
		mixinComparable: newMixinComparable[*AComparable[T], T](a),
		mixinCustom:     newMixinCustom[*AComparable[T], T](a),
	}

	return a
}

func Cmp[T comparable]() *AComparable[T] {
	return Comparable[T]()
}
