package main

import "fmt"

// PROBLEM: How to ensure only one instance of a resource exists?

type Database struct {
	connection string
}

// Tworzenie nowego obiektu za każdym razem
func NewDatabase() *Database {
	return &Database{connection: "db://localhost"}
}

func main() {

	// Każde wywołanie tworzy nową instancję
	db1 := NewDatabase()
	db2 := NewDatabase()
	db3 := NewDatabase()

	// Pokazujemy różne adresy w pamięci
	fmt.Printf("db1 address: %p\n", db1)
	fmt.Printf("db2 address: %p\n", db2)
	fmt.Printf("db3 address: %p\n", db3)

	// Porównanie wskaźników
	fmt.Printf("db1 == db2: %v\n", db1 == db2)
	fmt.Printf("db2 == db3: %v\n", db2 == db3)
}