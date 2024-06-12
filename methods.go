package PointerFactory

import (
	"errors"
	"strconv"
)

/* Получение указателя из строки */
func (w *WrapperObj) Decode(str string) (*PointerObj, error) {
	retObj := PointerObj{}
	bufLen := len(str)
	var err error

	//проверка на максимальную длину
	if bufLen > 16 {
		return &retObj, errors.New("Too long")
	}

	//проверка на минимальную длину
	if bufLen < 5 {
		return &retObj, errors.New("Too short")
	}

	// Получение типа
	retObj.t = wp.ParseType([]rune(str[0:1])[0])
	if retObj.t == TypeNone {
		return &retObj, errors.New("Unknown type [" + str[0:1] + "]")
	}

	//проверка контрольной суммы
	bufCRC := str[bufLen-1:]
	if bufCRC != wp.CRC(str[:bufLen-1]) {
		return &retObj, errors.New("Checksum verification error")
	}

	//Получение сервера
	bufServer, err := strconv.ParseUint(str[1:3], wp.NumBase, 16)
	if err != nil {
		return &retObj, err
	}
	retObj.s = uint16(bufServer)

	//получение указателя
	retObj.p, err = strconv.ParseUint(str[3:bufLen-1], wp.NumBase, 64)
	if err != nil {
		return &retObj, err
	}

	return &retObj, nil
}

/* Формирование строки из указателя */
func (w *WrapperObj) Encode(point *PointerObj) string {
	return point.String()
}

/* Формирование строки из указателя десятичным числом (без CRC) */
func (w *WrapperObj) EncodeINT(point *PointerObj) string {
	return point.StringINT()
}
