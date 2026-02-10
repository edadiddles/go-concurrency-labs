package main

import (
	"testing"
)

func check_report(report AccessReport, n int) []string {
	errs := make([]string, 0)

	if len(report.access_records) < n {
		errs = append(errs, "missing access records")
	} else if len(report.access_records) > n {
		errs = append(errs, "too many access records")
	}

	if !report.initialized {
		errs = append(errs, "never initialized")
	} else {
		init_time := report.initialization_time
		did_init := 0
		for _, rec := range report.access_records {
			if rec.performed_initialization {
				did_init += 1
			}

			if init_time.Compare(rec.access_time) > -1 {
				errs = append(errs, "accessed before initialization")
			}
		}

		if did_init > 1 {
			errs = append(errs, "too many inits")
		}
	}

	return errs
}

func TestBasic(t *testing.T) {
	n := 10
	low := 50
	high := 100

	for range 10 {
		report := run_once(n, low, high)

		errs := check_report(report, n)
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestManyWorkers(t *testing.T) {
	n := 100
	low := 50
	high := 100

	for range 10 {
		report := run_once(n, low, high)

		errs := check_report(report, n)
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestWideRange(t *testing.T) {
	n := 10
	low := 10
	high := 5000

	report := run_once(n, low, high)

	errs := check_report(report, n)
	for _, err := range errs {
		t.Error(err)
	}
}
