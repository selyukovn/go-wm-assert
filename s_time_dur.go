package assert

import (
	"time"
)

type ATimeDuration struct {
	*assert[time.Duration]
	*mixinComparable[*ATimeDuration, time.Duration]
	*mixinCustom[*ATimeDuration, time.Duration]
	*mixinOrdered[*ATimeDuration, time.Duration]
}

func timeDurationFnCmp(bigger, smaller time.Duration) bool {
	return bigger > smaller
}

// TimeDuration
//
// Deprecated: use TimeDur() instead
func TimeDuration() *ATimeDuration {
	a := new(ATimeDuration)

	*a = ATimeDuration{
		assert:          newAssert[time.Duration](),
		mixinComparable: newMixinComparable[*ATimeDuration, time.Duration](a),
		mixinOrdered:    newMixinOrdered[*ATimeDuration, time.Duration](a, timeDurationFnCmp),
	}

	return a
}

func TimeDur() *ATimeDuration {
	return TimeDuration()
}

// ---------------------------------------------------------------------------------------------------------------------
// Zero
// ---------------------------------------------------------------------------------------------------------------------

// Zero -- alias to Eq(time.Duration(0))
func (a *ATimeDuration) Zero(customErrMsg ...string) *ATimeDuration {
	a.Eq(time.Duration(0), customErrMsg...)
	return a
}

// NotZero -- alias to NotEq(time.Duration(0))
func (a *ATimeDuration) NotZero(customErrMsg ...string) *ATimeDuration {
	a.NotEq(time.Duration(0), customErrMsg...)
	return a
}

// ---------------------------------------------------------------------------------------------------------------------
