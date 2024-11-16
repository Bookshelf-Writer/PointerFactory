package PointerFactory

import "regexp"

////////////////////////////////////

var ValidateRegExp = regexp.MustCompile("^[a-z0-9]+$")

func (obj *GlobalObj) IsValid(uid string) error {
	if !obj.isActive {
		return ErrNotActive
	}

	if !ValidateRegExp.MatchString(uid) {
		return ErrInvalidChars
	}

	buf := []rune(uid)
	size := len(buf)
	if size <= 5 {
		return ErrValidLength
	}

	_, ok := obj.groups[buf[0]]
	if !ok {
		return ErrValidGroup
	}

	c1, c2 := buf[size-2:][0], buf[size-1:][0]
	r1, r2 := CRC(string(buf[:size-2]), obj.base)
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

	offset, ok := obj.groups[group]
	if !ok {
		return "", ErrGroupNotFound
	}

	return obj.newUID(group, offset), nil
}
