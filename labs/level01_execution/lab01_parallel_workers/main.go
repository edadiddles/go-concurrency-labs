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
	
	wg.Add(n)
	for i := range n {
		r := rand.Intn(high-low+1) + low
		go block(i, r, &wg)
	}

	wg.Wait()
}

func block(i, r int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d initialized blocking for %dms\n", i, r)
	time.Sleep(time.Duration(r) * time.Millisecond)
	fmt.Printf("Worker %d finished\n", i)
}
