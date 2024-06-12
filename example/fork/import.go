package main

import pf "github.com/Bookshelf-Writer/PointerFactory"

type WrapperObj struct {
	pf.WrapperObj
}

type CreatorObj struct {
	pf.CreatorObj
}

type PointerObj struct {
	pf.PointerObj
}

var SYS = pf.CreateWrapper()

func init() {
	SYS.GlobalVersion = GlobalVersion
	SYS.GlobalDateUpdate = GlobalDateUpdate
	SYS.GlobalName = GlobalName

	SYS.NumBase = NumBase
	SYS.YearPoint = YearPoint
}
