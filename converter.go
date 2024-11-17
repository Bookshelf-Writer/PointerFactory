package PointerFactory

import "strconv"

////////////////////////////////////

func NumToString(number uint64) string {
	return strconv.FormatUint(number, 36)
}

func StringToNum(text string) uint64 {
	number, _ := strconv.ParseUint(text, 36, 64)
	return number
}

////

func (obj *GlobalObj) NumToString(id uint64) string {
	return strconv.FormatUint(id, obj.base)
}

func (obj *GlobalObj) StringToNum(id string) uint64 {
	number, _ := strconv.ParseUint(id, obj.base, 64)
	return number
}
