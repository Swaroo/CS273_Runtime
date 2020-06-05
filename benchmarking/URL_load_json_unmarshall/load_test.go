package main

import (
	"testing"
)

//Benchmarking Serial Multiply
func BenchmarkSerial(b *testing.B) {
	for index := 0; index < b.N; index++ {
		serial()
	}
}

//Benchmarking Parallel Multiply
func BenchmarkParallel(b *testing.B) {
	for index := 0; index < b.N; index++ {
		parallel()
	}
}
