package main

import "fmt"

func main() {

	intChan := make(chan int) // unbuffered channel which holds an int
	defer close(intChan)

	go func() {
		intChan <- 3 // send 3 into the channel
	}()
	x := <-intChan // receive from the channel, and assign to x. In this case, x is 3.
	fmt.Println(x)

}
