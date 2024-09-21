package PointerFactory

////////////////////////////////////

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

//

func CRC(text string, base int32) (rune, rune) {
	sum := int32(1)

	for _, char := range text {
		sum += char
	}

	return chars[sum%base], chars[(sum % (base / 3))]
}
