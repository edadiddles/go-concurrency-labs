package main

import (
	"testing"
)


func TestBasic(t *testing.T) {
	n := 100
	w := 10
	low := 50
	high := 100

	actual := fan_in_out(n, w, low, high)
	
	expected := 5050
	if actual < expected {
		t.Error("Not all task's results were merged")
	} else if actual > expected {
		t.Error("Too many task's results were merged")
	}
}

func TestLongWorkRange(t *testing.T) {
	n := 10
	w := 5
	low := 10
	high := 3000

	actual := fan_in_out(n, w, low, high)
	
	expected := 55
	if actual < expected {
		t.Error("Not all task's results were merged")
	} else if actual > expected {
		t.Error("Too many task's results were merged")
	}
}
