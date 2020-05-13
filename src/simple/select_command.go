package main

import "fmt"
import "strconv"

func simpleProcessor(process, kill chan int){
	process_unit_id := 1
	for {
		select{
		case process <- process_unit_id:
			process_unit_id += 1
		case <-kill:
			fmt.Println("Killing Process")
			return 
		}
	}
}

func main(){
	process := make(chan int)
	kill := make(chan int)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println("Processing id " + strconv.Itoa(<-process))
		}
		kill <- 0
	}()
	simpleProcessor(process, kill)
}