package PointerFactory

import (
	"bytes"
	"strconv"
	"time"
)

////////////////////////////////////

func uintToString(num uint64) string {
	return strconv.FormatUint(num, 36)
}

//

type GroupType byte

func (t *GroupType) String() string {
	return uintToString(uint64(*t))
}

//

type ClusterType uint32

func (t *ClusterType) String() string {
	return uintToString(uint64(*t))
}

//

type OffsetType uint16

func (t *OffsetType) String() string {
	return uintToString(uint64(*t))
}

//

type TimeType time.Time

func (t *TimeType) String() string {
	return uintToString(uint64(time.Time(*t).Minute()))
}

////

type GlobalObj struct {
	Group   GroupType
	Cluster ClusterType
	Time    TimeType
	Offset  OffsetType
}

func (uid *GlobalObj) String() string {
	var buf bytes.Buffer

	buf.WriteString(uid.Group.String())
	buf.WriteString(uid.Cluster.String())
	buf.WriteString(uid.Time.String())
	buf.WriteString(uid.Offset.String())
	buf.WriteString(crc(buf.String()))

	return buf.String()
}

//

func crc(uid string) string {
	sum := int32(1)

	for _, char := range uid {
		sum += char
	}

	return uintToString(uint64(sum%36)) + uintToString(uint64(sum%8))
}

func validate(uid string) bool {
	length := len(uid)
	ccc := uid[length-2:]
	uid = uid[:length-2]

	return crc(uid) == ccc
}

////

type FinalObj struct {
	Group   GroupType
	Cluster uint16
	Minute  uint32
	Offset  uint16
}
