package assert

type mixinCustom[A assertInterface[T], T any] struct {
	assert A
}

func newMixinCustom[A assertInterface[T], T any](assert A) *mixinCustom[A, T] {
	return &mixinCustom[A, T]{assert: assert}
}

func (m *mixinCustom[A, T]) Custom(check func(v T) error) A {
	m.assert.addCheck(check)
	return m.assert
}
