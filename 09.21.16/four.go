package main

import "fmt"

func main() {
	// on to buffered channels... more on this later

	ch := make(chan int, 0) // capacity 0... unbuffered channel
	defer close(ch)

	// allocate buffered channel with a capacity of 10
	ch10 := make(chan int, 10)
	defer close(ch10)

	fmt.Println(cap(ch10))
	fmt.Println(len(ch10))
	ch10 <- 10
	fmt.Println(len(ch10))

}
