package assert

import (
	"errors"
	tAssert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_Assert(t *testing.T) {
	// Success
	// --------------------------------

	t.Run("no checks -> no error", func(t *testing.T) {
		a := newAssert[int]()
		tAssert.NoError(t, a.Check(42))
		tAssert.Len(t, a.CheckAll(42), 0)
		tAssert.NotPanics(t, func() { a.Must(42) })
		tAssert.NotPanics(t, func() { a.MustAll(42) })
	})

	t.Run("single passing check -> no error", func(t *testing.T) {
		a := newAssert[int]()
		a.addCheck(func(v int) error { return nil })
		tAssert.NoError(t, a.Check(42))
		tAssert.Len(t, a.CheckAll(42), 0)
		tAssert.NotPanics(t, func() { a.Must(42) })
		tAssert.NotPanics(t, func() { a.MustAll(42) })
	})

	t.Run("multiple checks, all pass -> no error", func(t *testing.T) {
		a := newAssert[string]()
		a.addCheck(func(v string) error { return nil })
		a.addCheck(func(v string) error { return nil })
		a.addCheck(func(v string) error { return nil })
		tAssert.NoError(t, a.Check("hello"))
		tAssert.Len(t, a.CheckAll("hello"), 0)
		tAssert.NotPanics(t, func() { a.Must("hello") })
		tAssert.NotPanics(t, func() { a.MustAll("hello") })
	})

	// Error
	// --------------------------------

	tErrMustAll := func(fnT func(), fnA func(errs []error)) {
		defer func() { fnA(recover().([]error)) }()
		fnT()
	}

	t.Run("single failing check -> returns error", func(t *testing.T) {
		rErr := errors.New("result error")
		a := newAssert[int]()
		a.addCheck(func(v int) error { return rErr })
		// check
		err := a.Check(42)
		tAssert.Error(t, err)
		tAssert.Equal(t, rErr, err)
		// check all
		errs := a.CheckAll(42)
		tAssert.Len(t, errs, 1)
		tAssert.Equal(t, rErr, errs[0])
		// must
		tAssert.PanicsWithValue(t, rErr, func() { a.Must(42) })
		// must all
		tErrMustAll(func() { a.MustAll(42) }, func(errs []error) {
			tAssert.Len(t, errs, 1)
			tAssert.Equal(t, rErr, errs[0])
		})
	})

	t.Run("multiple checks, several fail -> first error", func(t *testing.T) {
		rErr1 := errors.New("result error 1")
		rErr2 := errors.New("result error 2")
		rErr3 := errors.New("result error 3")
		a := newAssert[bool]()
		a.addCheck(func(v bool) error { return nil })
		a.addCheck(func(v bool) error { return nil })
		a.addCheck(func(v bool) error { return rErr1 })
		a.addCheck(func(v bool) error { return rErr2 })
		a.addCheck(func(v bool) error { return nil })
		a.addCheck(func(v bool) error { return rErr3 })
		// check
		err := a.Check(true)
		tAssert.Error(t, err)
		tAssert.Equal(t, rErr1, err)
		// check all
		errs := a.CheckAll(true)
		tAssert.Len(t, errs, 3)
		tAssert.Equal(t, rErr1, errs[0])
		tAssert.Equal(t, rErr2, errs[1])
		tAssert.Equal(t, rErr3, errs[2])
		// must
		tAssert.PanicsWithValue(t, rErr1, func() { a.Must(true) })
		// must all
		tErrMustAll(func() { a.MustAll(true) }, func(errs []error) {
			tAssert.Len(t, errs, 3)
			tAssert.Equal(t, rErr1, errs[0])
			tAssert.Equal(t, rErr2, errs[1])
			tAssert.Equal(t, rErr3, errs[2])
		})
	})

	// Custom message
	// --------------------------------

	t.Run("custom message", func(t *testing.T) {
		a := newAssert[string]()
		a.addCheck(func(v string) error { return nil })
		a.addCheck(func(v string) error { return errors.New("not this message") })
		a.addCheck(func(v string) error { return nil })
		err := a.Check("hello", "world1")
		tAssert.Error(t, err)
		tAssert.Equal(t, "world1", err.Error())
		tAssert.PanicsWithError(t, "world2", func() { a.Must("hello", "world2") })
	})
}
