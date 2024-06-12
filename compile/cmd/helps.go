package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

var helpArr = map[string][]*flagInfoObj{}

func initHelpArr() {
	globalArr := []*flagInfoObj{
		&flagsBuf.ver.info,
		&flagsBuf.jsonFormat.info,
		&flagsBuf.printTypes.info,
		&flagsBuf.validType.info,
	}
	helpArr["global"] = globalArr

	validArr := []*flagInfoObj{
		&flagsBuf.validString.info,
		&flagsBuf.printInfo.info,
	}
	helpArr["valid"] = validArr

	createNumArr := []*flagInfoObj{
		&flagsBuf.createPointer.info,
		&flagsBuf.setType.info,
		&flagsBuf.setServer.info,
		&flagsBuf.setPointer.info,

		&flagsBuf.setDateTime.info,
		&flagsBuf.setCount.info,
	}
	helpArr["create"] = createNumArr

}

func printHelp() {
	if *flagsBuf.jsonFormat.value {
		bufMap := map[string]map[string]string{
			"info": infoObj(),
		}

		for group, arr := range helpArr {
			buf := map[string]string{}
			for _, info := range arr {
				buf[info.name] = info.usage
			}
			bufMap[group] = buf
		}

		printJson(bufMap)
		return
	}

	fmt.Println(wp.GlobalName, wp.GlobalVersion, "\t["+strconv.Itoa(wp.YearPoint)+"]N^"+strconv.Itoa(wp.NumBase), "\n ")

	for group, arr := range helpArr {
		fmt.Println("  ", group+":")

		for _, info := range arr {
			fmt.Println("\t", "-"+info.name+strings.Repeat(" ", 8-len(info.name)), info.usage)
		}

		fmt.Println("")
	}
}
