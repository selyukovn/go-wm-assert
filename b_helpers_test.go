package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_FmtVal(t *testing.T) {
	// string
	// --------------------------------

	t.Run("string -> quoted", func(t *testing.T) {
		tAssert.Equal(t, `"hello"`, fmtVal("hello"))
		tAssert.Equal(t, `""`, fmtVal(""))
	})

	// fmt.Stringer
	// --------------------------------

	t.Run("fmt.Stringer", func(t *testing.T) {
		tAssert.Equal(t, "2h0m0s", fmtVal(2*time.Hour))
	})

	// %v
	// --------------------------------

	t.Run("formatted as %v", func(t *testing.T) {
		// int
		tAssert.Equal(t, "42", fmtVal(42))

		// float
		tAssert.Equal(t, "3.14", fmtVal(3.14))

		// bool
		tAssert.Equal(t, "true", fmtVal(true))
		tAssert.Equal(t, "false", fmtVal(false))

		// slice
		tAssert.Equal(t, "[1 2 3]", fmtVal([]int{1, 2, 3}))

		// map
		tAssert.Contains(t, fmtVal(map[string]int{"a": 1, "b": 2}), "a:1")
		tAssert.Contains(t, fmtVal(map[string]int{"a": 1, "b": 2}), "b:2")

		// struct
		structSample := struct {
			A int
			B string
		}{A: 1, B: "test"}
		tAssert.Equal(t, "{1 test}", fmtVal(structSample))

		// pointer
		tAssert.NotPanics(t, func() { fmtVal(&structSample) })
		tAssert.NotNil(t, fmtVal(&structSample))
		tAssert.NotEmpty(t, fmtVal(&structSample))
		var structPointerNil *struct{}
		tAssert.NotPanics(t, func() { fmtVal(structPointerNil) })
		tAssert.NotNil(t, fmtVal(structPointerNil))
		tAssert.NotEmpty(t, fmtVal(structPointerNil))

		// nil-interface
		var nilInterface any
		tAssert.Equal(t, "<nil>", fmtVal(nilInterface))
	})
}
