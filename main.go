package main

import (
	"fmt"
	"runtime"
)
import "time"

const (
	// how many ints we'll send across the channel
	count int = 1e7
	// this is the slice capacity
	batch int = 100
)

func main() {
	start := time.Now()
	ch := make(chan []int)
	go func() {
		slice := make([]int, 0, batch)
		for i := 1; i <= count; i++ {
			slice = append(slice, i)
			if len(slice) == cap(slice) {
				ch <- slice
				slice = make([]int, 0, batch)
			}
		}
		ch <- slice
		close(ch)
	}()
	sum := 0
	for slice := range ch {
		for _, i := range slice {
			sum += i
		}
	}
	elapsed := time.Since(start)
	ok := sum == count*(count+1)/2
	fmt.Printf("sum 1..%d: %d (sum verified:%v) took %s\n", count, sum, ok, elapsed)
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
