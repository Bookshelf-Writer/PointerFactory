package PointerFactory

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

const TypeTest TypeTag = 44

func init() {
	wp.TypeMAP['s'] = TypeTest
	wp.TypeMAP['1'] = 100
	wp.TypeMAP['2'] = 101
	wp.TypeMAP['3'] = 102
	wp.TypeMAP['4'] = 103
	wp.TypeMAP['5'] = 104
	wp.TypeMAP['6'] = 105
	wp.TypeMAP['7'] = 106
	wp.TypeMAP['8'] = 107
	wp.TypeMAP['9'] = 108
}

//////////////////////////////////////////////////////////`

func TestGlobal(t *testing.T) {
	wp.StartTests(t)
}

func BenchmarkGlobal(b *testing.B) {
	wp.StartBenchmarks(b)
}

// TestCreateMin Прогон теста на сохранение уникальности через минуту
func TestCreateMin(t *testing.T) {
	t.SkipNow()

	period := time.Duration(10 * time.Second)
	timeWork := time.Duration(1*time.Minute + 30*time.Second)

	//Инициализация метода фабрики указателей
	crt := wp.InitCreator(888)
	defer crt.Close()

	//Внутренние переменные для валидации теста или перехвата STOP
	bufID := []string{}
	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, os.Interrupt, syscall.SIGTERM)

	//Инициализация периода и времени работы
	ticker := time.NewTicker(period)
	endTime := time.Now().Add(timeWork)
	defer ticker.Stop()

	for {
		select {

		//Тик по периоду
		case timeN := <-ticker.C:
			id := crt.New(TypeNone)
			bufID = append(bufID, id.String())
			fmt.Println(timeN.Minute(), timeN.Second(), "\t|\t", id.String(), id.Uint())

		// Финальный перехват
		case <-time.After(time.Until(endTime)):
			mapID := map[string]int{}
			for _, id := range bufID {
				_, status := mapID[id]
				if !status {
					mapID[id] = 1
				} else {
					mapID[id] += 1
					t.Error(fmt.Errorf("Dublicate ID: %s ", id))
				}
			}
			return

		//Перехват команды остановки
		case <-closeCh:
			ticker.Stop()
			t.Errorf("BREACK")
			return

		}
	}
}
