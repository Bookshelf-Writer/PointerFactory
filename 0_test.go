package PointerFactory

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	groups := []rune{
		'a',
		'3',
	}

	obj, err := New(groups, 0, 36, time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(obj)

	for i := 0; i < 10; i++ {
		uid, e := obj.New(groups[0])
		fmt.Println("GENERATE\t", uid, e)
		if e == nil {
			fmt.Println("VALID\t", obj.IsValid(uid))
		}

		time.Sleep(10 * time.Second)
	}

	obj.Close()
}
