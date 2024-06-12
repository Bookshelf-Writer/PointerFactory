package PointerFactory

import (
	"sync"
	"time"
)

const IncrementMAX = 10000

func init() {
	wp.IncrementMAX = IncrementMAX
}

// IncrementObj Хранилише инкрементов
type IncrementObj struct {
	sync.Mutex

	inc uint16 // Последний инкремент за период
}

//////////////////////////////////////////////////////////////////////////////

// CreatorObj Класс создания уникальных инкрементов
type CreatorObj struct {
	pointRef time.Time // Точка отсчета
	server   uint16    // Указатель сервера

	syncWG    sync.WaitGroup // Блокировка для корректного завершения
	chanClose chan struct{}  // Канал закрытия всех системных потоков
	timeNow   time.Time      // Реальное время, для оптимизации асинхронных расчетов
	duration  uint64         // Количество минут от точки отсчета

	incrementBuf map[TypeTag]*IncrementObj // Буфер для поддержания уникализации
}

/* Инициализация метода создания новых индексов */
func (w *WrapperObj) InitCreator(server uint16) *CreatorObj {
	return w.InitCreatorDate(server, time.Now().UTC())
}

/* Инициализация для конкретной даты */
func (w *WrapperObj) InitCreatorDate(server uint16, timeNOW time.Time) *CreatorObj {
	obj := CreatorObj{}

	obj.server = server
	obj.pointRef = time.Date(wp.YearPoint, 0, 0, 0, 0, 0, 0, time.UTC)

	obj.chanClose = make(chan struct{}, 1)
	obj.syncWG.Add(1)

	obj.timeNow = time.Date(timeNOW.Year(), timeNOW.Month(), timeNOW.Day(), timeNOW.Hour(), 0, 0, 0, timeNOW.Location())
	obj.duration = uint64(obj.timeNow.Sub(obj.pointRef).Minutes())
	go obj.startUpdater()

	obj.incrementBuf = make(map[TypeTag]*IncrementObj, len(w.TypeMAP)-1)
	for _, typeTag := range w.TypeMAP {
		obj.incrementBuf[typeTag] = &IncrementObj{
			inc: 0,
		}
	}

	return &obj
}

/* Закрытие всех потоков и очистка буфера */
func (obj *CreatorObj) Close() {
	close(obj.chanClose)
	obj.syncWG.Wait()

	for key, _ := range obj.incrementBuf {
		obj.incrementBuf[key] = nil
	}
}

//.//

/* Создание фиксированного указателя */
func (obj *CreatorObj) NewFixed(types TypeTag, pointer uint64) *PointerObj {
	retObj := PointerObj{}

	retObj.t = types
	retObj.s = obj.server
	retObj.p = pointer

	return &retObj
}

/* Создание нового уникального указателя */
func (obj *CreatorObj) New(types TypeTag) *PointerObj {
	retObj := PointerObj{}

	retObj.t = types
	retObj.s = obj.server
	retObj.p = obj.uniqueID(types)

	return &retObj
}

//////////////////////////////////////////////////////////////////////////////

// uniqueID Генериция уникального ID
func (obj *CreatorObj) uniqueID(types TypeTag) uint64 {
	retID := obj.duration * wp.IncrementMAX

	//Работа с буфером инкрементов
	if incObj, ok := obj.incrementBuf[types]; ok {
		incObj.Lock()

		// Плюсуем инкремент
		retID += uint64(incObj.inc)

		incObj.inc += 1
		incObj.Unlock()
	}

	return retID
}

// startUpdater обновление каждую минуту
func (obj *CreatorObj) startUpdater() {
	defer obj.syncWG.Done()

	//Запуск тикера обновлений с необходимым смещением
	secT := 60 - time.Now().UTC().Second() - obj.timeNow.Second()
	tickerStart := time.NewTicker(time.Second * time.Duration(secT))
	defer tickerStart.Stop()

	//Установка тикера для минутного
	tickerUPD := time.NewTicker(time.Minute)
	defer tickerUPD.Stop()

	//обнуление буфера
	updBuf := func() {
		obj.duration += 1

		for _, bufObj := range obj.incrementBuf {
			bufObj.Lock()
			bufObj.inc = 0
			bufObj.Unlock()
		}
	}

	for {
		select {
		case <-tickerStart.C: //Первый запуск в четкое время
			tickerUPD = time.NewTicker(time.Minute)
			tickerStart.Stop()
			updBuf()

		case <-tickerUPD.C: //Обновляшка по периоду
			updBuf()

		case <-obj.chanClose: //Закрытие
			return

		}
	}
}
