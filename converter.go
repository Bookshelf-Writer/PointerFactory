package PointerFactory

import (
	"strconv"
)

////////////////////////////////////

func NumToString(number uint64, base int32) string {
	return strconv.FormatUint(number, int(base))
}

func StringToNum(text string, base int32) uint64 {
	number, _ := strconv.ParseUint(text, int(base), 64)
	return number
}
