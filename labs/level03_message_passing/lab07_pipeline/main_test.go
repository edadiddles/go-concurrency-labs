package main

import (
	"testing"
)

func check_pipeline_records(num_items, num_stages int, recs []PipelineRecord) []string {
	errs := make([]string, 0)

	if len(recs) < num_items {
		errs = append(errs, "items missing")
	} else if len(recs) > num_items {
		errs = append(errs, "too many records")
	}

	for _, rec := range recs {
		if len(rec.stages) < num_stages {
			errs = append(errs, "stages missing")
			continue
		} else if len(rec.stages) > num_stages {
			errs = append(errs, "too many stages")
			continue
		}

		for i := range num_stages-1 {
			if rec.stages[i] != rec.stages[i+1]-1 {
				errs = append(errs, "stages not executed in order")
			}
		}
	}

	return errs
}

func TestPipeline(t *testing.T) {
	n := 10
	s := 5
	low := 50
	high := 100

	recs := pipeline(n, s, low, high)

	errs := check_pipeline_records(n, s, recs)

	for _, err := range errs {
		t.Error(err)
	}
}

func TestSkewedItems(t *testing.T) {
	n := 1000
	s := 5
	low := 50
	high := 100

	recs := pipeline(n, s, low, high)

	errs := check_pipeline_records(n, s, recs)

	for _, err := range errs {
		t.Error(err)
	}
}

func TestSkewedStages(t *testing.T) {
	n := 10
	s := 200
	low := 50
	high := 100

	recs := pipeline(n, s, low, high)

	errs := check_pipeline_records(n, s, recs)

	for _, err := range errs {
		t.Error(err)
	}
}

func TestLargeLoad(t *testing.T) {
	n := 100
	s := 20
	low := 100
	high := 2000

	recs := pipeline(n, s, low, high)

	errs := check_pipeline_records(n, s, recs)

	for _, err := range errs {
		t.Error(err)
	}
}
