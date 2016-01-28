package main

import (
	"math/rand"
	"runtime"
	"time"
)

func PI(samples int) float64 {
	var inside int

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

func MultiPI(samples int) float64 {
	cpus := runtime.NumCPU()
	threadSamples := samples / cpus
	results := make(chan float64, cpus)

	for j := 0; j < cpus; j++ {
		go func() {
			var inside int
			for i := 0; i < threadSamples; i++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				x, y := r.Float64(), r.Float64()

				if x*x+y*y <= 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}
