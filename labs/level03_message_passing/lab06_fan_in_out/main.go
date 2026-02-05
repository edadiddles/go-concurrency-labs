package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
	"sync"
)

type WorkItem struct {
	id int
	task func() int
}

type Result struct {
	work_item_id int
	val int
}

func main() {
	fmt.Println("Hello Lab06")

	n, _ := strconv.Atoi(os.Args[1])
	w, _ := strconv.Atoi(os.Args[2])
	low, _ := strconv.Atoi(os.Args[3])
	high, _ := strconv.Atoi(os.Args[4])

	fmt.Println("Number of work items:", n)
	fmt.Println("Number of workers:", w)
	fmt.Printf("Workload Range: [%d,%d]\n", low, high)

	res := fan_in_out(n, w, low, high)

	fmt.Println("Sum of work tasks:", res)

}

func fan_in_out(n, w, low, high int) int {
	var worker_wg sync.WaitGroup
	work_channel := make(chan WorkItem, n)
	result_channel := make(chan Result, n)

	// fan in
	for range w {
		worker_wg.Add(1)
		go func(in_ch <- chan WorkItem, out_ch chan <- Result, wg *sync.WaitGroup) {
			defer worker_wg.Done()
			for work_item := range in_ch {
				val := work_item.task()
				res := Result{
					work_item_id: work_item.id,
					val: val,
				}
				out_ch <- res
			}
		}(work_channel, result_channel, &worker_wg)
	}

	// fan out
	for k := range n {
		d := rand.Intn(high-low+1) + low
		task := func() int {
			time.Sleep(time.Duration(d) * time.Millisecond)
			return k+1
		}
		work_channel <- WorkItem{
			id: k,
			task: task,
		}
	}

	close(work_channel)
	worker_wg.Wait()
	close(result_channel)

	sum := 0
	for res := range result_channel {
		sum += res.val
	}

	return sum
}
