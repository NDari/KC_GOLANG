package main

import "fmt"

type Dog struct{}

func (d *Dog) Speak() string { return "Arf! Arf!" }

type NoiseMaker interface {
	Speak() string
}

func makeNoise(a NoiseMaker) {
	fmt.Println("The animal says:", a.Speak())
}

func main() {
	myDog := new(Dog)
	makeNoise(myDog)
}
