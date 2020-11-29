package channeltiming

import (
	"time"
)

// GenerateInts is ...
func GenerateInts(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// ConsumeInts is ...
func ConsumeInts(ch <-chan int, expect int) int {
	defer TimeTrack(time.Now(), "int", expect)
	sum := 0
	for i := range ch {
		sum += i
	}
	return sum
}
