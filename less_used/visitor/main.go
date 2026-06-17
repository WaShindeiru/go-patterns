package main

import "fmt"

// PROBLEM: Jak wykonywać różne operacje na typach bez rozbudowanej hierarchii klas?

type PDFDocument struct{}
type WordDocument struct{}

// operacja przez type switch
func Export(doc interface{}) {

	switch doc.(type) {
	case PDFDocument:
		fmt.Println("Export PDF")
	case WordDocument:
		fmt.Println("Export Word")
	}
}

func main() {

	doc := PDFDocument{}

	// wykonanie operacji
	Export(doc)
}