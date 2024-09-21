package PointerFactory

////////////////////////////////////

type chObj struct {
	group     rune
	retOffset chan uint32
}

//

func (obj *GlobalObj) sendChan(group rune) uint32 {
	ch := make(chan uint32)
	defer close(ch)

	obj.ch <- &chObj{
		group:     group,
		retOffset: ch,
	}

	return <-ch
}
