package main

import (
	"fmt"
	"log"
	"runtime"
)
import "time"

const (
	// how many ints we'll send across the channel
	count int = 1e7
)

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

func ConsumeInts(ch <-chan int) int {
	defer TimeTrack(time.Now(), "int", count)
	sum := 0
	for i := range ch {
		sum += i
	}
	return sum
}

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

func ConsumeIntPointers(ch <-chan *int) int {
	defer TimeTrack(time.Now(), "*int", count)
	sum := 0
	for i := range ch {
		sum += *i
	}
	return sum
}

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

func ConsumeSlices(ch <-chan []int, batch int) int {
	defer TimeTrack(time.Now(), fmt.Sprintf("[%d]int", batch), count)
	sum := 0
	for slice := range ch {
		for _, i := range slice {
			sum += i
		}
	}
	return sum
}

func Check_sum(n int, sum int) {
	expected := count * (count + 1) / 2
	ok := sum == expected
	if !ok {
		log.Fatalf("Expected %d got %d", expected, sum)
	}

}

func TimeTrack(start time.Time, what string, count int) {
	name := fmt.Sprintf("Chanel of %10s", what)
	elapsed := time.Since(start)
	if count > 0 {
		rate := float64(count) / elapsed.Seconds()
		log.Printf("%s took %s, count: %d rate: %.1e/s", name, elapsed, count, rate)
	} else {
		log.Printf("%s took %s", name, elapsed)
	}
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// ch_i := GenerateInts(count)
	// sum := ConsumeInts(ch_i)
	// Check_sum(sum)

	Check_sum(count, ConsumeInts(GenerateInts(count)))
	Check_sum(count, ConsumeIntPointers(GenerateIntPointers(count)))
	Check_sum(count, ConsumeSlices(GenerateSlices(count, 1), 1))
	Check_sum(count, ConsumeSlices(GenerateSlices(count, 10), 10))
	Check_sum(count, ConsumeSlices(GenerateSlices(count, 100), 100))
	Check_sum(count, ConsumeSlices(GenerateSlices(count, 1e3), 1e3))
	Check_sum(count, ConsumeSlices(GenerateSlices(count, 1e4), 1e4))

	// start := time.Now()
	// ch := GenerateSlices(count, batch)
	// sum := 0
	// for slice := range ch {
	// 	for _, i := range slice {
	// 		sum += i
	// 	}
	// }
	// elapsed := time.Since(start)
	// ok := sum == count*(count+1)/2
	// fmt.Printf("sum 1..%d: %d (sum verified:%v) took %s\n", count, sum, ok, elapsed)
}
