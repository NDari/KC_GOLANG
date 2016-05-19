package main

import "fmt"

type Dog struct{}

func (d *Dog) Speak() string { return "Arf! Arf!" }

func makeNoise(d *Dog) {
	fmt.Println("The animal says:", d.Speak())
}

func main() {
	myDog := new(Dog)
	makeNoise(myDog)
}
