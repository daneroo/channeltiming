package main

import (
	ct "github.com/daneroo/channeltiming"
)

const (
// how many ints we'll send across the channel
// count int = 1e7
)

func main() {

	var count int = 1e7

	ct.CheckSumErr(count, ct.ConsumeInts(ct.GenerateInts(count), count))
	ct.CheckSumErr(count, ct.ConsumeIntPointers(ct.GenerateIntPointers(count), count))

	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e0), 1e0, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e2), 1e1, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e2), 1e2, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e3), 1e3, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e4), 1e4, count))

	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count, 1e0), 1e0, count))
	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count, 1e1), 1e1, count))
	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count, 1e2), 1e2, count))
	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count, 1e3), 1e3, count))
	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count, 1e4), 1e4, count))

}
