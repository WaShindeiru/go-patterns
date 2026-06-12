package main

import (
	"fmt"
	"sync"
)

// Two different pipeline branches running in parallel:
//
//   numbers 1-5  ──> [square workers x2] ──┐
//                                            ├──> merge ──> results
//   numbers 6-10 ──> [cube workers x2]   ──┘

type Result struct {
	Input  int
	Output int
	Op     string
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func merge[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	forward := func(ch <-chan T) {
		defer wg.Done()
		for v := range ch {
			out <- v
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go forward(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func worker(id int, op string, fn func(int) int, in <-chan int) <-chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		for n := range in {
			result := fn(n)
			fmt.Printf("  worker %d [%s]: %d -> %d\n", id, op, n, result)
			out <- Result{Input: n, Output: result, Op: op}
		}
	}()
	return out
}

func fanOut(n int, op string, fn func(int) int, in <-chan int) <-chan Result {
	workers := make([]<-chan Result, n)
	for i := range n {
		workers[i] = worker(i, op, fn, in)
	}
	return merge(workers...)
}

func main() {
	smallNums := generate(1, 2, 3, 4, 5)
	largeNums := generate(6, 7, 8, 9, 10)
	
	squares := fanOut(2, "square", func(n int) int { return n * n }, smallNums)
	cubes := fanOut(2, "cube", func(n int) int { return n * n * n }, largeNums)

	for r := range merge(squares, cubes) {
		fmt.Printf("  %s(%d) = %d\n", r.Op, r.Input, r.Output)
	}
}
