package main

import (
	"fmt"
	"time"
)

func GenerateSlices(n int, batch int) chan []int {
	ch := make(chan []int)
	go func() {
		slice := make([]int, 0, batch)
		for i := 1; i <= n; i++ {
			slice = append(slice, i)
			if len(slice) == cap(slice) {
				ch <- slice
				slice = make([]int, 0, batch)
			}
		}
		ch <- slice
		close(ch)
	}()
	return ch
}

func ConsumeSlices(ch <-chan []int, batch int, expect int) int {
	defer TimeTrack(time.Now(), fmt.Sprintf("[%d]int", batch), expect)
	sum := 0
	for slice := range ch {
		for _, i := range slice {
			sum += i
		}
	}
	return sum
}
