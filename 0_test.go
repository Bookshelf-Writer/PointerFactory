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

	obj, err := New(groups, 0, time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(obj)

	time.Sleep(110 * time.Second)
	obj.Close()
}
