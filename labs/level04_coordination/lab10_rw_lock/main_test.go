package main

import (
	"testing"
	"time"
)

func check_report(report RWReport, r, w, k int) []string {
	errs := make([]string, 0)

	if report.num_reads < r*k {
		errs = append(errs, "missing reads")
	} else if report.num_reads > r*k {
		errs = append(errs, "too many reads")
	}

	if report.num_writes < w*k {
		errs = append(errs, "missing writes")
	} else if report.num_reads > r*k {
		errs = append(errs, "too many writes")
	}

	prev_val := 0
	prev_time := time.Now().Add(-1*24*time.Hour)
	for _, rec := range report.access_records {
		if prev_time.Compare(rec.access_time) > -1 {
			errs = append(errs, "access records out of order")
		}

		if rec.access_type == "write" {
			prev_val = rec.val
			prev_time = rec.access_time
			continue
		}

		if prev_val != rec.val {
			errs = append(errs, "read wrong value")
		}
	}


	return errs
}

func TestEmpty(t *testing.T) {
	r := 10
	w := 10
	k := 10
	low := 50
	high := 100

	report := rw_lock(r, w, k, low, high)

	errs := check_report(report, r, w, k)
	for _, err := range errs{
		t.Error(err)
	}
}
