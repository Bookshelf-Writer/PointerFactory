package PointerFactory

import "bytes"

////////////////////////////////////

func (obj *GlobalObj) newUID(group rune, offset uint32) string {
	var buf bytes.Buffer

	buf.WriteRune(group)

	buf.WriteString(NumToString(uint64(obj.cluster), obj.base))
	buf.WriteString(NumToString(uint64(obj.minute), obj.base))
	buf.WriteString(NumToString(uint64(offset), obj.base))

	crc1, crc2 := CRC(buf.String(), obj.base)
	buf.WriteRune(crc1)
	buf.WriteRune(crc2)

	return buf.String()
}
