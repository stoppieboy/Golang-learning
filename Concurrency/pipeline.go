package main

import (
	"fmt"
	"time"
)

// sliceToChannel function returns a read-only channel of integer type
func sliceToChannel(nums []int) <-chan int{
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}

func main() {
	p := fmt.Printf
	// input
	nums := []int{2,3,4,7,1}
	start := time.Now()
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
	// time.Sleep(time.Nanosecond *4000)
	end := time.Now()
	diff := end.Sub(start)
	p("Time taken: %10.9f",diff.Seconds())
}