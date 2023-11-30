package main

import (
	fmt "fmt"
)

type Element string

// This interface defines the methods supported by any associative collection.
type Associative interface {
	GetElement(name string) Element
}

// This interface defines the methods supported by any periodic table of elements.
type Periodic interface {
	GetElement(name string) Element
}

// This concrete type implements the Periodic interface.
type Table struct {
}

func (t *Table) GetElement(name string) Element {
	return "Hydrogen"
}

func main() {
	var periodic Periodic = &Table{}
	switch periodic.(type) {
	case Associative:
		fmt.Println("Found an Associative collection.")
	case Periodic:
		fmt.Println("Found a Periodic table.")
	}
	fmt.Printf("The element is %v.\n", periodic.GetElement("H"))
}
