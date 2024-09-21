package PointerFactory

import "errors"

////////////////////////////////////

var (
	ErrEmpyGroups          = errors.New("нет заданых групп")
	ErrInvalidStartPoint   = errors.New("startPoint указано для будущего")
	ErrInvalidGroupElement = errors.New("один из элементов группы не есть валидным")
)
