package main

import (
	"math/rand"
	"testing"
)

//Benchmarking Serial Multiply
func BenchmarkSerialMultiply(b *testing.B) {
	arr := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		arr[i] = rand.Int() % 1000
	}
	for index := 0; index < b.N; index++ {
		_ = SerialMultiply(arr)
	}
}

//Benchmarking Parallel Multiply
func BenchmarkParallelMultiply(b *testing.B) {
	arr := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		arr[i] = rand.Int() % 10000
	}
	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]
	ch := make(chan int)
	for index := 0; index < b.N; index++ {
		go multiplyAndSendToChannel(arr1, ch)
		go multiplyAndSendToChannel(arr2, ch)
		_ = <-ch * <-ch
	}
}
