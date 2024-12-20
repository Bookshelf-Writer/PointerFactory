package PointerFactory

import (
	"bytes"
	"sync/atomic"
)

////////////////////////////////////

func (obj *GlobalObj) newUID(group rune, offset *atomic.Uint32) string {
	var buf bytes.Buffer

	buf.WriteRune(group)

	buf.WriteString(obj.NumToString(uint64(obj.cluster)))
	buf.WriteString(obj.NumToString(uint64(obj.minute.Load())))
	buf.WriteString(obj.NumToString(uint64(offset.Load())))

	crc1, crc2 := obj.CRC(buf.String())
	buf.WriteRune(crc1)
	buf.WriteRune(crc2)

	offset.Add(1)
	return buf.String()
}
