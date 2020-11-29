package main

import (
	"fmt"
	"log"
	"time"

	ct "github.com/daneroo/channeltiming"
)

const (
	fmtRFC3339Millis = "2006-01-02T15:04:05.000Z07:00"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format(fmtRFC3339Millis) + " - " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	ct.PrintSizes()
	var count int = 1e7

	ct.CheckSumErr(count, ct.ConsumeInts(ct.GenerateInts(count), count))
	ct.CheckSumErr(count, ct.ConsumeIntPointers(ct.GenerateIntPointers(count), count))
	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count), count))

	log.Println("")
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e0), 1e0, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e2), 1e1, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e2), 1e2, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e3), 1e3, count))
	ct.CheckSumErr(count, ct.ConsumeSlices(ct.GenerateSlices(count, 1e4), 1e4, count))

	log.Println("")
	ct.CheckSumErr(count, ct.ConsumeEntries(ct.GenerateEntries(count), count))
	ct.CheckSumErr(count, ct.ConsumeEntrySlices(ct.GenerateEntrySlices(count, 1e0), 1e0, count))
	ct.CheckSumErr(count, ct.ConsumeEntrySlices(ct.GenerateEntrySlices(count, 1e1), 1e1, count))
	ct.CheckSumErr(count, ct.ConsumeEntrySlices(ct.GenerateEntrySlices(count, 1e2), 1e2, count))
	ct.CheckSumErr(count, ct.ConsumeEntrySlices(ct.GenerateEntrySlices(count, 1e3), 1e3, count))
	ct.CheckSumErr(count, ct.ConsumeEntrySlices(ct.GenerateEntrySlices(count, 1e4), 1e4, count))

}
