package main

import (
	"testing"
)

func check_records(records []Record, capacity, total_expected_records int) []string {
	errs := make([]string, 0)

	n := 0
	cnt := 0
	for _, rec := range records {
		n += 1
		switch rec.op {
		case "+":
			cnt += 1
		case "-":
			cnt -= 1
		}

		if cnt > capacity {
			errs = append(errs, "buffer overflow")
		} else if cnt < 0 {
			errs = append(errs, "buffer underflow")
		}

	}

	if n < total_expected_records {
		errs = append(errs, "did not capture all expected records")
	} else if n > total_expected_records {
		errs = append(errs, "too many records captured")
	}

	return errs
}

func TestBoundedBufferManyProducers(t *testing.T) {
	p := 10
	c := 2
	b := 10
	n := 100

	records := bounded_buffer(p, c, b, n)

	errs := check_records(records, b, 2*p*n)

	for _, err := range errs {
		t.Error(err)
	}
}

func TestBoundedBufferManyConsumers(t *testing.T) {
	p := 2
	c := 10
	b := 10
	n := 100

	records := bounded_buffer(p, c, b, n)

	errs := check_records(records, b, 2*p*n)

	for _, err := range errs {
		t.Error(err)
	}
}

func TestBoundedBufferLargeCapacity(t *testing.T) {
	p := 10
	c := 10
	b := 1000
	n := 1000

	records := bounded_buffer(p, c, b, n)

	errs := check_records(records, b, 2*p*n)

	for _, err := range errs {
		t.Error(err)
	}
}

func TestBoundedBufferManyRecords(t *testing.T) {
	p := 10
	c := 10
	b := 100
	n := 10000

	records := bounded_buffer(p, c, b, n)

	errs := check_records(records, b, 2*p*n)

	for _, err := range errs {
		t.Error(err)
	}
}
