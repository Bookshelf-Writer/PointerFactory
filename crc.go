package PointerFactory

////////////////////////////////////

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

//

func CRC(val string, base int32) (rune, rune) {
	sum := uint64(1)
	b := uint64(base)

	for _, char := range val {
		sum += uint64(char)
	}

	return chars[sum%b], chars[(sum % (b / 3))]
}
