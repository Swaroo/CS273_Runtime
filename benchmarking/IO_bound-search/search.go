package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func serialSearch() {
	for i := 0; i < 50; i++ {
		filename := "./search_directory/file_" + strconv.Itoa(i) + ".txt"
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Some error in reading file")
			return
		}
		_ = strings.Count(string(data), "en")

	}
}

func search(filename string, ch chan int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		//fmt.Println("Some error in reading file")
		ch <- 0
	}
	ch <- strings.Count(string(data), "en")

}
func parallelSearch() {
	ch := make(chan int, 50)
	for i := 0; i < 50; i++ {
		filename := "./search_directory/file_" + strconv.Itoa(i) + ".txt"
		go search(filename, ch)
	}
	for i := 0; i < 10; i++ {
		_ = <-ch
	}

}

func main() {
	serialSearch()
	parallelSearch()
}
