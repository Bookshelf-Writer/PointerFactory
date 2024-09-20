package PointerFactory

////////////////////////////////////

func NumToString(number uint64, base int32) string {
	var result string
	bb := uint64(base)

	for number > 0 {
		remainder := number % bb
		result = string(chars[remainder]) + result
		number /= bb
	}

	return result
}

func StringToNum(text string, base int32) uint64 {
	var result uint64
	bb := uint64(base)

	for _, c := range text {
		digit := uint64(charsMap[c])
		result = result*bb + digit
	}

	return result
}
