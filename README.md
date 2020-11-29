# Channel Communications Timing Experiment

## TODO

- Sync.pool (Olivier Gagnon Gophers-Slack/#golang-newbies)
  https://play.golang.org/p/bgG5tEh5tq
  https://play.golang.org/p/MsFkuoWMOk

## Running the experiment

    # to run as main
    go run main.go of*.go

    # run tests
    go test -v

    # run tests and benchmarks
    go test --bench .

## Presentation

Question:
-I am reading 7,000,000 records from a MySQL database (in a goroutine)
-This takes about 9.5s
-When I also send the row `struct{time.Time,int}` to a channel, and consume them from a concurrent goroutine, the time goes to 12.8s

Snippets: Timing for channel comms.

    chan int: https://play.golang.org/p/AGVkCCd3Zz
    chan *int: https://play.golang.org/p/xL6o-onkbp
    chan []int https://play.golang.org/p/OR3OiDaZVU

Does this seem like a lot of overhead for channel communications?
(_That means channel communications take about a third of the time it takes to stream the data from the database, which seems high_)

I have done this snippet: https://play.golang.org/p/AGVkCCd3Zz
which just times sending 10,000,000 ints
_Of course it takes 0 time in the playground!_
but locally I get a similar result as the database example above: 2.7s for `1e7 ints`
`sum 1..10000000: 50000005000000 took 2.698350052s`
Also `GOMAXPROCS: 4` (on a Macbook Air)

I just tried something else: sending slices of 100 ints across the channel helps a lot:
sending 1e7 ints, in 100 int slice chunks, brings the time down from _2.7s_ to under _100ms_
This is the new code: https://play.golang.org/p/OR3OiDaZVU
`sum 1..10000000: 50000005000000 (sum verified:true) took 96.601072ms`
