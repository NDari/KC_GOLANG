package main

import (
	"fmt"
	"sync"
)

func main() {
	numRows := 8
	numCols := 10
	s := make([][]float64, numRows)
	for i := range s {
		s[i] = make([]float64, numCols)
	}

	var wg sync.WaitGroup
	for i := 0; i < numRows; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < len(s[i]); j++ {
				s[i][j] = float64(i)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(s)
}
