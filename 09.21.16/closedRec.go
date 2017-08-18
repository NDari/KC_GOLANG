package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 3
	}()
	fmt.Println(<-ch)
	close(ch)

	// what do we get from this?
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
