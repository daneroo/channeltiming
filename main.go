package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)
import "time"

const (
// how many ints we'll send across the channel
// count int = 1e7
)

var (
	SilentTimeTrack  bool = false // to silence TimeTrack for Benchmarks
	anInt            int  = 0
	sizeOfInt        int  = int(reflect.TypeOf(anInt).Size())
	sizeOfIntPointer int  = int(reflect.TypeOf(&anInt).Size())
)

func CheckSum(n int, sum int) error {
	expected := n * (n + 1) / 2
	var err error
	ok := sum == expected
	if !ok {
		err = fmt.Errorf("Expected %d got %d", expected, sum)
	}
	return err
}

func CheckSumErr(n int, sum int) {
	if err := CheckSum(n, sum); err != nil {
		log.Fatal(err)
	}
}

func TimeTrack(start time.Time, what string, n int) {
	if SilentTimeTrack {
		return
	}
	name := fmt.Sprintf("Chanel of %12s", what)
	elapsed := time.Since(start)
	if n > 0 {
		rate := float64(n) / elapsed.Seconds()
		log.Printf("%s n: %d rate: %.1e/s time: %s", name, n, rate, elapsed)
	} else {
		log.Printf("%s took %s", name, elapsed)
	}
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Sizeof(int): %d\n", sizeOfInt)
	fmt.Printf("time.Unix(1e7): %v\n", time.Unix(1e7, 0))
	fmt.Printf("time.Unix(1e8): %v\n", time.Unix(1e8, 0))
	fmt.Printf("time.Unix(1e9): %v\n", time.Unix(1e9, 0))

	var count int = 1e7

	CheckSumErr(count, ConsumeInts(GenerateInts(count), count))
	CheckSumErr(count, ConsumeIntPointers(GenerateIntPointers(count), count))
	CheckSumErr(count, ConsumeSlices(GenerateSlices(count, 1), 1, count))
	CheckSumErr(count, ConsumeSlices(GenerateSlices(count, 10), 10, count))
	CheckSumErr(count, ConsumeSlices(GenerateSlices(count, 100), 100, count))
	CheckSumErr(count, ConsumeSlices(GenerateSlices(count, 1e3), 1e3, count))
	CheckSumErr(count, ConsumeSlices(GenerateSlices(count, 1e4), 1e4, count))

	CheckSumErr(count, ConsumeEntries(GenerateEntries(count, 1), 1, count))
	CheckSumErr(count, ConsumeEntries(GenerateEntries(count, 1e3), 1e3, count))

}
