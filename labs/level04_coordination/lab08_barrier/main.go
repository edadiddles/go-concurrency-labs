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

	var mu sync.Mutex
	counter := BarrierCounter{
		cond: sync.NewCond(&mu),
		cnt: 0,
		chk_val: n,
	}
	ch := make(chan WorkerReport)
	for id := range n {
		go worker(id, p, durations[p*id:p*(id+1)], &counter, ch)
	}

	worker_reports := make([]WorkerReport, n)
	cnt := 0
	for r := range ch {
		worker_reports[cnt] = r
		cnt += 1
		if cnt == n {
			close(ch)
		}
	}

	return  worker_reports
}

func worker(id, num_phases int, durations []int, counter *BarrierCounter, ch chan WorkerReport) {
	worker_report := WorkerReport{
		id: id,
		barriers: make([]BarrierReport, 0),
	}
	for phase_id := range num_phases {
		arrival_time := time.Now()
		time.Sleep(time.Duration(durations[phase_id])*time.Millisecond)	

		counter.cond.L.Lock()
		counter.cnt += 1
		for counter.cnt < counter.chk_val {
			counter.cond.Wait()
		}
		counter.cond.Broadcast()
		counter.cond.L.Unlock()

		report := BarrierReport{
			id: phase_id,
			arrival_time: arrival_time,
			departure_time: time.Now(),
		}

		 worker_report.barriers = append(worker_report.barriers, report)
	}

	ch <- worker_report	
}
