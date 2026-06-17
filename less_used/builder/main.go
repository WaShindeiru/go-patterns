// PROBLEM: How to create an object with many optional fields without using many constructors? And allow to set fields in any order?
// go run main.go 

package main

import "fmt"

// Struktura reprezentująca użytkownika.
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {

	// Bezpośrednie utworzenie obiektu.
	user := User{
		Name: "John",
		Age:  25,
	}

	// Dodanie pola później.
	user.Email = "john@example.com"

	// Wyświetlenie wyniku.
	fmt.Println(user)
}