package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"sync"
	"math/rand"
)

type BarrierCounter struct {
	cond *sync.Cond
	generation int
	cnt int
	chk_val int
}

type WorkerReport struct {
	id int
	barriers []BarrierReport
}

type BarrierReport struct {
	id int
	arrival_time time.Time
	departure_time time.Time
}

func main() {
	fmt.Println("Hello Lab08")

	n, _ := strconv.Atoi(os.Args[1])
	p, _ := strconv.Atoi(os.Args[2])
	low, _ := strconv.Atoi(os.Args[3])
	high, _ := strconv.Atoi(os.Args[4])

	fmt.Println("Number of workers:", n)
	fmt.Println("Number of phases:", p)
	fmt.Printf("Work load time range: [%d,%d]\n", low, high)

	reports := barrier(n, p, low, high)

	for _, report := range reports {
		fmt.Println(report)
	}
}

func barrier(n, p, low, high int) []WorkerReport {
	durations := make([]int, n*p)
	for k := range n*p {
		durations[k] = rand.Intn(high-low+1) + low
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := BarrierCounter{
		cond: sync.NewCond(&mu),
		cnt: 0,
		chk_val: n,
	}
	ch := make(chan WorkerReport)
	for id := range n {
		wg.Add(1)
		go worker(id, p, durations[p*id:p*(id+1)], &counter, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	worker_reports := make([]WorkerReport, 0)
	for r := range ch {
		worker_reports = append(worker_reports, r)
	}

	return  worker_reports
}

func worker(id, p int, durations []int, b *BarrierCounter, wg *sync.WaitGroup, ch chan WorkerReport) {
	defer wg.Done()

	worker_report := WorkerReport{
		id: id,
		barriers: make([]BarrierReport, 0),
	}
	for phase_id := range p {
		time.Sleep(time.Duration(durations[phase_id])*time.Millisecond)	
		arrival_time := time.Now()

		b.cond.L.Lock()
		generation := b.generation
		b.cnt += 1
		if b.cnt == b.chk_val {
			b.cnt = 0
			b.generation += 1
			b.cond.Broadcast()
		}
		for generation == b.generation {
			b.cond.Wait()
		}	
		b.cond.L.Unlock()

		report := BarrierReport{
			id: phase_id,
			arrival_time: arrival_time,
			departure_time: time.Now(),
		}

		 worker_report.barriers = append(worker_report.barriers, report)
	}

	ch <- worker_report	
}
