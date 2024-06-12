package PointerFactory

import (
	"testing"
	"time"
)

func (w *WrapperObj) StartTests(t *testing.T) {
	wp.TypeMAP['#'] = 250
	wp.TypeMAP['$'] = 251
	wp.TypeMAP['&'] = 252
	wp.YearPoint = 2024
	wp.NumBase = 36

	for n, f := range testArr {
		if !t.Run(n, f) {
			break
		}
	}
}

var testArr = map[string]func(t *testing.T){
	"CRC":         testCRC,
	"Transform":   testTransform,
	"testCreator": testCreator,
}

//////////////////////////////////////////////////////////////////////////////

// testCRC Проверка контрольной суммы
func testCRC(t *testing.T) {
	buf := map[string]string{
		"c": "0oo18xvqp06",
		"r": "0oo18xw64lg",
		"x": "0oo18xvb9cf",
		"u": "gotest://PointerFactory#BenchmarkAdd",
	}

	for crc, value := range buf {
		if crc != wp.CRC(value) {
			t.Error("Not valid\t", crc, wp.CRC(value))
		}
	}
}

// testTransform проверка на трансформации
func testTransform(t *testing.T) {
	serverBuf := uint16(444)
	pointerBuf := uint64(99999999999)
	uidBuf := "#cc19xtf1tra"
	pointerBufObj := PointerObj{
		t: 250,
		s: serverBuf,
		p: pointerBuf,
	}

	//парсим
	parseObj, err := wp.Decode(uidBuf)
	if err != nil {
		t.Fatal(err)
	}

	//Проверям по строке
	if pointerBufObj.String() != uidBuf {
		t.Fatal("Fail Encode\t", pointerBufObj.String(), uidBuf)
	}

	if parseObj.t != pointerBufObj.t {
		t.Fatal("Invalid Type\t", parseObj.t, pointerBufObj.t)
	}
	if parseObj.s != pointerBufObj.s {
		t.Fatal("Invalid Server\t", parseObj.s, pointerBufObj.s)
	}
	if parseObj.p != pointerBufObj.p {
		t.Fatal("Invalid Pointer\t", parseObj.p, pointerBufObj.p)
	}
}

// testCreator Проверка генератора
func testCreator(t *testing.T) {
	dateTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	crtDateObj := wp.InitCreatorDate(444, dateTime)

	t.Run("Auto", func(t *testing.T) {
		dateArr := map[string]bool{
			"$cc2ms7i80m": true,
			"$cc2ms7i81n": true,
			"$cc2ms7i82k": true,
		}

		for uid, _ := range dateArr {
			id := crtDateObj.New(251)
			if uid != id.String() {
				t.Error("Not valid\t", uid, id.String())
			}
		}
	})

	t.Run("Fixed", func(t *testing.T) {
		dateArr := map[string]uint64{
			"&cc9ixe":     12345,
			"&ccqhjhp":    1235789,
			"&cc4jc8luu7": 9876543654,
		}

		for uid, pointer := range dateArr {
			id := crtDateObj.NewFixed(252, pointer)
			if uid != id.String() {
				t.Error("Not valid\t", uid, id.String())
			}
		}
	})
}
