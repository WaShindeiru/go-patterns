package main

import (
	"fmt"
	"sync"
)

// PROBLEM: How to create a single shared instance safely in concurrent environment?

type Database struct {
	connection string
}

var instance *Database

// NIEBEZPIECZNY: brak synchronizacji (race condition w goroutines)
// Ten kod działa poprawnie tylko w single-threaded execution
// W concurrency może stworzyć więcej niż jedną instancję
func GetInstance() *Database {

	if instance == nil {
		instance = &Database{connection: "db://localhost"}
	}

	return instance
}

func main() {

	// Single-threaded: działa poprawnie
	db1 := GetInstance()
	db2 := GetInstance()

	fmt.Println("--- Single-threaded ---")
	fmt.Printf("db1 address: %p\n", db1)
	fmt.Printf("db2 address: %p\n", db2)
	fmt.Printf("db1 == db2: %v\n", db1 == db2)

	// reset dla testu concurrency
	instance = nil

	var wg sync.WaitGroup
	results := make([]*Database, 100000)

	fmt.Println("\n--- Multi-threaded (data race possible) ---")

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

	fmt.Printf("All goroutines got same instance: %v\n", allSame)
}