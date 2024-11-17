package PointerFactory

////////////////////////////////////

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

//

func (obj *GlobalObj) CRC(val string) (rune, rune) {
	sum := uint32(1)
	b := uint32(obj.base)
	bDiv3 := b / 3

	for i := 0; i < len(val); i++ {
		sum += uint32(val[i])
	}

	return chars[sum&(b-1)], chars[sum&(bDiv3-1)]
}
