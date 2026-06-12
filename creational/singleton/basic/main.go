package main

import (
	"fmt"
	"sync"
)

type Database struct {
	connection string
}

var instance *Database

func GetInstance() *Database {
	if instance == nil {
		instance = &Database{connection: "db://localhost"}
	}
	return instance
}

func main() {
	// Single-threaded: works correctly
	db1 := GetInstance()
	db2 := GetInstance()
	fmt.Println("--- Single-threaded ---")
	fmt.Printf("db1 address: %p\n", db1)
	fmt.Printf("db2 address: %p\n", db2)
	fmt.Printf("db1 == db2: %v\n", db1 == db2)

	// Multi-threaded: data race — run with: go run -race .
	instance = nil
	var wg sync.WaitGroup
	results := make([]*Database, 100000)

	fmt.Println("\n--- Multi-threaded (run with -race to see the data race) ---")
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
	fmt.Printf("All goroutines got the same instance: %v\n", allSame)
}
