package main

import "testing"

func TestGlobal(t *testing.T) {
	SYS.StartTests(t)
}

func BenchmarkGlobal(b *testing.B) {
	SYS.StartBenchmarks(b)
}
