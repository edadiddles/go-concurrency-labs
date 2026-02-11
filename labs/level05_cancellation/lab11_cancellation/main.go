package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"sync"
	"math/rand"
)

type CancelTrigger struct {
	mu *sync.RWMutex
	canceled bool
	cancel_time time.Time
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

	var wg sync.WaitGroup
	var mu sync.RWMutex
	work_start := time.Now()
	trigger := CancelTrigger{
		mu: &mu,
		canceled: false,
	}
	ch := make(chan WorkerReport)
	for id := range n {
		wg.Add(1)
		go worker(id, k, low, high, &trigger, ch, &wg)
	}

	wg.Add(1)
	go check_timeout(t, work_start, &trigger, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	worker_reports := make([]WorkerReport, 0)
	for wr := range ch {
		worker_reports = append(worker_reports, wr)
	}

	trigger.mu.RLock()
	defer trigger.mu.RUnlock()
	return CancelReport{
		cancel_time: trigger.cancel_time,
		worker_reports: worker_reports,
	}
}

func worker(id, k, low, high int, trigger *CancelTrigger, ch chan WorkerReport, wg *sync.WaitGroup) {
	defer wg.Done()
	canceled := false
	for range k {
		d := rand.Intn(high-low+1)+low
		time.Sleep(time.Duration(d)*time.Millisecond)
		
		trigger.mu.RLock()
		if trigger.canceled {
			canceled = true
			trigger.mu.RUnlock()
			break
		}
		trigger.mu.RUnlock()
	}

	ch <- WorkerReport{
		id: id,
		interrupted: canceled,
		exit_time: time.Now(),
	}
}

func check_timeout(timeout int, work_start time.Time, trigger *CancelTrigger, wg *sync.WaitGroup) {
	defer wg.Done()
	for time.Since(work_start) < time.Duration(timeout)*time.Millisecond {
		time.Sleep(time.Duration(100)*time.Millisecond)
	}

	trigger.mu.Lock()
	trigger.canceled = true
	trigger.cancel_time = time.Now()
	trigger.mu.Unlock()
}

