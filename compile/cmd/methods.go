package cmd

import (
	"fmt"
	pf "github.com/Bookshelf-Writer/PointerFactory"
	"strings"
	"time"
)

func validType(str string) bool {
	if len(str) == 0 {
		printError("TYPE", "The type cannot be empty")
		return false
	}

	if len(str) > 1 {
		printError("TYPE", "A type can be only one character")
		return false
	}

	_, status := wp.TypeMAP[[]rune(str)[0]]
	if !status {
		printError("TYPE", "This type is not supported")
		return false
	}

	return true
}

func parseTime(dateStr string) (time.Time, error) {
	if len(dateStr) == 0 {
		return time.Time{}, fmt.Errorf("Date cannot be empty")
	}

	formats := []string{
		time.RFC3339, // 2006-01-02T15:04:05Z07:00
		"2006-01-02T15:04:05.000Z",
		"2006-01-02",
		"20060102T150405Z",
		"2006-01-02 15:04:05",
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, dateStr)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("Date format does not match ISO-8601")
}

//////////////////////////////////////////////////////////////////////////////

func methodVER() {
	if *flagsBuf.jsonFormat.value {
		printJson(infoObj())
		return
	}

	fmt.Println(wp.GlobalName, wp.GlobalVersion)
}

func methodTypes() {
	buf := []string{}
	for char, t := range wp.TypeMAP {
		if t != pf.TypeNone {
			buf = append(buf, string(char))
		}
	}

	if *flagsBuf.jsonFormat.value {
		printJson(buf)
		return
	}

	fmt.Println(buf)
}

func methodType(str string) {
	if validType(str) {
		printOK()
	}
}

func methodValidKey(key string, info bool) {
	obj, err := wp.Decode(key)
	if err != nil {
		printError("VALID", err.Error())
		return
	}

	if !info {
		printOK()
		return
	}

	if *flagsBuf.jsonFormat.value {
		type retInfoObj struct {
			Type    string `json:"type"`
			Server  uint16 `json:"server"`
			Pointer uint64 `json:"pointer"`
		}
		printJson(retInfoObj{
			Type:    obj.Type(),
			Server:  obj.ServerN(),
			Pointer: obj.PointerN(),
		})
	} else {
		fmt.Println(strings.Repeat(" ", 8-len("Type"))+"Type:", obj.Type())
		fmt.Println(strings.Repeat(" ", 8-len("Server"))+"Server:", obj.ServerN())
		fmt.Println(strings.Repeat(" ", 8-len("Pointer"))+"Pointer:", obj.PointerN())
	}
}

func methodCreateNum(types string, server uint, pointer uint64) {
	if !validType(types) {
		return
	}

	if server == 0 {
		printError("CREATE-N", "The server cannot be 0")
		return
	}
	if server > 1294 {
		printError("CREATE-N", "The server can't be bigger than 1294")
		return
	}

	if pointer == 0 {
		printError("CREATE-N", "The pointer cannot be 0")
		return
	}
	if pointer > 99999999999999999 {
		printError("CREATE-N", "The pointer can't be bigger than 1294")
		return
	}

	//

	crt := wp.InitCreator(uint16(server))
	defer crt.Close()

	t, _ := wp.TypeMAP[[]rune(types)[0]]

	printString(crt.NewFixed(t, pointer).String())
}

func methodCreateDate(types string, server uint, date string, count uint) {
	if !validType(types) {
		return
	}

	if server == 0 {
		printError("CREATE-D", "The server cannot be 0")
		return
	}
	if server > 1294 {
		printError("CREATE-D", "The server can't be bigger than 1294")
		return
	}

	if count == 0 {
		count = 1
	}
	if count > 9999 {
		printError("CREATE-D", "The count can't be bigger than 9999")
		return
	}

	dateTime, err := parseTime(date)
	if err != nil {
		printError("CREATE-D", err.Error())
		return
	}

	if dateTime.Year() < wp.YearPoint {
		printError("CREATE-D", "A year cannot be less than the starting point")
		return
	}

	//.//

	crt := wp.InitCreatorDate(uint16(server), dateTime)
	defer crt.Close()

	t, _ := wp.TypeMAP[[]rune(types)[0]]
	bufKeys := []string{}

	for i := uint(0); i < count; i++ {
		bufKeys = append(bufKeys, crt.New(t).String())
	}

	if *flagsBuf.jsonFormat.value {
		printJson(bufKeys)
	} else {
		for _, uid := range bufKeys {
			fmt.Println(uid)
		}
	}
}
