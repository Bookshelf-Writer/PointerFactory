package PointerFactory

import "errors"

////////////////////////////////////

var (
	ErrEmpyGroups          = errors.New("no groups defined")
	ErrInvalidBase         = errors.New("invalid radix size")
	ErrInvalidStartPoint   = errors.New("startPoint is specified for the future")
	ErrInvalidGroupElement = errors.New("one of the group elements is not valid")

	ErrNotActive     = errors.New("factory is not running")
	ErrInvalidChars  = errors.New("uid has invalid characters")
	ErrGroupNotFound = errors.New("no such group found")

	ErrValidLength = errors.New("invalid length")
	ErrValidGroup  = errors.New("invalid group")
	ErrValidCRC    = errors.New("invalid CRC")
)
