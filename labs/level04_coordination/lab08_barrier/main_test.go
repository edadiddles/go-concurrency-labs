package main

import (
	"testing"
	"time"
)

type DepartureStats struct {
	min time.Time
	max time.Time
}

func check_reports(n, p int, reports []WorkerReport) []string {
	errs := make([]string, 0)

	if len(reports) < n {
		errs = append(errs, "missing worker reports")
	} else if len(reports) > n {
		errs = append(errs, "too many worker reports")
	}

	phase_departures := make([]DepartureStats, 0)
	for _, report := range reports {
		if len(report.barriers) < p {
			errs = append(errs, "missing barrier reports")
		} else if len(report.barriers) > p {
			errs = append(errs, "too many barrier reports")
		}

		for _, b_report := range report.barriers {
			if len(phase_departures) <= b_report.id {
				phase_departure := DepartureStats{
					min: b_report.departure_time,
					max: b_report.departure_time,
				}
				phase_departures = append(phase_departures, phase_departure)
				continue
			}

			if phase_departures[b_report.id].min.After(b_report.departure_time) {
				phase_departures[b_report.id].min = b_report.departure_time
			}
			if phase_departures[b_report.id].max.Before(b_report.departure_time) {
				phase_departures[b_report.id].max = b_report.departure_time
			}
		}
	}

	for _, departures := range phase_departures {
		if departures.max.Sub(departures.min) > time.Duration(10)*time.Millisecond {
			errs = append(errs, "departures not within threshold")
		}
	}
	return errs
}

func TestBasic(t *testing.T) {
	n := 10
	p := 5
	low := 100
	high := 300

	reports := barrier(n, p, low, high)
	errs := check_reports(n, p, reports)
	for _, err := range errs {
		t.Error(err)
	}
}

func TestManyWorkers(t *testing.T) {
	n := 100
	p := 5
	low := 100
	high := 300

	reports := barrier(n, p, low, high)
	errs := check_reports(n, p, reports)
	for _, err := range errs {
		t.Error(err)
	}
}

func TestManyPhases(t *testing.T) {
	n := 10
	p := 50
	low := 100
	high := 300

	reports := barrier(n, p, low, high)
	errs := check_reports(n, p, reports)
	for _, err := range errs {
		t.Error(err)
	}
}

func TestLoad(t *testing.T) {
	n := 100
	p := 50
	low := 10
	high := 1000

	reports := barrier(n, p, low, high)
	errs := check_reports(n, p, reports)
	for _, err := range errs {
		t.Error(err)
	}
}
