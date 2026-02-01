package main

import (
	"testing"
	"math/rand"
)

func TestCounter(t *testing.T) {
	for range 100 {
		n := rand.Intn(500) + 500
		k := rand.Intn(100) + 100

		actual := shared_counter(n, k)

		expected := n*k
		if actual != expected {
			t.Error("counter corruption expected", expected, "actual", actual)
		}
	}
}
