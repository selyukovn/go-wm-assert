package assert

import (
	"errors"
	"fmt"
)

func ternary[T any](condition bool, rTrue T, rFalse T) T {
	if condition {
		return rTrue
	}
	return rFalse
}

func fmtVal[T any](v T) string {
	switch tv := any(v).(type) {
	case string:
		// raw strings better to show in quotes
		return fmt.Sprintf("%q", tv)
	case fmt.Stringer:
		// e.g. time.Duration.String() provides better format -- better than %v
		return fmt.Sprintf("%s", tv)
	default:
		return fmt.Sprintf("%v", tv)
	}
}

func mkCustomErr(customErrMsg []string) error {
	if len(customErrMsg) > 0 && customErrMsg[0] != "" {
		return errors.New(customErrMsg[0])
	}
	return nil
}

func mkCheckErr(defaultMsg string, customErrMsg []string) error {
	if err := mkCustomErr(customErrMsg); err != nil {
		return err
	}

	return errors.New(defaultMsg)
}
