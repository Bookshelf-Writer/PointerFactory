package PointerFactory

import (
	"strconv"
	"strings"
)

/* Обьект указателя формальный */
type PointerObj struct {
	t TypeTag // `Type`

	s uint16 // `Server`
	p uint64 // `Pointer`
}

//////////////////////////////////////////////////////////////////////////////

func (obj *PointerObj) String() string {
	var strBuf strings.Builder

	//получение типа
	strBuf.WriteString(obj.Type())

	//Формирование сервера
	strBuf.WriteString(obj.Server())

	//Формирование указателя
	strBuf.WriteString(obj.Pointer())

	//Добавление контрольной суммы
	return strBuf.String() + wp.CRC(strBuf.String())
}

func (obj *PointerObj) StringINT() string {
	return strconv.FormatUint(uint64(obj.s), 10) + strconv.FormatUint(obj.p, 10)
}

func (obj *PointerObj) Uint() uint64 {
	num, _ := strconv.ParseUint(obj.StringINT(), 10, 64)
	return num
}

//.//

func (obj *PointerObj) Compare(pointerObj *PointerObj) bool {

	if obj.t != pointerObj.t {
		return false
	}

	if obj.s != pointerObj.s {
		return false
	}

	if obj.p != pointerObj.p {
		return false
	}

	return true
}

//////////////////////////////////////////////////////////////////////////////

// uintToHEX Формирование строки в нужном счислении из десятичного числа
func (w *WrapperObj) uintToHEX(pointer uint64) string {
	return strconv.FormatUint(pointer, wp.NumBase)
}

// serverToString Формирование валидной строки указателя на сервер
func (w *WrapperObj) serverToString(serverPoint uint16) string {
	retStr := wp.uintToHEX(uint64(serverPoint))

	if len(retStr) == 1 {
		return "0" + retStr
	}

	return retStr
}

// pointerToString Получение символьного указателя
func (w *WrapperObj) pointerToString(pointerTag TypeTag) string {
	for char, typeTag := range w.TypeMAP {
		if pointerTag == typeTag {
			return string(char)
		}
	}
	return "0"
}

//////////////////////////////////////////////////////////////////////////////

func (obj *PointerObj) Type() string {
	return wp.pointerToString(obj.t)
}
func (obj *PointerObj) TypeN() byte {
	return byte(obj.t)
}

func (obj *PointerObj) Server() string {
	return wp.serverToString(obj.s)
}
func (obj *PointerObj) ServerN() uint16 {
	return obj.s
}

func (obj *PointerObj) Pointer() string {
	return wp.uintToHEX(obj.p)
}
func (obj *PointerObj) PointerN() uint64 {
	return obj.p
}

//////////////////////////////////////////////////////////////////////////////
