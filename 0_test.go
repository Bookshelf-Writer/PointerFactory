package PointerFactory

import (
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
	base       int32  = 36
	cluster    uint16 = 0
)

////////

func TestBase(t *testing.T) {
	num := rand.Uint64()
	uid := NumToString(num, base)

	if StringToNum(uid, base) != num {
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

	for !uid.IsActive() {
		time.Sleep(10 * time.Millisecond)
	}

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

	for !uid.IsActive() {
		time.Sleep(10 * time.Millisecond)
	}

	err = uid.IsValid("u07qpgy0ya")
	if err != nil {
		t.Error(err)
	}

	err = uid.IsValid("r07qpgy0v7")
	if err != nil {
		t.Error(err)
	}

	err = uid.IsValid("g085fb055")
	if err == nil {
		t.Error("Invalid Valid")
	}
}

////

func BenchmarkNew(b *testing.B) {
	uid, err := New(groups, cluster, base, startPoint)
	if err != nil {
		b.Fatal(err)
	}
	defer uid.Close()

	for !uid.IsActive() {
		time.Sleep(10 * time.Millisecond)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uid.New(groups[0])
	}
}
