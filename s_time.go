package assert

import (
	"time"
)

type ATime struct {
	*assert[time.Time]
	*mixinComparable[*ATime, time.Time]
	*mixinCustom[*ATime, time.Time]
	*mixinOrdered[*ATime, time.Time]
}

func timeFnCmp(bigger, smaller time.Time) bool {
	return smaller.Before(bigger)
}

func Time() *ATime {
	a := new(ATime)

	*a = ATime{
		assert:          newAssert[time.Time](),
		mixinComparable: newMixinComparable[*ATime, time.Time](a),
		mixinCustom:     newMixinCustom[*ATime, time.Time](a),
		mixinOrdered:    newMixinOrdered[*ATime, time.Time](a, timeFnCmp),
	}

	return a
}

// ---------------------------------------------------------------------------------------------------------------------
// Zero
// ---------------------------------------------------------------------------------------------------------------------

// Zero -- alias to Eq(time.Time{})
func (a *ATime) Zero(customErrMsg ...string) *ATime {
	a.Eq(time.Time{}, customErrMsg...)
	return a
}

// NotZero -- alias to NotEq(time.Time{})
func (a *ATime) NotZero(customErrMsg ...string) *ATime {
	a.NotEq(time.Time{}, customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
