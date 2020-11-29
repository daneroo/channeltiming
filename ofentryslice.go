package channeltiming

import (
	"fmt"
	"time"
)

// GenerateEntrySlices is ...
func GenerateEntrySlices(n int, batch int) chan []Entry {
	ch := make(chan []Entry)
	go func() {
		slice := make([]Entry, 0, batch)
		for i := 1; i <= n; i++ {
			// entry := Entry{Watt: i}
			entry := Entry{Stamp: time.Unix(int64(i), 0), Watt: i}
			slice = append(slice, entry)
			if len(slice) == cap(slice) {
				ch <- slice
				slice = make([]Entry, 0, batch)
			}
		}
		ch <- slice
		close(ch)
	}()
	return ch
}

// ConsumeEntrySlices is ...
func ConsumeEntrySlices(ch <-chan []Entry, batch int, expect int) int {
	defer TimeTrack(time.Now(), fmt.Sprintf("[%d]Entry", batch), expect)
	sum := 0
	for slice := range ch {
		for _, entry := range slice {
			i := entry.Watt
			sum += i
		}
	}
	return sum
}
