package main

import (
	"fmt"
	"sync"
)

func main() {
	numRows := 800
	numCols := 10
	numWorkers := 4
	jobs := make(chan int)

	s := make([][]float64, numRows)
	for i := range s {
		s[i] = make([]float64, numCols)
	}

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for job := range jobs {
				for j := 0; j < len(s[job]); j++ {
					s[job][j] = float64(job)
				}
			}
			wg.Done()
		}()
	}

	for i := 0; i < numRows; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	fmt.Println(s)
}
