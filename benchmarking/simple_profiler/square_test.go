package main

import(
	"testing"
)

func BenchmarkSquare(b *testing.B){
	for index:=0; index < b.N; index++{
		Square(5)
	}
}