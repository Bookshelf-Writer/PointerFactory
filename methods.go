package PointerFactory

////////////////////////////////////

func (obj *GlobalObj) IsValid(uid string) error {
	if !obj.isActive {
		return ErrNotActive
	}

	size := len([]rune(uid))
	if size <= 5 {
		return ErrValidLength
	}

	group := rune(uid[0])
	_, ok := obj.groups[group]
	if !ok {
		return ErrValidGroup
	}

	c1, c2 := rune(uid[size-2:][0]), rune(uid[size-1:][0])
	r1, r2 := CRC(uid[:size-2], obj.base)
	if c1 != r1 || c2 != r2 {
		return ErrValidCRC
	}

	return nil
}

////

func (obj *GlobalObj) New(group rune) (string, error) {
	if !obj.isActive {
		return "", ErrNotActive
	}

	_, ok := obj.groups[group]
	if !ok {
		return "", ErrGroupNotFound
	}

	uid := obj.newUID(group, obj.sendChan(group))
	return uid.String(obj.base), nil
}
