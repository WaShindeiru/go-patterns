package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Number int
}

type Result struct {
	Number  int
	IsPrime bool
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- Result{Number: job.Number, IsPrime: isPrime(job.Number)}
	}
}

func main() {
	const numWorkers = 4
	const numJobs = 20

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup
	for i := range numWorkers {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= numJobs; i++ {
		jobs <- Job{Number: i}
	}
	close(jobs)

	primes := []int{}
	for r := range results {
		if r.IsPrime {
			primes = append(primes, r.Number)
		}
	}

	fmt.Printf("Primes in 1..%d: %v\n", numJobs, primes)
}
