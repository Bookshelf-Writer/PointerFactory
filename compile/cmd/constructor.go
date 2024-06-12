package cmd

import (
	"flag"
)

//////////////////////////////////////////////////////////////////////////////

var errFlagMAP = map[string]bool{}

func (obj *flagInfoObj) flagMapInit() bool {
	_, status := errFlagMAP[obj.name]
	if status {
		printError("BUILD", "Dublicate before init:"+obj.name)
		return true
	}

	errFlagMAP[obj.name] = true
	return false
}

//.//

func (obj *flagInfoObj) addInfo(name string, usage string) {
	_, status := errMAP[name]
	if status {
		printError("BUILD", "Dublicate key:"+name)
		isValid = false
	}

	errMAP[name] = true
	obj.name = name
	obj.usage = usage
}

func addFlag(input addFlagInterface) {
	input.add()
}

//////////////////////////////////////////////////////////////////////////////

type addFlagInterface interface {
	add()
}

//.//

func (obj *flagPointBoolObj) add() {
	if obj.info.flagMapInit() {
		return
	}
	obj.value = flag.Bool(obj.info.name, false, obj.info.usage)
}

func (obj *flagPointStringObj) add() {
	if obj.info.flagMapInit() {
		return
	}
	obj.value = flag.String(obj.info.name, "", obj.info.usage)
}

func (obj *flagPointUintObj) add() {
	if obj.info.flagMapInit() {
		return
	}
	obj.value = flag.Uint(obj.info.name, 0, obj.info.usage)
}
func (obj *flagPointUint64Obj) add() {
	if obj.info.flagMapInit() {
		return
	}
	obj.value = flag.Uint64(obj.info.name, 0, obj.info.usage)
}
