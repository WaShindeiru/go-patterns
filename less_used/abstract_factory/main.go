package main

import "fmt"

// PROBLEM: Jak tworzyć rodziny powiązanych obiektów bez rozbudowanej hierarchii fabryk?

type Button interface {
	Render()
}

// Windows implementacja
type WindowsButton struct{}

func (WindowsButton) Render() {
	fmt.Println("Windows Button")
}

// Mac implementacja
type MacButton struct{}

func (MacButton) Render() {
	fmt.Println("Mac Button")
}

// prosta funkcja zamiast fabryk
func NewButton(os string) Button {
	switch os {
	case "windows":
		return WindowsButton{}
	case "mac":
		return MacButton{}
	default:
		return nil
	}
}

func main() {

	// wybór implementacji
	button := NewButton("windows")

	// użycie interfejsu
	button.Render()
}