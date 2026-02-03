package main

import (
	"testing"
)

func TestManyProducers(t *testing.T) {
	p := 20
	w := 5
	n := 10
	low := 10
	high := 50


	stats := work_queues(p, w, n, low, high)

	if len(stats) < p*n {
		t.Error("not all work items processed")
	} else if len(stats) > p*n {
		t.Error("too many work items processed")
	}
}

func TestManyConsumers(t *testing.T) {
	p := 5
	w := 20
	n := 10
	low := 50
	high := 100


	stats := work_queues(p, w, n, low, high)

	if len(stats) < p*n {
		t.Error("not all work items processed")
	} else if len(stats) > p*n {
		t.Error("too many work items processed")
	}
}

func TestManyTasks(t *testing.T) {
	p := 20
	w := 20
	n := 100
	low := 50
	high := 100


	stats := work_queues(p, w, n, low, high)

	if len(stats) < p*n {
		t.Error("not all work items processed")
	} else if len(stats) > p*n {
		t.Error("too many work items processed")
	}
}

func TestHighDurationRange(t *testing.T) {
	p := 20
	w := 20
	n := 20
	low := 10
	high := 1000


	stats := work_queues(p, w, n, low, high)

	if len(stats) < p*n {
		t.Error("not all work items processed")
	} else if len(stats) > p*n {
		t.Error("too many work items processed")
	}
}
