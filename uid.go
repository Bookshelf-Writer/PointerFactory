package PointerFactory

import (
	"bytes"
	"sync/atomic"
)

////////////////////////////////////

func (obj *GlobalObj) newUID(group rune, offset *atomic.Uint32) string {
	var buf bytes.Buffer

	buf.WriteRune(group)

	buf.WriteString(NumToString(uint64(obj.cluster), obj.base))
	buf.WriteString(NumToString(uint64(obj.minute.Load()), obj.base))
	buf.WriteString(NumToString(uint64(offset.Load()), obj.base))

	crc1, crc2 := CRC(buf.String(), obj.base)
	buf.WriteRune(crc1)
	buf.WriteRune(crc2)

	offset.Add(1)
	return buf.String()
}
