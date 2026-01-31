package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("hello lab03")

	n, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])

	count := shared_counter(n, k)

	fmt.Println("current count:", count)
}


func shared_counter(n, k int) int {
	counter := 0
	for range n {
		go increment_counter(&counter, k)
	}
	return counter
}

func increment_counter(cnt *int, k int) {
	for range k {
		*cnt += 1
	}
}
