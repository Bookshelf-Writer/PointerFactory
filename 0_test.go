package PointerFactory

import (
	"fmt"
	"strconv"
	"testing"
)

////////////////////////////////////

func TestName(t *testing.T) {
	num := uint64(1234534324)
	base := size

	str := NumToString(num, base)
	fmt.Println(str, strconv.FormatUint(num, 36))

	newNum := StringToNum(str, base)
	fmt.Println(newNum)
}
