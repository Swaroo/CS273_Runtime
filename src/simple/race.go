package main

import (
	"fmt"
)

func main() {
	var sharedVariable int
	go func() {
		sharedVariable++
	}()
	if sharedVariable == 0 {
		fmt.Printf("The value is %v .\n", sharedVariable)
	}
}
