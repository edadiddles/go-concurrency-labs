package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"math/rand"
)

type AccessRecord struct {
	id int
	val int
	access_type string
	access_time time.Time
}

type RWReport struct {
	num_reads int
	num_writes int
	access_records []AccessRecord
}

type SharedResource struct {
	mu *sync.RWMutex
	val int
}

func main() {
	fmt.Println("Hello lab10")

	r, _ := strconv.Atoi(os.Args[1])
	w, _ := strconv.Atoi(os.Args[2])
	k, _ := strconv.Atoi(os.Args[3])
	low, _ := strconv.Atoi(os.Args[4])
	high, _ := strconv.Atoi(os.Args[5])

	fmt.Println("Readers:", r)
	fmt.Println("Writers:", w)
	fmt.Println("Repeats per Reader/Writer:", k)
	fmt.Printf("Workload Range: [%d,%d]\n", low, high)

	rw_report := rw_lock(r, w, k, low, high)

	fmt.Println(rw_report)
}

func rw_lock(r, w, k, low, high int) RWReport {

	var mu sync.RWMutex
	resource := SharedResource{
		mu: &mu,
		val: 0,
	}

	var wg sync.WaitGroup
	ch := make(chan AccessRecord)
	for id := range w {
		d := rand.Intn(high-low+1) + low
		wg.Add(1)
		go write(id, k, d, &resource, ch, &wg)
	}

	for id := range r {
		d := rand.Intn(high-low+1) + low
		wg.Add(1)
		go read(id, k, d, &resource, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	num_reads := 0
	num_writes := 0
	access_records := make([]AccessRecord, 0)
	for rec := range ch {
		if rec.access_type == "read" {
			num_reads += 1
		} else {
			num_writes += 1
		}
		access_records = append(access_records, rec)
	}
	return RWReport{
		num_reads: num_reads,
		num_writes: num_writes,
		access_records: access_records,
	}
}

func write(id, k, d int, r *SharedResource, ch chan AccessRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range k {
		r.mu.Lock()
		v := (id+i+1)*d
		r.val = v
		t := time.Now()
		time.Sleep(time.Duration(d)*time.Millisecond)
		r.mu.Unlock()
		ch <- AccessRecord{
			id: id,
			val: v,
			access_type: "write",
			access_time: t,
		}
	}
}

func read(id, k, d int, r *SharedResource, ch chan AccessRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	for range k {
		r.mu.RLock()
		time.Sleep(time.Duration(d)*time.Millisecond)
		v := r.val
		t := time.Now()
		r.mu.RUnlock()
		ch <- AccessRecord{
			id: id,
			val: v,
			access_type: "read",
			access_time: t,
		}
	}
}
