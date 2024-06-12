package main

import (
	pf "github.com/Bookshelf-Writer/PointerFactory"
	fWASM "github.com/Bookshelf-Writer/PointerFactory/compile/wasm"
)

var SYS = pf.CreateWrapper()

/* Добавляем типы */

const (
	TypeN1 pf.TypeTag = iota + 1
	TypeN2
	TypeS
)

func init() {
	SYS.TypeMAP['1'] = TypeN1
	SYS.TypeMAP['2'] = TypeN2
	SYS.TypeMAP['s'] = TypeS

	SYS.GlobalVersion = "T.1.1"
	SYS.GlobalDateUpdate = "2000-20-20"
	SYS.GlobalName = "TestWASM"
}

//////////////////////////////////////////////////////////////////////////////

func main() {
	fWASM.InitWASM(&SYS)
}
