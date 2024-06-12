package PointerFactory

func init() {
	wp.checksumSymbols = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
}

/* Односимвольная контрольная сумма по входной строке */
func (w *WrapperObj) CRC(str string) string {
	var sum byte
	maxSize := byte(len(w.checksumSymbols))

	for pos, _ := range str {
		sum ^= str[pos]
	}

	for sum >= maxSize {
		sum %= maxSize
	}

	return w.checksumSymbols[sum]
}
