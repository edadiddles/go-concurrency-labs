package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
)

type PipelineRecord struct {
	id int
	stages []int
}

func main() {
	fmt.Println("Hello Lab07")

	n, _ := strconv.Atoi(os.Args[1])
	s, _ := strconv.Atoi(os.Args[2])
	low, _ := strconv.Atoi(os.Args[3])
	high, _ := strconv.Atoi(os.Args[4])

	fmt.Println("Number of items:", n)
	fmt.Println("Number of stages:", s)
	fmt.Printf("Range of stage duration: [%d,%d]\n", low, high)

	records := pipeline(n, s, low, high)

	fmt.Println("Output Records")
	for _, rec := range records {
		fmt.Println(rec)
	}
}

func pipeline(num_items, num_stages, low, high int) []PipelineRecord {
	channels := make([]chan PipelineRecord, num_stages+1)
	for i := range num_stages+1 {
		channels[i] = make(chan PipelineRecord)
	}

	for i := range num_stages {
		d := rand.Intn(high-low+1)+low
		go stage(i, d, channels[i], channels[i+1])
	}

	go func() {
		for id := range num_items {
			rec := PipelineRecord{
				id: id,
				stages: make([]int, 0),
			}

			channels[0] <- rec
		}

		close(channels[0])
	}()	
	records := make([]PipelineRecord, 0)
	for rec := range channels[len(channels)-1] {
		records = append(records, rec)
	}

	return records
}

func stage(n, d int, in_ch <- chan PipelineRecord, out_ch chan <- PipelineRecord) {
	for rec := range in_ch {
		rec.stages = append(rec.stages, n)
		time.Sleep(time.Duration(d)*time.Millisecond)
		out_ch <- rec
	}
	close(out_ch)
}
