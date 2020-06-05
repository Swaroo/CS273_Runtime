package main

import (
	"fmt"
)

func Square(x int) int {
	result := x * x
	return result
}

func main(){
	fmt.Println("Benchmarking")
	//fmt.Println(Square(5))
}