package main

import (
	"testing"
)

func baselineSum(n int) error {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return CheckSum(n, sum)
}

func TestSum(t *testing.T) {
	n := 100
	if err := baselineSum(n); err != nil {
		t.Error(err)
	}
}

func TestChanInt(t *testing.T) {
	SilentTimeTrack = true
	var n int = 1e5
	if err := CheckSum(n, ConsumeInts(GenerateInts(n), n)); err != nil {
		t.Error(err)
	}
}
func TestChanIntPointer(t *testing.T) {
	SilentTimeTrack = true
	var n int = 1e5
	if err := CheckSum(n, ConsumeIntPointers(GenerateIntPointers(n), n)); err != nil {
		t.Error(err)
	}
}
func TestChanInSlices100(t *testing.T) {
	SilentTimeTrack = true
	var n int = 1e5
	var batch int = 100
	if err := CheckSum(n, ConsumeSlices(GenerateSlices(n, batch), batch, n)); err != nil {
		t.Error(err)
	}
}

func BenchmarkSum(b *testing.B) {
	SilentTimeTrack = true
	b.SetBytes(int64(sizeOfInt))
	if err := baselineSum(b.N); err != nil {
		b.Error(err)
	}
}

func BenchmarkChanInt(b *testing.B) {
	SilentTimeTrack = true
	b.SetBytes(int64(sizeOfInt))
	if err := CheckSum(b.N, ConsumeInts(GenerateInts(b.N), b.N)); err != nil {
		b.Error(err)
	}
}

func BenchmarkChanIntPointer(b *testing.B) {
	SilentTimeTrack = true
	b.SetBytes(int64(sizeOfIntPointer))
	if err := CheckSum(b.N, ConsumeIntPointers(GenerateIntPointers(b.N), b.N)); err != nil {
		b.Error(err)
	}
}

func benchForSliceSize(batch int, b *testing.B) {
	SilentTimeTrack = true
	b.SetBytes(int64(sizeOfInt))

	if err := CheckSum(b.N, ConsumeSlices(GenerateSlices(b.N, batch), batch, b.N)); err != nil {
		b.Error(err)
	}
}

func BenchmarkChanIntSlices1(b *testing.B) {
	benchForSliceSize(1, b)
}
func BenchmarkChanIntSlices10(b *testing.B) {
	benchForSliceSize(10, b)
}
func BenchmarkChanIntSlices100(b *testing.B) {
	benchForSliceSize(100, b)
}
func BenchmarkChanIntSlices1000(b *testing.B) {
	benchForSliceSize(1000, b)
}
func BenchmarkChanIntSlices10000(b *testing.B) {
	benchForSliceSize(10000, b)
}

func BenchmarkChanEntries1000(b *testing.B) {
	SilentTimeTrack = true
	b.SetBytes(int64(sizeOfInt))
	batch := 1000
	if err := CheckSum(b.N, ConsumeEntries(GenerateEntries(b.N, batch), batch, b.N)); err != nil {
		b.Error(err)
	}
}
