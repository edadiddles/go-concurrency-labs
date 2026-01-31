package main

import (
	"testing"
	"time"
)

func check_stats(t *testing.T, stats []Statistics, n int, start, end time.Time) {
	max_worker_duration := int64(0)
	sum_worker_duration := int64(0)
	concurrency_exists := false
	for i, s1 := range stats {
		sum_worker_duration += s1.duration
		if s1.duration > max_worker_duration {
			max_worker_duration = s1.duration
		}

		if !concurrency_exists {
			for _, s2 := range stats[:i] {
				if s1.start.Compare(s2.start) > 0 && s1.end.Compare(s2.end) < 0 {
					concurrency_exists = true
				} else if s1.start.Compare(s2.start) < 0 && s1.end.Compare(s2.end) > 0 {
					concurrency_exists = true
				}
			}
		}
	}

	// check all workers returned records
	if n < len(stats) {
		t.Error("worker records are missing")
	} else if n > len(stats) {
		t.Error("too many worker records")
	}

	// check wall-clock runtime is correctly bounded
	wallclock_duration := end.Sub(start).Milliseconds()
	if max_worker_duration > wallclock_duration || sum_worker_duration < wallclock_duration {
		t.Error("wall-clock runtime is not bounded by max and sum of worker durations")
	}

	// check proof of concurrency
	if !concurrency_exists {
		t.Error("Proof of concurrency not detected")
	}
}

func TestBlocking(t *testing.T) {
	n := 10000
	low := 100
	high := 1000

	start := time.Now()
	stats := saturated_execution(n, low, high, "blocking")
	end := time.Now()
	
	check_stats(t, stats, n, start, end)
}

func TestCPU(t *testing.T) {	
	n := 10000
	low := 100
	high := 1000

	start := time.Now()
	stats := saturated_execution(n, low, high, "cpu")
	end := time.Now()
	
	check_stats(t, stats, n, start, end)
}

func TestMixed(t *testing.T) {	
	n := 10000
	low := 100
	high := 1000

	start := time.Now()
	stats := saturated_execution(n, low, high, "mixed")
	end := time.Now()
	
	check_stats(t, stats, n, start, end)
}
