package assert

import (
	tAssert "github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Any_unimplementableIsNil(t *testing.T) {
	fnIsNil := func(v any) bool {
		switch v.(type) {
		case nil:
			return true
		default:
			rve := reflect.ValueOf(v).Elem()
			if rve.IsValid() {
				return rve.IsNil()
			}
			return true
		}
	}

	var i interface{}
	tAssert.True(t, i == nil)
	tAssert.True(t, fnIsNil(i))

	var p *int
	tAssert.True(t, p == nil)
	tAssert.True(t, fnIsNil(p))

	i = p
	tAssert.False(t, i == nil)
	tAssert.True(t, fnIsNil(i)) // differs from Go's default behavior
}

func Test_Any_isZeroValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		tAssert.True(t, isZeroValue(nil))
	})

	t.Run("not nillable", func(t *testing.T) {
		// integers
		tAssert.True(t, isZeroValue(0))
		tAssert.True(t, isZeroValue(00))
		tAssert.False(t, isZeroValue(1))

		// floats
		tAssert.True(t, isZeroValue(0))
		tAssert.True(t, isZeroValue(0.0))
		tAssert.False(t, isZeroValue(0.1))
		tAssert.False(t, isZeroValue(1.0))

		// bool
		tAssert.True(t, isZeroValue(false))
		tAssert.False(t, isZeroValue(true))

		// string
		tAssert.True(t, isZeroValue(""))
		tAssert.False(t, isZeroValue("0"))
		tAssert.False(t, isZeroValue(" "))
		tAssert.False(t, isZeroValue("\n"))

		// struct
		tAssert.True(t, isZeroValue(struct{}{}))
		tAssert.True(t, isZeroValue(struct{ i int }{}))
		tAssert.False(t, isZeroValue(struct{ i int }{1}))

		// array
		tAssert.True(t, isZeroValue([0]int{}))
		tAssert.True(t, isZeroValue([1]int{}))
		tAssert.False(t, isZeroValue([1]int{1}))
	})

	t.Run("pointer", func(t *testing.T) {
		var p *int
		tAssert.True(t, p == nil)
		tAssert.True(t, isZeroValue(p))

		var pp **int
		tAssert.True(t, pp == nil)
		tAssert.True(t, isZeroValue(pp))
		pp = &p
		tAssert.False(t, pp == nil)
		tAssert.False(t, isZeroValue(pp))

		var ppp ***int
		tAssert.True(t, ppp == nil)
		tAssert.True(t, isZeroValue(ppp))
		ppp = &pp
		tAssert.False(t, ppp == nil)
		tAssert.False(t, isZeroValue(ppp))
	})

	t.Run("interface", func(t *testing.T) {
		var p *int

		var ip interface{}
		tAssert.True(t, ip == nil)
		tAssert.True(t, isZeroValue(ip))
		ip = p
		tAssert.False(t, ip == nil)
		tAssert.True(t, isZeroValue(ip)) // differs from Go's default behavior.
	})

	t.Run("slice", func(t *testing.T) {
		var sN []int
		tAssert.True(t, sN == nil)
		tAssert.True(t, isZeroValue(sN))

		s0 := make([]int, 0)
		tAssert.False(t, s0 == nil)
		tAssert.False(t, isZeroValue(s0))
	})

	t.Run("map", func(t *testing.T) {
		var mN map[int]int
		tAssert.True(t, mN == nil)
		tAssert.True(t, isZeroValue(mN))

		m0 := make(map[int]int)
		tAssert.False(t, m0 == nil)
		tAssert.False(t, isZeroValue(m0))
	})

	t.Run("chan", func(t *testing.T) {
		var cN chan int
		tAssert.True(t, cN == nil)
		tAssert.True(t, isZeroValue(cN))

		c0 := make(chan int)
		tAssert.False(t, c0 == nil)
		tAssert.False(t, isZeroValue(c0))
	})

	t.Run("func", func(t *testing.T) {
		var fN func()
		tAssert.True(t, fN == nil)
		tAssert.True(t, isZeroValue(fN))

		f0 := func() {}
		tAssert.False(t, f0 == nil)
		tAssert.False(t, isZeroValue(f0))
	})
}

