package main

import (
	"fmt"
	"sync"
)

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
