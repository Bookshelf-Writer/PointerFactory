package cmd

import (
	"flag"
	pf "github.com/Bookshelf-Writer/PointerFactory"
)

func InitCMD(wpImport *pf.WrapperObj) {
	if !isValid {
		printError("BUILD", " Flags not init")
		return
	}

	wp = *wpImport

	addFlag(&flagsBuf.ver)
	addFlag(&flagsBuf.jsonFormat)

	addFlag(&flagsBuf.printTypes)
	addFlag(&flagsBuf.validType)

	addFlag(&flagsBuf.createPointer)
	addFlag(&flagsBuf.setType)
	addFlag(&flagsBuf.setServer)
	addFlag(&flagsBuf.setPointer)

	addFlag(&flagsBuf.setDateTime)
	addFlag(&flagsBuf.setCount)

	addFlag(&flagsBuf.validString)
	addFlag(&flagsBuf.printInfo)

	flag.Parse()

	switchRoad()
}

func switchRoad() {
	switch {

	case *flagsBuf.ver.value:
		methodVER()
	case *flagsBuf.printTypes.value:
		methodTypes()
	case len(*flagsBuf.validType.value) > 0:
		methodType(*flagsBuf.validType.value)

	case len(*flagsBuf.validString.value) > 0:
		methodValidKey(*flagsBuf.validString.value, *flagsBuf.printInfo.value)

	case *flagsBuf.createPointer.value:
		if len(*flagsBuf.setDateTime.value) > 0 {
			methodCreateDate(*flagsBuf.setType.value, *flagsBuf.setServer.value, *flagsBuf.setDateTime.value, *flagsBuf.setCount.value)
		} else {
			methodCreateNum(*flagsBuf.setType.value, *flagsBuf.setServer.value, *flagsBuf.setPointer.value)
		}

	default:
		printHelp()
	}
}
