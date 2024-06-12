package PointerFactory

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func (w *WrapperObj) StartBenchmarks(b *testing.B) {
	for n, f := range benchmarkArr {
		if !b.Run(n, f) {
			break
		}
	}
}

var benchmarkArr = map[string]func(b *testing.B){
	"CRC": benchmarkChecksum,

	"Linear generation":                         benchmarkAdd,
	"Asynchronous generation":                   benchmarkAddPG,
	"Asynchronous generation (within a thread)": benchmarkAddP,
}

//////////////////////////////////////////////////////////////////////////////

// benchmarkChecksum Контрольная сумма
func benchmarkChecksum(b *testing.B) {
	str := strconv.Itoa(time.Now().Nanosecond())

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wp.CRC(str)
	}
	b.StopTimer()
}

// benchmarkAdd потокое созадание указателей
func benchmarkAdd(b *testing.B) {
	crt := wp.InitCreator(999)
	defer crt.Close()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range wp.TypeMAP {
			crt.New(tag)
		}
	}
	b.StopTimer()
}

// benchmarkAddP Асинхронное создание указателей в рамках потока
func benchmarkAddP(b *testing.B) {
	crt := wp.InitCreator(999)
	defer crt.Close()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(len(wp.TypeMAP))

		for _, tag := range wp.TypeMAP {
			bufTag := tag

			go func() {
				crt.New(bufTag)
				wg.Done()
			}()

		}

		wg.Wait()
	}
	b.StopTimer()
}

// benchmarkAddPG Асинхронное создание указателей за пределами потока
func benchmarkAddPG(b *testing.B) {
	crt := wp.InitCreator(999)
	defer crt.Close()

	var wg sync.WaitGroup
	wg.Add(len(wp.TypeMAP) * b.N)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range wp.TypeMAP {
			bufTag := tag

			go func() {
				crt.New(bufTag)
				wg.Done()
			}()

		}
	}

	wg.Wait()
	b.StopTimer()
}
