package main

import (
	"testing"
)

func BenchmarkParkingProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recParkOuter(seq, 0, 0, recParkInner)
	}
	return
}
