package assert

import "fmt"

// #####################################################################################################################
// INTERFACE
// #####################################################################################################################

type assertInterface[T any] interface {
	// addCheck
	//
	// Registers custom validation check.
	//
	// Panics, if check is nil.
	addCheck(check func(v T) error)

	// Check
	//
	// Runs registered validation checks one by one against the given value and returns an error from assert first failed check.
	// Returns nil, if all checks pass.
	Check(v T, customErrMsg ...string) error

	// CheckAll
	//
	// Runs registered validation checks one by one against the given value and returns errors from all failed checks.
	// Returns empty slice, if all checks pass.
	CheckAll(v T) []error

	// Must
	//
	// Calls Check and panics with the error if validation fails.
	Must(v T, customErrMsg ...string)

	// MustAll
	//
	// Calls CheckAll and panics with the error slice if validation fails.
	MustAll(v T)
}

// #####################################################################################################################
// ASSERT
// #####################################################################################################################

type assert[T any] struct {
	checks []func(v T) error
}

func newAssert[T any]() *assert[T] {
	return &assert[T]{
		checks: make([]func(v T) error, 0, 1),
	}
}

// AddCheck
//
// Registers custom validation check.
//
// Panics, if check is nil.
func (a *assert[T]) addCheck(check func(v T) error) {
	if check == nil {
		panic(fmt.Errorf("%T.addCheck expects not nil check", a))
	}

	a.checks = append(a.checks, check)
}

// Check
//
// Runs registered validation checks one by one against the given value and returns an error from a first failed check.
// Returns nil, if all checks pass.
func (a *assert[T]) Check(v T, customErrMsg ...string) error {
	for _, check := range a.checks {
		if err := check(v); err != nil {
			if customErr := mkCustomErr(customErrMsg); customErr != nil {
				return customErr
			}
			return err
		}
	}
	return nil
}

// CheckAll
//
// Runs registered validation checks one by one against the given value and returns errors from all failed checks.
// Returns empty slice, if all checks pass.
func (a *assert[T]) CheckAll(v T) []error {
	errs := make([]error, 0, len(a.checks))
	for _, check := range a.checks {
		err := check(v)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// Must
//
// Calls Check and panics with the error if validation fails.
func (a *assert[T]) Must(v T, customErrMsg ...string) {
	err := a.Check(v, customErrMsg...)
	if err != nil {
		panic(err)
	}
}

// MustAll
//
// Calls CheckAll and panics with the error slice if validation fails.
func (a *assert[T]) MustAll(v T) {
	errs := a.CheckAll(v)
	if len(errs) > 0 {
		panic(errs)
	}
}

// #####################################################################################################################
