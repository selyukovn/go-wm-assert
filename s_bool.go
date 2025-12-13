package assert

type ABool struct {
	*assert[bool]
	*mixinComparable[*ABool, bool]
	*mixinCustom[*ABool, bool]
}

func Bool() *ABool {
	a := new(ABool)

	*a = ABool{
		assert:          newAssert[bool](),
		mixinComparable: newMixinComparable[*ABool, bool](a),
		mixinCustom:     newMixinCustom[*ABool, bool](a),
	}

	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// True
// ---------------------------------------------------------------------------------------------------------------------

// True -- alias to Eq(true)
func (a *ABool) True(customErrMsg ...string) *ABool {
	a.Eq(true, customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// False
// ---------------------------------------------------------------------------------------------------------------------

// False -- alias to Eq(false)
func (a *ABool) False(customErrMsg ...string) *ABool {
	a.Eq(false, customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
