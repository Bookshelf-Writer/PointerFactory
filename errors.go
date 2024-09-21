package PointerFactory

import "errors"

////////////////////////////////////

var (
	ErrEmpyGroups          = errors.New("нет заданых групп")
	ErrInvalidBase         = errors.New("недопустимый размер базы")
	ErrInvalidStartPoint   = errors.New("startPoint указано для будущего")
	ErrInvalidGroupElement = errors.New("один из элементов группы не есть валидным")

	ErrNotActive     = errors.New("экземпляр не запущен")
	ErrGroupNotFound = errors.New("такая группа не найдена")

	ErrValidLength = errors.New("Недопустимая длинна")
	ErrValidGroup  = errors.New("недопустимая группа")
	ErrValidCRC    = errors.New("CRC не сошелся")
)
