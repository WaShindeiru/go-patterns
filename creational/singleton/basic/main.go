package main

import (
	"fmt"
	"sync"
)

type Database struct {
	connection string
}

// GLOBALNA INSTANCJA SINGLETONA
// to jest miejsce gdzie “trzymamy” jedyny obiekt
var instance *Database

// GETTER SINGLETONA (LAZY INITIALIZATION)
// to jest główna część wzorca Singleton:
// - sprawdza czy instancja istnieje
// - jeśli nie → tworzy ją
// - jeśli tak → zwraca istniejącą
func GetInstance() *Database {

	// WARUNEK SINGLETONA
	// jeśli nie istnieje jeszcze instancja → tworzymy ją
	if instance == nil {
		instance = &Database{connection: "db://localhost"}
	}

	// zawsze zwracamy tę samą instancję
	return instance
}

func main() {

	// UŻYCIE SINGLETONA
	// wszystkie wywołania powinny zwracać tę samą instancję
	db1 := GetInstance()
	db2 := GetInstance()

	fmt.Println("--- Single-threaded ---")
	fmt.Printf("db1 address: %p\n", db1)
	fmt.Printf("db2 address: %p\n", db2)
	fmt.Printf("db1 == db2: %v\n", db1 == db2)

	// RESET (tylko do pokazania problemu concurrency)
	instance = nil

	var wg sync.WaitGroup
	results := make([]*Database, 100000)

	fmt.Println("\n--- Multi-threaded (problem concurrency) ---")

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