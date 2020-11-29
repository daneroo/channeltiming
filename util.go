package channeltiming

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

var (
	// SilentTimeTrack is used to silence TimeTrack for Benchmarks
	SilentTimeTrack  bool = false
	anInt            int  = 0
	sizeOfInt        int  = int(reflect.TypeOf(anInt).Size())
	sizeOfIntPointer int  = int(reflect.TypeOf(&anInt).Size())
)

// PrintSizes is ...
func PrintSizes() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Sizeof(int): %d\n", sizeOfInt)
	fmt.Printf("time.Unix(1e7): %v\n", time.Unix(1e7, 0))
	fmt.Printf("time.Unix(1e8): %v\n", time.Unix(1e8, 0))
	fmt.Printf("time.Unix(1e9): %v\n", time.Unix(1e9, 0))

}

// CheckSum is ...
func CheckSum(n int, sum int) error {
	expected := n * (n + 1) / 2
	var err error
	ok := sum == expected
	if !ok {
		err = fmt.Errorf("Expected %d got %d", expected, sum)
	}
	return err
}

// CheckSumErr is ...
func CheckSumErr(n int, sum int) {
	if err := CheckSum(n, sum); err != nil {
		log.Fatal(err)
	}
}

// TimeTrack is ...
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
