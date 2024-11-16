package PointerFactory

import (
	"time"
)

////////////////////////////////////

func (obj *GlobalObj) loop() {

	secT := 60 - time.Now().Second() - obj.timeNow().Second()
	tickerStart := time.NewTicker(time.Second * time.Duration(secT))
	defer tickerStart.Stop()

	tickerUPD := time.NewTicker(time.Minute)
	defer tickerUPD.Stop()

	//

	defer func(obj *GlobalObj) { obj.isActive = false }(obj)

	upd := func() {
		obj.minute.Add(0)

		for r, _ := range obj.groups {
			obj.groups[r].Store(0)
		}
	}

	//

	for {
		select {
		case <-tickerStart.C:
			tickerUPD = time.NewTicker(time.Minute)
			tickerStart.Stop()
			upd()

		case <-tickerUPD.C:
			upd()

		case <-obj.ctx.Done():
			return

		default:
			obj.isActive = true
		}
	}
}
