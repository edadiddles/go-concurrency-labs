package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand"
	"sync"
)

type SharedResource struct {
	initialized bool
	initialization_time time.Time
	val int
	cond sync.Cond
}

type WorkerAccessRecord struct {
	id int
	performed_initialization bool
	access_time time.Time
}

type AccessReport struct {
	initialized bool
	access_records []WorkerAccessRecord
}

func main() {
	fmt.Println("Hello lab09")

	n, _ := strconv.Atoi(os.Args[1])
	low, _ := strconv.Atoi(os.Args[2])
	high, _ := strconv.Atoi(os.Args[3])

	fmt.Println("Number of workers:", n)
	fmt.Printf("Workload range: [%d,%d]\n", low, high)

	report := run_once(n, low, high)

	fmt.Println("Initialized:", report.initialized)
	for _, rec := range report.access_records {
		fmt.Println(rec)
	}
}

func run_once(n, low, high int) AccessReport {
	durations := make([]int, n)
	for k := range n {
		durations[k] = rand.Intn(high-low+1) + low
	}

	var wg sync.WaitGroup
	ch := make(chan WorkerAccessRecord)

	var mu sync.Mutex
	r := SharedResource{
		initialized: false,
		cond: *sync.NewCond(&mu),
	}
	for id := range n {
		wg.Add(1)
		go worker(id, durations[id], &r, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	is_initialized := false
	access_records := make([]WorkerAccessRecord, 0)
	for rec := range ch {
		if rec.performed_initialization {
			is_initialized = true
		}
		access_records = append(access_records, rec)
	}

	return AccessReport{
		initialized: is_initialized,
		access_records: access_records,
	}
}


func worker(id, d int, r *SharedResource, ch chan WorkerAccessRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	performed_initialized := false

	r.cond.L.Lock()
	if !r.initialized {
		initialize(r, d)
		performed_initialized = true
		r.cond.Broadcast()
	}
	for !r.initialized {
		r.cond.Wait()
	}
	access_time := access(r, id, d)
	r.cond.L.Unlock()

	
	ch <- WorkerAccessRecord{
		id: id,
		performed_initialization: performed_initialized,
		access_time: access_time,
	}
}

func initialize(r *SharedResource, d int) {
	time.Sleep(time.Duration(d)*time.Millisecond)
	r.initialized = true
	r.initialization_time = time.Now()
}

func access(r *SharedResource, id, d int) time.Time {
	r.val = id
	time.Sleep(time.Duration(d)*time.Millisecond)
	return time.Now()
}
