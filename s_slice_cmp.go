package assert

type ASliceCmp[S sliceType[E], E comparable] struct {
	*assert[S]
	*mixinCustom[*ASliceCmp[S, E], S]
	*mixinSliceCmp[*ASliceCmp[S, E], S, E]
}

func SliceCmp[S sliceType[E], E comparable]() *ASliceCmp[S, E] {
	a := new(ASliceCmp[S, E])

	*a = ASliceCmp[S, E]{
		assert:        newAssert[S](),
		mixinCustom:   newMixinCustom[*ASliceCmp[S, E], S](a),
		mixinSliceCmp: newMixinSliceCmp[*ASliceCmp[S, E], S](a),
	}

	return a
}
