package Fibbonachi

import (
	"testing"
)
var input uint = 45

func BenchmarkRunOrig(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		RunOrig(input)
	}
}

func BenchmarkRunOptimized1(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		RunOptimized1(input)
	}
}

func BenchmarkRunOptimized2(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		RunOptimized2(input)
	}
}

func BenchmarkRunOptimized3(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		RunOptimized3(input)
	}
}