package main

import "fmt"

func main() {
	ints := make(chan int)
	sqrs := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			ints <- x
		}
		close(ints) // dont need to close channels, unless you want to signal that you're done
	}()
	go func() {
		for x := range ints { // Range over a channel means keep going until its closed
			sqrs <- x * x
		}
		close(sqrs)
	}()
	for x := range sqrs { // using the range over channel again
		fmt.Println(x)
	}
}
