package assert

type ASliceAny[S sliceType[E], E any] struct {
	*assert[S]
	*mixinCustom[*ASliceAny[S, E], S]
	*mixinSliceAny[*ASliceAny[S, E], S, E]
}

func SliceAny[S sliceType[E], E any]() *ASliceAny[S, E] {
	a := new(ASliceAny[S, E])

	*a = ASliceAny[S, E]{
		assert:        newAssert[S](),
		mixinCustom:   newMixinCustom[*ASliceAny[S, E], S](a),
		mixinSliceAny: newMixinSliceAny[*ASliceAny[S, E], S, E](a),
	}

	return a
}
