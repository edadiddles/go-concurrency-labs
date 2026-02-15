package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	id         int
	duration   int
	timeout    int
	queue_time time.Time
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
			queue <- Task{
				id:         task_id,
				duration:   rand.Intn(high-low-+1) + low,
				timeout:    p,
				queue_time: time.Now(),
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
	discarded := 0
	timeout_ch := make(chan int)
	for t := range queue {
		submitted += 1
		if  time.Since(t.queue_time) > time.Duration(t.timeout)*time.Millisecond {
			discarded += 1
			continue
		}
		go timeout(t, timeout_ch)
	loop:
		for {
			select {
			case to := <-timeout_ch:
				if to == t.id {
					timed_out += 1
					break loop
				}
			case <-time.After(time.Duration(t.duration) * time.Millisecond):
				completed += 1
				break loop
			}
		}
	}

	close(timeout_ch)
	ch <- TimeoutReport{
		submitted: submitted,
		completed: completed,
		timed_out: timed_out,
		discarded: discarded,
	}
}

func timeout(task Task, ch chan int) {
	loop:
	for {
		select {
		case <- ch:
			break loop
		case <- time.After(time.Duration(task.timeout)*time.Millisecond):
			ch <- task.id
			break loop
		}
	}
}
