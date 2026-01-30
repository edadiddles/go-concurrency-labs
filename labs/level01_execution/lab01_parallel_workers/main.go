package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"sync"
)

type Statistics struct {
	worker int
	start time.Time
	end time.Time
}

func main() {
	fmt.Println("Hello Lab01")

	n, _ := strconv.Atoi(os.Args[1])
	low, _ := strconv.Atoi(os.Args[2])
	high, _ := strconv.Atoi(os.Args[3])
	fmt.Printf("Received N=%d and range=[%d,%d]\n", n, low, high)

	parallel_workers(n, low, high)
}

func parallel_workers(n, low, high int) []Statistics {
	var wg sync.WaitGroup

	// define durations upfront
	durations := make([]int, n)
	for i := range n {
		durations[i] = rand.Intn(high-low+1) + low
	}

	// create channels
	c := make(chan Statistics)	
	for i := range n {
		wg.Add(1)
		go block(i, durations[i], c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var stats []Statistics
	for ch := range c {
		log(ch.worker, ch.start, ch.end)
		stats = append(stats, ch)
	}
	
	return stats
}

func block(i, wait int, c chan Statistics, wg *sync.WaitGroup) {
	defer wg.Done()

	t0 := time.Now()
	time.Sleep(time.Duration(wait) * time.Millisecond)
	t1 := time.Now()
	c <- Statistics{
		worker: i,
		start: t0,
		end: t1,
	};
}

func log(i int, t0, t1 time.Time) {
	fmt.Printf("Worker %d started at %s ended at %s and ran for %dns\n", i, t0, t1, t1.Sub(t0))
}
