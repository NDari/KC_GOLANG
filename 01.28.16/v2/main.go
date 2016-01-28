package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PI(samples int) float64 {
	var inside int = 0

	for i := 0; i < samples; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := r.Float64()
		y := r.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(samples)

	return ratio * 4
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Our value of Pi after 100 runs:\t\t\t", PI(100))
	fmt.Println("Our value of Pi after 1,000 runs:\t\t", PI(1000))
	fmt.Println("Our value of Pi after 10,000 runs:\t\t", PI(10000))
	fmt.Println("Our value of Pi after 100,000 runs:\t\t", PI(100000))
	fmt.Println("Our value of Pi after 1,000,000 runs:\t\t", PI(1000000))
	fmt.Println("Our value of Pi after 10,000,000 runs:\t\t", PI(10000000))
	fmt.Println("Our value of Pi after 100,000,000 runs:\t\t", PI(100000000))
}
