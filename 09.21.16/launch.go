package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read one byte
		abort <- struct{}{}
	}()
	tick := time.Tick(1 * time.Second) // goroutine leak
	for x := 2; x > 0; x-- {
		fmt.Println(x)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("launch aborted")
			return
		}
	}
	fmt.Println("Launching...")
	panic("oh nooo")
}