func Test_Any_isNilInDepth(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		tAssert.True(t, isNilInDepth(nil))
	})

	t.Run("not nillable", func(t *testing.T) {
		// integers
		tAssert.False(t, isNilInDepth(0))

		// floats
		tAssert.False(t, isNilInDepth(0.0))

		// bool
		tAssert.False(t, isNilInDepth(false))

		// string
		tAssert.False(t, isNilInDepth(""))

		// struct
		tAssert.False(t, isNilInDepth(struct{}{}))
		tAssert.False(t, isNilInDepth(struct{ i int }{}))

		// array
		tAssert.False(t, isNilInDepth([0]int{}))
		tAssert.False(t, isNilInDepth([1]int{}))
	})

	t.Run("pointer", func(t *testing.T) {
		var p *int
		tAssert.True(t, p == nil)
		tAssert.True(t, isNilInDepth(p))

		var pp **int
		tAssert.True(t, pp == nil)
		tAssert.True(t, isNilInDepth(pp))
		pp = &p
		tAssert.False(t, pp == nil)
		tAssert.True(t, isNilInDepth(pp)) // deep nil -- (*int)(*int)(nil)

		var ppp ***int
		tAssert.True(t, ppp == nil)
		tAssert.True(t, isNilInDepth(ppp))
		ppp = &pp
		tAssert.False(t, ppp == nil)
		tAssert.True(t, isNilInDepth(ppp)) // deep nil -- (*int)(*int)(*int)(nil)
	})

	t.Run("interface", func(t *testing.T) {
		var p *int

		var ip interface{}
		tAssert.True(t, ip == nil)
		tAssert.True(t, isNilInDepth(ip))
		ip = p
		tAssert.False(t, ip == nil)
		tAssert.True(t, isNilInDepth(ip)) // deep nil -- interface|(*int)(nil)

		var pip *interface{}
		tAssert.True(t, pip == nil)
		tAssert.True(t, isNilInDepth(pip))
		pip = &ip
		tAssert.False(t, pip == nil)
		tAssert.True(t, isNilInDepth(pip)) // deep nil -- *interface|(*int)(nil)

		var ipip interface{}
		tAssert.True(t, ipip == nil)
		tAssert.True(t, isNilInDepth(ipip))
		ipip = pip
		tAssert.False(t, ipip == nil)
		tAssert.True(t, isNilInDepth(ipip)) // deep nil -- interface|(*interface|(*int)(nil))
	})

	t.Run("slice", func(t *testing.T) {
		var sN []int
		tAssert.True(t, sN == nil)
		tAssert.True(t, isNilInDepth(sN))

		s0 := make([]int, 0)
		tAssert.False(t, s0 == nil)
		tAssert.False(t, isNilInDepth(s0))
	})

	t.Run("map", func(t *testing.T) {
		var mN map[int]int
		tAssert.True(t, mN == nil)
		tAssert.True(t, isNilInDepth(mN))

		m0 := make(map[int]int)
		tAssert.False(t, m0 == nil)
		tAssert.False(t, isNilInDepth(m0))
	})

	t.Run("chan", func(t *testing.T) {
		var cN chan int
		tAssert.True(t, cN == nil)
		tAssert.True(t, isNilInDepth(cN))

		c0 := make(chan int)
		tAssert.False(t, c0 == nil)
		tAssert.False(t, isNilInDepth(c0))
	})

	t.Run("func", func(t *testing.T) {
		var fN func()
		tAssert.True(t, fN == nil)
		tAssert.True(t, isNilInDepth(fN))

		f0 := func() {}
		tAssert.False(t, f0 == nil)
		tAssert.False(t, isNilInDepth(f0))
	})
}
