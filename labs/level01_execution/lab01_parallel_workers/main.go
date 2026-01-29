package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"sync"
)

func main() {
	fmt.Println("Hello Lab01")

	n, _ := strconv.Atoi(os.Args[1])
	low, _ := strconv.Atoi(os.Args[2])
	high, _ := strconv.Atoi(os.Args[3])
	fmt.Printf("Received N=%d and range=[%d,%d]\n", n, low, high)

	parallel_workers(n, low, high)
}

func parallel_workers(n, low, high int) {
	var wg sync.WaitGroup

	// define durations upfront
	durations := make([]int, n)
	for i := range n {
		durations[i] = rand.Intn(high-low+1) + low
	}
	
	wg.Add(n)
	for i := range n {
		go block(i, durations[i], &wg)
	}

	wg.Wait()
}

func block(i, wait int, wg *sync.WaitGroup) {
	defer wg.Done()
	

	t0 := time.Now()
	time.Sleep(time.Duration(wait) * time.Millisecond)
	t1 := time.Now()
	fmt.Printf("Worker %d started at %s ended at %s and ran for %dns\n", i, t0, t1, t1.Sub(t0))
}
