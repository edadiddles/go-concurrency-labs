package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"math"
	"math/rand"
	"time"
)

type Statistics struct {
	id int
	start time.Time
	end time.Time
	duration int64
}

func main() {
	fmt.Println("hello lab02")

	n, _ := strconv.Atoi(os.Args[1])
	mode := os.Args[2]

	fmt.Printf("Workers: %d\nMode: %s\n\n", n, mode)

	stats := saturated_execution(n, mode)

	log_stats(stats)
}

func saturated_execution(n int, mode string) []Statistics {
	worker_modes := make([]string, n)
	fmt.Println("starting saturated execution")
	for i := range n {
		m := mode
		if m == "mixed" {
			m = "blocking"
			if rand.Intn(2) == 0 {
				m = "cpu"
			}
		}
			
		switch m {
		case "blocking":
			worker_modes[i] = m
		case "cpu":
			worker_modes[i] = m
		default:
			panic("mode is not known should be: cpu, blocking, mixed")
		}
	}

	ch := make(chan Statistics)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go perform_work(i, worker_modes[i], &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	stats := make([]Statistics, 0)
	for c := range ch {
		stats = append(stats, c)
	}

	return stats
}

func perform_work(id int, mode string, wg *sync.WaitGroup, ch chan Statistics) {
	defer wg.Done()

	start := time.Now()
	switch mode {
	case "blocking":
		blocking_op()
	case "cpu":
		cpu_op()
	}
	end := time.Now()

	ch <- Statistics{
		id: id,
		start: start,
		end: end,
		duration: end.Sub(start).Milliseconds(),
	}
}

func blocking_op() {
	time.Sleep(time.Duration(100) * time.Millisecond)
}

func cpu_op() {
	start := time.Now()
	for time.Since(start).Milliseconds() < int64(100) {
		// perform some non-trival computation
		a := float64(time.Now().UnixMicro()) / math.Pi
		b := math.Sqrt(a) + math.Phi
		_ = math.Atan(b)
	}
}

func log_stats(stats []Statistics) {
	for _, s := range stats {
		fmt.Printf(
			"Worker %d: (%s, %s, %d)\n",
			s.id, s.start.Format("05.000"), s.end.Format("05.000"), s.duration,
		)
	}
}
