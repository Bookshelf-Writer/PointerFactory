package PointerFactory

////////////////////////////////////

type chObj struct {
	group     rune
	retOffset chan uint32
}

//

func (obj *GlobalObj) sendChan(group rune) uint32 {
	ch := make(chan uint32)

	obj.ch <- &chObj{
		group:     group,
		retOffset: ch,
	}

	offset := <-ch
	close(ch)

	return offset
}
