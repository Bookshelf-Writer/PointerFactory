package PointerFactory

import (
	"math"
)

// //////////////////////////////////
const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func parseUint(s string, base int) uint64 {
	if s == "" {
		return 0
	}

	cutoff := math.MaxUint64/uint64(base) + 1

	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		case 'a' <= c && c <= 'z':
			d = c - 'a' + 10
		default:
			return 0
		}

		if d >= byte(base) {
			return 0
		}

		if n >= cutoff {
			return 0
		}
		n *= uint64(base)

		n1 := n + uint64(d)
		if n1 < n || n1 > math.MaxUint64 {
			return 0
		}
		n = n1
	}

	return n
}

func formatUint(u uint64, base int) (s string) {
	var a [64 + 1]byte
	i := len(a)

	b := uint64(base)
	for u >= b {
		i--
		q := u / b
		a[i] = digits[uint(u-q*b)]
		u = q
	}

	i--
	a[i] = digits[uint(u)]

	s = string(a[i:])
	return
}

////

func NumToString(number uint64) string {
	return formatUint(number, 36)
}

func StringToNum(text string) uint64 {
	return parseUint(text, 36)
}

//

func (obj *GlobalObj) NumToString(id uint64) string {
	return formatUint(id, obj.base)
}

func (obj *GlobalObj) StringToNum(id string) uint64 {
	return parseUint(id, obj.base)
}
