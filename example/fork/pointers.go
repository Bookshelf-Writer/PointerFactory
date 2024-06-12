package main

import pf "github.com/Bookshelf-Writer/PointerFactory"

const (
	TypeN1 pf.TypeTag = iota + 1
	TypeN2
	TypeN3
)

func init() {
	SYS.TypeMAP['a'] = TypeN1
	SYS.TypeMAP['b'] = TypeN2
	SYS.TypeMAP['c'] = TypeN3
}
