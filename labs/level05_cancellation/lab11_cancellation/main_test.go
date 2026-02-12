package main

import (
	"testing"
	"time"
)

func check_report(report CancelReport, n int) []string {
	errs := make([]string, 0)

	cancel_time := time.Now()
	if !report.cancel_time.IsZero(){
		cancel_time = report.cancel_time
	}

	if len(report.worker_reports) > n {
		errs = append(errs, "too many worker reports")
	} else if len(report.worker_reports) < n {
		errs = append(errs, "missing worker reports")
	}


	for _, worker_report := range report.worker_reports {
		if worker_report.exit_time.Compare(cancel_time) > 0 && !worker_report.interrupted {
			errs = append(errs, "worker finished after cancellation without interruption")
		} else if worker_report.exit_time.Compare(cancel_time) < 0 && worker_report.interrupted {
			errs = append(errs, "worker interrupted before cancellation")
		}

	}
	return errs
}

func TestBasic(t *testing.T) {

	n := 10
	k := 10
	p := 300
	low := 10
	high := 50

	for range 10 {
		report := cancellation(n, k, p, low, high)

		errs := check_report(report, n)
		for err := range errs {
			t.Error(err)
		}
	}
}

func TestManyWorkers(t *testing.T) {

	n := 100
	k := 10
	p := 300
	low := 10
	high := 50

	for range 10 {
		report := cancellation(n, k, p, low, high)

		errs := check_report(report, n)
		for err := range errs {
			t.Error(err)
		}
	}
}

func TestManyWorkUnits(t *testing.T) {

	n := 10
	k := 100
	p := 3000
	low := 10
	high := 50

	for range 10 {
		report := cancellation(n, k, p, low, high)

		errs := check_report(report, n)
		for err := range errs {
			t.Error(err)
		}
	}
}
