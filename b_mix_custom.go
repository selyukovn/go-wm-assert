package assert

type mixinCustom[A assertInterface[T], T comparable] struct {
	assert A
}

func newMixinCustom[A assertInterface[T], T comparable](assert A) *mixinCustom[A, T] {
	return &mixinCustom[A, T]{assert: assert}
}

func (m *mixinCustom[A, T]) Custom(check func(v T) error) A {
	m.assert.addCheck(check)
	return m.assert
}
