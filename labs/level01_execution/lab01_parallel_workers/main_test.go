package main

import (
	"testing"
	"runtime"
)


func TestConcurrencySmallRange(t *testing.T) {
	iterations := 10
	n := 10000
	low := 10
	high := 15

	for iter := range iterations {
		_ = iter
		parallel_workers(n, low, high)
	}
}

func TestConcurrencyLargeRange(t *testing.T) {
	iterations := 10
	n := 10000
	low := 10
	high := 1000

	for iter := range iterations {
		_ = iter
		parallel_workers(n, low, high)
	}
}

func TestConcurrencySchedulingPressure(t *testing.T) {
	runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(runtime.NumCPU())

	iterations := 10
	n := 10000
	low := 10
	high := 100
	
	for iter := range iterations {
		_ = iter
		parallel_workers(n, low, high)
	}
}
