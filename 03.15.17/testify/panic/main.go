package main

import (
	"fmt"
	"os"
)

func cat(x, y string) string {
	if y == "panic" {
		panic("Oh snap!")
	}
	return x + y
}

func main() {
	fmt.Println(cat(os.Args[1], os.Args[2]))
}
