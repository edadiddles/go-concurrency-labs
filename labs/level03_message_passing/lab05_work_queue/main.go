package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"math/rand"
	"time"
)

type Statistics struct {
	id int
}

func main() {
	fmt.Println("Hello lab05")

	p, _ := strconv.Atoi(os.Args[1])
	w, _ := strconv.Atoi(os.Args[2])
	n, _ := strconv.Atoi(os.Args[3])
	low, _ := strconv.Atoi(os.Args[4])
	high, _ := strconv.Atoi(os.Args[5])

	work_queues(p, w, n, low, high)
}

func work_queues(p, w, n, low, high int) []Statistics {
	var wg_producers sync.WaitGroup
	var wg_consumers sync.WaitGroup

	durations := make([]int, p*n)
	for i := range p*n {
		durations[i] = rand.Intn(high-low+1) + low
	}

	work_queue := make(chan func() int)

	stats_queue := make(chan Statistics, n*p)
	for range w {
		wg_consumers.Add(1)
		go worker(work_queue, stats_queue, &wg_consumers)
	}


	for id := range p {
		wg_producers.Add(1)
		go producer(id, n, durations[id*n:(id+1)*n], work_queue, &wg_producers)
	}


	go func() {
		wg_producers.Wait()
		close(work_queue)
		wg_consumers.Wait()
		close(stats_queue)
	}()
	stats := make([]Statistics, 0)
	for stat := range stats_queue {
		stats = append(stats, stat)
	}

	return stats
}

func producer(id, n int, durations []int, work_queue chan<-func() int, wg *sync.WaitGroup) {
	defer wg.Done()

	for k := range n {
		task_id := (id+1)*(k+1) + k
		work_queue <- func() int {
			time.Sleep(time.Duration(durations[k])*time.Millisecond)
			return task_id
		}
	}

}

func worker(work_queue <-chan func() int, stats_queue chan Statistics, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range work_queue {
		val := task()
		stats_queue <- Statistics {
			id: val,
		}
	}
}
