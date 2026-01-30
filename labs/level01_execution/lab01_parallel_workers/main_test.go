package main

import (
	"testing"
	"runtime"
	"time"
)

func check_stats(n, low, high int, start, end time.Time, stats []Statistics) []string {
	errs := make([]string, 0)


	var sequential_duration int64 = 0
	var min_start time.Time
	var max_end time.Time
	for _, s := range stats {
		if s.start.Compare(min_start) < 0 {
			min_start = s.start
		}
		if s.end.Compare(max_end) > 0 {
			max_end = s.end
		}
		sequential_duration = sequential_duration + (s.end.Sub(s.start).Milliseconds())
	}

	// check N workers completed
	if len(stats) != 10000 {
		errs = append(errs, "Not all logs returned")
	}

	return errs
}

func TestConcurrencySmallRange(t *testing.T) {
	iterations := 10
	n := 10000
	low := 10
	high := 15

	for iter := range iterations {
		_ = iter
		start := time.Now()
		stats := parallel_workers(n, low, high)
		end := time.Now()

		for _, e := range check_stats(n, low, high, start, end, stats) {
			t.Error(e)
		}

	}
}

func TestConcurrencyLargeRange(t *testing.T) {
	iterations := 10
	n := 10000
	low := 10
	high := 1000

	for iter := range iterations {
		_ = iter
		start := time.Now()
		stats := parallel_workers(n, low, high)
		end := time.Now()
		
		for _, e := range check_stats(n, low, high, start, end, stats) {
			t.Error(e)
		}
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
		start := time.Now()
		stats := parallel_workers(n, low, high)
		end := time.Now()
		
		for _, e := range check_stats(n, low, high, start, end, stats) {
			t.Error(e)
		}
	}
}
