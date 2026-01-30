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
		worker_duration := int64(s.end.Sub(s.start).Milliseconds())
		sequential_duration = sequential_duration + worker_duration

		// check start < end
		if s.start.Compare(s.end) >= 0 {
			errs = append(errs, "start time does not occur before end time")
		}

		// check duration >= low
		if s.end.Sub(s.start).Milliseconds() < int64(low) {
			errs = append(errs, "duration does not exceed minimum value")
		}
	}

	// check N workers completed
	if len(stats) != 10000 {
		errs = append(errs, "Not all logs returned")
	}

	// check total measured runtime does not exceed max duration by an appreciable amount.
	//if end.Sub(start).Milliseconds() > int64(high + 10)  {
	//	errs = append(errs, "measured runtime exceeded expectations")
	//}

	// check for proof of concurrency
	is_concurrent := false
	for i := 1; i < len(stats); i++ {
		for j := 0; j < i; j++ {
			if stats[i].start.Compare(stats[j].start) < 0 && stats[j].end.Compare(stats[i].end) < 0 {
				is_concurrent = true
			} else if stats[i].start.Compare(stats[j].start) > 0 && stats[j].end.Compare(stats[i].end) > 0 {
				is_concurrent = true
			}
		}
	}
	if !is_concurrent {
		errs = append(errs, "proof of concurrency dne")
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
