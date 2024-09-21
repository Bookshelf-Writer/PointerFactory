package PointerFactory

import (
	"time"
)

////////////////////////////////////

func (obj *GlobalObj) loop() {
	obj.ch = make(chan *chObj, 1)
	defer close(obj.ch)

	secT := 60 - time.Now().Second() - obj.timeNow().Second()
	tickerStart := time.NewTicker(time.Second * time.Duration(secT))
	defer tickerStart.Stop()

	tickerUPD := time.NewTicker(time.Minute)
	defer tickerUPD.Stop()

	//

	defer func(obj *GlobalObj) { obj.isActive = false }(obj)

	upd := func() {
		obj.minute += 1

		for _, group := range obj.groupsBuf {
			obj.groups[group] = 0
		}
	}

	//

	for {
		obj.isActive = true

		select {
		case <-tickerStart.C:
			tickerUPD = time.NewTicker(time.Minute)
			tickerStart.Stop()
			upd()

		case <-tickerUPD.C:
			upd()

		case pointer := <-obj.ch:
			pointer.retOffset <- obj.groups[pointer.group]
			obj.groups[pointer.group]++

		case <-obj.ctx.Done():
			return
		}
	}
}
