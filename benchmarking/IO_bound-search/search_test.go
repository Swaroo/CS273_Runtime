package main

import (
	"testing"
)

//Benchmarking Serial Multiply
func BenchmarkSerialSearch(b *testing.B) {
	for index := 0; index < b.N; index++ {
		serialSearch()
	}
}

//Benchmarking Parallel Multiply
func BenchmarkParallelSearch(b *testing.B) {
	for index := 0; index < b.N; index++ {
		parallelSearch()
	}
}
