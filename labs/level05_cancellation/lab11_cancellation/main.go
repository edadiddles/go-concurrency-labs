package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"sync"
	"math/rand"
)

type CancelSignal struct {
	done chan struct{}
	cancel_time time.Time
	once *sync.Once
}

type WorkerReport struct {
	id int
	interrupted bool
	exit_time time.Time
}

type CancelReport struct {
	cancel_time time.Time
	worker_reports []WorkerReport
}

func main() {
	fmt.Println("Hello lab11")

	n, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	t, _ := strconv.Atoi(os.Args[3])
	low, _ := strconv.Atoi(os.Args[4])
	high, _ := strconv.Atoi(os.Args[5])

	fmt.Println("Number of workers:", n)
	fmt.Println("Number of worker units per worker:", k)
	fmt.Println("Cancellation trigger:", t)
	fmt.Printf("Workload Range: [%d,%d]\n", low, high)

	report := cancellation(n, k, t, low, high)

	fmt.Println(report)
}

func cancellation(n, k, t, low, high int) CancelReport {
	
	ch := make(chan WorkerReport)
	done := make(chan struct{})

	var wg sync.WaitGroup
	var once sync.Once
	cancel_signal := CancelSignal{
		done: done,
		once: &once,
	}
	for id := range n {
		wg.Add(1)
		go worker(id, k, low, high, ch, &cancel_signal, &wg)
	}

	go check_timeout(t, &cancel_signal)

	go func(cancel_signal *CancelSignal) {
		wg.Wait()
		close(ch)
		cancel_signal.once.Do(func() {
			// this is not a cancellation. this is clean up
			close(cancel_signal.done)
		})
	}(&cancel_signal)

	worker_reports := make([]WorkerReport, 0)
	for wr := range ch {
		worker_reports = append(worker_reports, wr)
	}

	return CancelReport{
		cancel_time: cancel_signal.cancel_time,
		worker_reports: worker_reports,
	}
}

func worker(id, k, low, high int, ch chan WorkerReport, cancel_signal *CancelSignal, wg *sync.WaitGroup) {
	defer wg.Done()
	interrupted := false
	loop:
		for range k {
			d := rand.Intn(high-low+1)+low
			select {
			case <-cancel_signal.done:
				interrupted = true
				break loop
			case <-time.After(time.Duration(d)*time.Millisecond):
			}
		}

	ch <- WorkerReport{
		id: id,
		interrupted: interrupted,
		exit_time: time.Now(),
	}
}

func check_timeout(timeout int, cancel_signal *CancelSignal) {
	time.Sleep(time.Duration(timeout)*time.Millisecond)
	cancel_signal.once.Do(func() {
		cancel_signal.cancel_time = time.Now()
		close(cancel_signal.done)
	})
}
