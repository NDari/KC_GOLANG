package main

import "fmt"

type Dog struct{}
type Cat struct{}

func (d *Dog) Speak() string { return "Arf! Arf!" }
func (d *Cat) Speak() string { return "Meow!" }

type NoiseMaker interface {
	Speak() string
}

func makeNoise(a NoiseMaker) { fmt.Println("The animal says:", a.Speak()) }

func main() {
	myDog := new(Dog)
	myCat := new(Cat)
	makeNoise(myDog)
	makeNoise(myCat)
}
