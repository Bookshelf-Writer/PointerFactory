package PointerFactory

import "bytes"

////////////////////////////////////

type UidObj struct {
	Group   rune
	Cluster uint16
	Minute  uint32
	Offset  uint32
}

////

func (uid *UidObj) String(base int32) string {
	var buf bytes.Buffer

	buf.WriteRune(uid.Group)

	buf.WriteString(NumToString(uint64(uid.Cluster), base))
	buf.WriteString(NumToString(uint64(uid.Minute), base))
	buf.WriteString(NumToString(uint64(uid.Offset), base))

	crc1, crc2 := CRC(buf.String(), base)
	buf.WriteRune(crc1)
	buf.WriteRune(crc2)

	return buf.String()
}
