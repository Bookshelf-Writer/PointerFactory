package PointerFactory

////////////////////////////////////

func CRC(text string, base int32) (rune, rune) {
	sum := int32(1)

	for _, char := range text {
		sum += char
	}

	return sum % base, sum % (base / 2)
}
