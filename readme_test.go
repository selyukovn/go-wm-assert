package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Readme(t *testing.T) {
	t.Run("example_1", func(t *testing.T) {
		value := "Gopher"
		err := String().
			Word().
			LenMax(100, "Too long, isn’t it?" /* <-- custom error message only for this fail */).
			NotIn([]string{"some", "bad", "words"}).
			Custom(func(value string) error {
				// if ... { return err }
				return nil
			}).
			Check(value, "Name is incorrect" /* <-- custom error message that overrides any other in the chain */)

		tAssert.NoError(t, err)
	})

	t.Run("example_2", func(t *testing.T) {
		tAssert.Panics(t, func() {
			Time().NotZero().Must(time.Time{})
		})
		tAssert.NotPanics(t, func() {
			Time().NotZero().Must(time.Now())
		})

		type eventCollection struct{}
		tAssert.Panics(t, func() {
			var p *eventCollection
			Comparable[*eventCollection]().NotEq(nil).Must(p)
		})
		tAssert.NotPanics(t, func() {
			p := new(eventCollection)
			Comparable[*eventCollection]().NotEq(nil).Must(p)
		})
	})
}
