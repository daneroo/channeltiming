# Channel Communications Timing Experiment

## Running the experiment

```bash
# to run as main
go run cmd/main.go

# run tests
go test -v

# run tests and benchmarks
go test --bench .
```

## Benchmarks

Program output

```bash
$ go run cmd/main.go
GOMAXPROCS: 8
Sizeof(int): 8
Sizeof(&int): 8
Sizeof(Entry): 32
time.Unix(1e7): 1970-04-26 13:46:40 -0400 EDT
time.Unix(1e8): 1973-03-03 04:46:40 -0500 EST
time.Unix(1e9): 2001-09-08 21:46:40 -0400 EDT
2020-11-29T03:08:17.464Z - Chanel of          int n: 10000000 rate:    3.2M/s time: 3.093056535s
2020-11-29T03:08:21.012Z - Chanel of         *int n: 10000000 rate:    2.8M/s time: 3.547594892s
2020-11-29T03:08:24.136Z - Chanel of        Entry n: 10000000 rate:    3.2M/s time: 3.123901462s
2020-11-29T03:08:24.136Z -
2020-11-29T03:08:27.604Z - Chanel of       [1]int n: 10000000 rate:    2.9M/s time: 3.468035824s
2020-11-29T03:08:27.682Z - Chanel of      [10]int n: 10000000 rate:  128.7M/s time: 77.721686ms
2020-11-29T03:08:27.760Z - Chanel of     [100]int n: 10000000 rate:  127.5M/s time: 78.427139ms
2020-11-29T03:08:27.804Z - Chanel of    [1000]int n: 10000000 rate:  227.7M/s time: 43.91152ms
2020-11-29T03:08:27.838Z - Chanel of   [10000]int n: 10000000 rate:  292.1M/s time: 34.240222ms
2020-11-29T03:08:27.838Z -
2020-11-29T03:08:30.946Z - Chanel of        Entry n: 10000000 rate:    3.2M/s time: 3.107337941s
2020-11-29T03:08:35.073Z - Chanel of     [1]Entry n: 10000000 rate:    2.4M/s time: 4.127524525s
2020-11-29T03:08:35.652Z - Chanel of    [10]Entry n: 10000000 rate:   17.3M/s time: 578.53299ms
2020-11-29T03:08:35.860Z - Chanel of   [100]Entry n: 10000000 rate:   48.0M/s time: 208.402072ms
2020-11-29T03:08:36.014Z - Chanel of  [1000]Entry n: 10000000 rate:   65.0M/s time: 153.838789ms
2020-11-29T03:08:36.173Z - Chanel of [10000]Entry n: 10000000 rate:   62.8M/s time: 159.253376ms
```

``

```bash
$ go test --bench .
goos: darwin
goarch: amd64
pkg: github.com/daneroo/channeltiming
BenchmarkSum-8                    1000000000  0.367 ns/op 21809.38 MB/s

BenchmarkChanInt-8                4135244        286.0 ns/op   27.99 MB/s
BenchmarkChanIntPointer-8         3693967        328.0 ns/op   24.42 MB/s
BenchmarkChanIntSlices1-8         3676503        341.0 ns/op   23.43 MB/s
BenchmarkChanIntSlices10-8        30612799        37.8 ns/op  211.37 MB/s
BenchmarkChanIntSlices100-8       156610070       7.50 ns/op 1066.56 MB/s
BenchmarkChanIntSlices1000-8      286873788       4.17 ns/op 1917.76 MB/s
BenchmarkChanIntSlices10000-8     370633615       3.31 ns/op 2414.90 MB/s

BenchmarkChanEntries-8            4129183        295.0 ns/op  108.33 MB/s
BenchmarkChanEntrySlices1-8       3214213        377.0 ns/op   84.79 MB/s
BenchmarkChanEntrySlices10-8      21651412        55.3 ns/op  578.15 MB/s
BenchmarkChanEntrySlices100-8     59819854        21.1 ns/op 1517.88 MB/s
BenchmarkChanEntrySlices200-8     62558145        19.2 ns/op 1663.12 MB/s
BenchmarkChanEntrySlices500-8     76730238        16.4 ns/op 1950.38 MB/s
BenchmarkChanEntrySlices1000-8    76182301        15.6 ns/op 2047.96 MB/s
BenchmarkChanEntrySlices10000-8   69555434        15.8 ns/op 2024.58 MB/s
PASS
ok  github.com/daneroo/channeltiming 21.896s
```

## Process

Question:
-I am reading 7,000,000 records from a MySQL database (in a goroutine)
-This takes about 9.5s
-When I also send the row `struct {time.Time,int}` to a channel, and consume them from a concurrent goroutine, the time goes to 12.8s

Snippets: Timing for channel communications.

- chan int: <https://play.golang.org/p/AGVkCCd3Zz>
- chan \*int: <https://play.golang.org/p/xL6o-onkbp>
- chan []int: <https://play.golang.org/p/OR3OiDaZVU>

Does this seem like a lot of overhead for channel communications?
(_That means channel communications take about a third of the time it takes to stream the data from the database, which seems high_)

I have done this snippet: <https://play.golang.org/p/AGVkCCd3Zz>
which just times sending 10,000,000 ints
_Of course it takes 0 time in the playground!_
but locally I get a similar result as the database example above: 2.7s for `1e7 ints`
`sum 1..10000000: 50000005000000 took 2.698350052s`
Also `GOMAXPROCS: 4` (on a Macbook Air)

I just tried something else: sending slices of 100 ints across the channel helps a lot.
Sending 1e7 ints, in 100 int slice chunks, brings the time down from _2.7s_ to under _100ms_
This is the new code: <https://play.golang.org/p/OR3OiDaZVU>

`sum 1..10000000: 50000005000000 (sum verified:true) took 96.601072ms`
