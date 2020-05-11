package main

import "fmt"
import "strconv"

//Finds the product within a slice in the array, and returns the product to a channel
func multiply(arr []int, ch chan int){
	out:= 1
	for _,val := range arr{
		out *= val
	}
	//Send multiplied output back to the channel
	ch <- out
}

func main() {
	arr:= []int{1,2,3,4,5,6}
	
	//Instantiate a channel which passes integer value
	ch:= make(chan int)

	// Call two goroutines to calculate product of each half of the array.
	// The two goroutines ruc in parallel
	go multiply(arr[:len(arr)/2], ch)
	go multiply(arr[len(arr)/2:], ch)

	//This instruction will block until there is some data in the channel
	first_product, second_product:= <-ch, <-ch

	fmt.Println("First half product "+strconv.Itoa(second_product))
	fmt.Println("Second half product "+strconv.Itoa(first_product))
	fmt.Println("Final Product "+strconv.Itoa(first_product * second_product))
}