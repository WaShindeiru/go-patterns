package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Database struct {
	connection string
}

// ATOMIC POINTER + MUTEX
// to są mechanizmy zapewniające bezpieczeństwo w concurrency
var (
	instance atomic.Pointer[Database] // przechowuje singleton
	mu       sync.Mutex               // chroni
)

// BEZPIECZNA IMPLEMENTACJA SINGLETONA
// gwarantuje: tylko jedna instancja w całym programie
func GetInstance() *Database {

	// SZYBKA ŚCIEŻKA (lock-free check)
	// sprawdzamy czy instancja już istnieje
	if p := instance.Load(); p != nil {
		return p
	}

	// tylko jedna goroutine może wejść tutaj naraz
	mu.Lock()
	defer mu.Unlock()

	// DOUBLE CHECK
	// po wejściu do locka sprawdzamy ponownie
	if instance.Load() == nil {
		// tworzymy singleton tylko raz
		instance.Store(&Database{connection: "db://localhost"})
	}

	// zwracamy zawsze tę samą instancję
	return instance.Load()
}

func main() {

	var wg sync.WaitGroup
	results := make([]*Database, 100000)

	// TEST CONCURRENCY
	// wiele goroutines próbuje dostać singleton jednocześnie
	for i := range results {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			results[idx] = GetInstance()
		}(i)
	}

	wg.Wait()

	// sprawdzamy czy wszystkie referencje są identyczne
	first := results[0]
	allSame := true

	for _, inst := range results {
		if inst != first {
			allSame = false
			break
		}
	}

	fmt.Printf("Instance address: %p\n", first)
	fmt.Printf("All goroutines got the same instance: %v\n", allSame)
}