package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("hello lab03")

	n, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])

	count := shared_counter(n, k)

	fmt.Println("current count:", count)
	fmt.Println("expected count:", n*k)
}


func shared_counter(n, k int) int {
	var wg sync.WaitGroup
	counter := 0
	for range n {
		wg.Add(1)
		go increment_counter(&counter, k, &wg)
	}
	wg.Wait()
	return counter
}

func increment_counter(cnt *int, k int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range k {
		*cnt += 1
	}
}
