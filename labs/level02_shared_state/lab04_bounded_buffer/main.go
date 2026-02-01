package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

type BoundedBuffer struct {
	mux      *sync.Mutex
	items    []int
	length   int
	capacity int
}

type Record struct {
	t time.Time
	val int
	op string
}

func main() {
	fmt.Println("Hello lab04")

	p, _ := strconv.Atoi(os.Args[1])
	c, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])
	n, _ := strconv.Atoi(os.Args[4])

	fmt.Println("Producers:", p)
	fmt.Println("Consumers:", c)
	fmt.Println("Capacity:", b)
	fmt.Println("Num Items per Producer:", n)

	records := bounded_buffer(p, c, b, n)

	for _, rec := range records {
		fmt.Printf("(%s, %d)\n", rec.op, rec.val)
	}
}

func bounded_buffer(p, c, b, n int) []Record {
	var wg_producer sync.WaitGroup
	var wg_consumer sync.WaitGroup

	buffer := BoundedBuffer{
		mux:      &sync.Mutex{},
		items:    make([]int, b),
		length:   0,
		capacity: b,
	}


	ch := make(chan Record, p*n)

	for i := range p {
		wg_producer.Add(1)
		go producer(i, n, &buffer, ch, &wg_producer)
	}

	for range c {
		wg_consumer.Add(1)
		go consumer(&buffer, ch, &wg_consumer)
	}

	go func() {
		wg_producer.Wait()
		wg_consumer.Wait()
		close(ch)
	}()


	records := make([]Record, 2*n*p)
	cnt := 0
	for c := range ch {
		records[cnt] = c
		cnt += 1
	}

	sort(&records)

	return records
}

func producer(id, n int, buffer *BoundedBuffer, ch chan Record, wg *sync.WaitGroup) {
	defer wg.Done()

	cnt := 0
	for cnt < n {
		buffer.mux.Lock()
		if buffer.length < buffer.capacity {
			val := int(math.Pow10(id+6)) + cnt
			buffer.items[buffer.length] = val
			buffer.length += 1
			cnt += 1
			ch <- Record {
				t: time.Now(),
				val: val,
				op: "+",
			}
		}
		buffer.mux.Unlock()
	}
}

func consumer(buffer *BoundedBuffer, ch chan Record, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("consumer online")

	last_update := time.Now()
	for time.Since(last_update).Milliseconds() < int64(1000) {
		buffer.mux.Lock()
		if buffer.length > 0 {
			val := buffer.items[buffer.length-1]
			buffer.length -= 1
			last_update = time.Now()
			ch <- Record{
				t: last_update,
				val: val,
				op: "-",
			}
		}
		buffer.mux.Unlock()
	}
}

func sort(records *[]Record) {
	for j := 1; j < len(*records); j++ {
		for i := 0; i < j; i++ {
			if (*records)[i].t.Compare((*records)[j].t) > 0 {
				temp := (*records)[i]
				(*records)[i] = (*records)[j]
				(*records)[j] = temp
			}
		}
	}
}
