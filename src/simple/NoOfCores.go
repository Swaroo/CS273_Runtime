package main

import (
	"runtime"
	"fmt"
)

func main(){
	noOfCores := runtime.NumCPU()
	fmt.Print("Number of cores : ")
	fmt.Println(noOfCores)
	fmt.Println("Setting the number of cores to 2")
	newNoOfCores := runtime.GOMAXPROCS(1)
	fmt.Print("New number of cores : ")
	fmt.Println(newNoOfCores)
	fmt.Println(runtime.NumCPU())
}
