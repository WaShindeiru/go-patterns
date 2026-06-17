package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Database struct {
	connection string
}

var (
	instance atomic.Pointer[Database]
	mu       sync.Mutex
)

// BEZPIECZNY: implementacja thread-safe Singleton
// używa double-check locking + atomic pointer
// gwarantuje tylko jedną instancję nawet w goroutines
func GetInstance() *Database {

	// szybka ścieżka (bez locka)
	if p := instance.Load(); p != nil {
		return p
	}

	// sekcja krytyczna
	mu.Lock()
	defer mu.Unlock()

	// ponowne sprawdzenie po wejściu do locka
	if instance.Load() == nil {
		instance.Store(&Database{connection: "db://localhost"})
	}

	return instance.Load()
}

func main() {

	var wg sync.WaitGroup
	results := make([]*Database, 100000)

	for i := range results {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			results[idx] = GetInstance()
		}(i)
	}

	wg.Wait()

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