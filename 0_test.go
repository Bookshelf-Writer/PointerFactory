package PointerFactory

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

////////////////////////////////////

var (
	groups = []rune{
		'u',
		'r',
	}
	startPoint        = time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)
	base       uint8  = 36
	cluster    uint16 = math.MaxUint16 - 1
)

////////

func TestBase(t *testing.T) {
	num := rand.Uint64()
	uid := NumToString(num)

	if StringToNum(uid) != num {
		t.Error("Invalid Methods convert")
	}
}

////

func TestNew(t *testing.T) {
	uid, err := New(groups, cluster, base, startPoint)
	if err != nil {
		t.Fatal(err)
	}
	defer uid.Close()

	_, err = uid.New(groups[0])
	if err != nil {
		t.Error(err)
	}
}

func TestValid(t *testing.T) {

	uid, err := New(groups, cluster, base, startPoint)
	if err != nil {
		t.Fatal(err)
	}
	defer uid.Close()

	err = uid.IsValid("u07sgxy008")
	if err != nil {
		t.Error(err)
	}

	err = uid.IsValid("r07sgxy011")
	if err != nil {
		t.Error(err)
	}

	err = uid.IsValid("g085fb055")
	if err == nil {
		t.Error("Invalid Valid")
	}
}

////

func BenchmarkGlobal(b *testing.B) {
	uid, err := New(groups, cluster, base, time.Date(-2000, 1, 1, 1, 1, 1, 1, time.UTC))
	if err != nil {
		b.Fatal(err)
	}
	defer uid.Close()

	//

	b.Run("New", func(bx *testing.B) {
		for i := 0; i < bx.N; i++ {
			uid.New(groups[0])
		}
	})

	//

	b.Run("ConvertToNum", func(bx *testing.B) {
		for i := 0; i < bx.N; i++ {
			StringToNum("u07sgxy008")
		}
	})

	x := rand.Uint64()
	b.Run("ConvertToString", func(bx *testing.B) {
		for i := 0; i < bx.N; i++ {
			NumToString(x)
		}
	})
}
