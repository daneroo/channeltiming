package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	n := 100
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	expected := n * (n + 1) / 2
	if sum != expected {
		t.Errorf("Sum is wrong")
	}
}

func BenchmarkChanInt(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	sum := 0
	for i := range ch {
		sum += i
	}
}
func BenchmarkChanIntPointer(b *testing.B) {
	ch := make(chan *int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- &i
		}
		close(ch)
	}()
	sum := 0
	for i := range ch {
		sum += *i
	}
}
func BenchmarkChanIntSlices(b *testing.B) {
	batch := 10
	ch := make(chan []int)
	go func() {
		slice := make([]int, 0, batch)
		for i := 0; i < b.N; i++ {
			slice = append(slice, i)
			if len(slice) == cap(slice) {
				ch <- slice
				slice = make([]int, 0, batch)
			}
		}
		ch <- slice
		close(ch)
	}()
	sum := 0
	for slice := range ch {
		for _, i := range slice {
			sum += i
		}
	}
}
