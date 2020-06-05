package main

func SerialMultiply(arr []int) int {
	res := 1
	for _, val := range arr {
		res *= val
	}
	return res
}

func multiplyAndSendToChannel(arr []int, ch chan int) {
	res := 1
	for _, val := range arr {
		res *= val
	}
	ch <- res
}

func ParallelMultiply(arr []int) int {
	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]
	ch := make(chan int)
	go multiplyAndSendToChannel(arr1, ch)
	go multiplyAndSendToChannel(arr2, ch)
	return <-ch * <-ch

}

func main() {
	// 	arr := make([]int, 1000)
	// 	for i := 0; i < 1000; i++ {
	// 		arr[i] = rand.Int() % 1000
	// 	}
	// 	fmt.Println(SerialMultiply)
	// 	fmt.Println(ParallelMultiply)
}
