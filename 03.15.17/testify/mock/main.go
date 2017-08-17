package main

import "fmt"

type Thing1 struct {
}

func (t *Thing1) add(x, y int) int {
	return x + y
}

type Adder interface {
	add(int, int) int
}

func callAdd(t2 Adder) int {
	return t2.add(1, 2)
}

func main() {
	r1 := &Thing1{}

	fmt.Println(callAdd(r1))
}
