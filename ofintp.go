package channeltiming

import (
	"time"
)

func GenerateIntPointers(n int) chan *int {
	ch := make(chan *int)
	go func() {
		for i := 1; i <= n; i++ {
			ii := i // cant send i's address?
			ch <- &ii
		}
		close(ch)
	}()
	return ch
}

func ConsumeIntPointers(ch <-chan *int, expect int) int {
	defer TimeTrack(time.Now(), "*int", expect)
	sum := 0
	for i := range ch {
		sum += *i
	}
	return sum
}
