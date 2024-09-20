package PointerFactory

////////////////////////////////////

const template = "0123456789abcdefghijklmnopqrstuvwxyz"

var (
	chars    []rune
	charsMap map[rune]int
	size     int32
)

////

func init() {
	charsMap = make(map[rune]int)

	for pos, ch := range template {
		chars = append(chars, ch)
		charsMap[ch] = pos
	}

	size = int32(len(chars))
}
