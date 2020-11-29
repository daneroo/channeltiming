package channeltiming

import (
	"time"
)

// GenerateEntries is ...
func GenerateEntries(n int) chan Entry {
	ch := make(chan Entry)
	go func() {
		for i := 1; i <= n; i++ {
			// entry := Entry{Watt: i}
			entry := Entry{Stamp: time.Unix(int64(i), 0), Watt: i}
			ch <- entry
		}
		close(ch)
	}()
	return ch
}

// ConsumeEntries is ...
func ConsumeEntries(ch <-chan Entry, expect int) int {
	defer TimeTrack(time.Now(), "Entry", expect)
	sum := 0
	for entry := range ch {
		i := entry.Watt
		sum += i
	}
	return sum
}
