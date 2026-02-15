package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	id       int
	duration int
	ctx      context.Context
	cancel   context.CancelFunc
}

type TimeoutReport struct {
	submitted int
	completed int
	timed_out int
	discarded int
}

func main() {
	fmt.Println("Hello lab12")

	n, _ := strconv.Atoi(os.Args[1])
	c, _ := strconv.Atoi(os.Args[2])
	p, _ := strconv.Atoi(os.Args[3])
	low, _ := strconv.Atoi(os.Args[4])
	high, _ := strconv.Atoi(os.Args[5])

	fmt.Println("Number of Workers:", n)
	fmt.Println("Queue Capacity:", c)
	fmt.Println("Task Timeout:", p)
	fmt.Printf("Workload Range: [%d,%d]\n", low, high)

	report := timeout_queue(n, c, p, low, high)
	fmt.Println(report)
}

func timeout_queue(n, c, p, low, high int) TimeoutReport {

	var wg sync.WaitGroup
	queue := make(chan Task, c)

	ch := make(chan TimeoutReport)
	for id := range n {
		wg.Add(1)
		go worker(id, queue, ch, &wg)
	}

	go func() {
		for task_id := range 1000 {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p)*time.Millisecond)
			task := Task{
				id:       task_id,
				duration: rand.Intn(high-low-+1) + low,
				ctx:      ctx,
				cancel:   cancel,
			}

			select {
			case queue <- task:
				//accepted
			case <-ctx.Done():
				cancel()
				ch <- TimeoutReport{submitted: 1, discarded: 1}
				continue
			}

			time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
		}
		close(queue)
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	submitted := 0
	completed := 0
	timed_out := 0
	discarded := 0

	for rep := range ch {
		submitted += rep.submitted
		completed += rep.completed
		timed_out += rep.timed_out
		discarded += rep.discarded
	}

	return TimeoutReport{
		submitted: submitted,
		completed: completed,
		timed_out: timed_out,
		discarded: discarded,
	}
}

func worker(id int, queue chan Task, ch chan TimeoutReport, wg *sync.WaitGroup) {
	defer wg.Done()
	_ = id
	submitted := 0
	completed := 0
	timed_out := 0
	for t := range queue {
		submitted += 1
		select {
		case <-t.ctx.Done():
			timed_out += 1
			t.cancel()
		case <-time.After(time.Duration(t.duration) * time.Millisecond):
			completed += 1
			t.cancel()
		}
	}

	ch <- TimeoutReport{
		submitted: submitted,
		completed: completed,
		timed_out: timed_out,
	}
}
