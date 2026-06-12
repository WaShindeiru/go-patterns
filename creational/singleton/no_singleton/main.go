package main

import "fmt"

type Database struct {
	connection string
}

func NewDatabase() *Database {
	return &Database{connection: "db://localhost"}
}

func main() {
	db1 := NewDatabase()
	db2 := NewDatabase()
	db3 := NewDatabase()

	fmt.Printf("db1 address: %p\n", db1)
	fmt.Printf("db2 address: %p\n", db2)
	fmt.Printf("db3 address: %p\n", db3)
	fmt.Printf("db1 == db2: %v\n", db1 == db2)
	fmt.Printf("db2 == db3: %v\n", db2 == db3)
}
