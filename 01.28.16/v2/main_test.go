package main

import (
	"testing"
)

const samples = 100000

func BenchmarkPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PI(samples)
	}
}
