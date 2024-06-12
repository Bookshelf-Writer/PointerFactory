package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func infoObj() map[string]string {
	return map[string]string{
		"name":  wp.GlobalName,
		"ver":   wp.GlobalVersion,
		"upd":   wp.GlobalDateUpdate,
		"num":   strconv.Itoa(wp.NumBase),
		"point": strconv.Itoa(wp.YearPoint),
	}
}

func printOK() {
	if *flagsBuf.jsonFormat.value {
		fmt.Printf("{\"status\": true}")
	} else {
		fmt.Println("OK")
	}
}

func printString(text string) {
	if *flagsBuf.jsonFormat.value {
		fmt.Printf("{\"status\": true, \"data\":\"%s\"}", text)
	} else {
		fmt.Println(text)
	}
}

func printError(typeError string, textError string) {
	if *flagsBuf.jsonFormat.value {
		fmt.Printf("{\"status\": false, \"type\":\"%s\" \"error\":\"%s\"}", typeError, textError)
	} else {
		fmt.Println("[ERROR]["+typeError+"]\t", textError)
	}
}

func printJson(obj interface{}) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		printError("JSON", err.Error())
		return
	}

	fmt.Printf("{\"status\": true, \"data\":%s}", jsonData)
}
