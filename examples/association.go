package main

import (
	fmt "fmt"
)

/********************************* package.go *********************************/

// PACKAGE TYPES

// Specialized Types

// This specialized type definition represents a specialization of the primitive
// Go `any` data type to represent the key type in a key-value pair.
type Key any

// This specialized type definition represents a specialization of the primitive
// Go `any` data type to represent the value type in a key-value pair.
type Value any

// PACKAGE CONSTANTS

// Private Constants

// These private constants define maps to reference multiple singletons for each
// generic class type structure.
var (
	associationClassSingletons = map[string]any{}
)

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all binding association types.
type Binding[K Key, V Value] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all association-like types.
type AssociationLike[K Key, V Value] interface {
	Binding[K, V]
}

// PACKAGE CLASSES

// This function returns a reference to a specific association class type and
// initializes any class constants.
func Association[K Key, V Value]() *associationClass_[K, V] {
	var class *associationClass_[K, V]
	var key = fmt.Sprintf("%T", class)
	var value = associationClassSingletons[key]
	switch actual := value.(type) {
	case *associationClass_[K, V]:
		class = actual
	default:
		class = &associationClass_[K, V]{
			// This class has no class constants.
		}
		associationClassSingletons[key] = class
	}
	return class
}

/******************************* association.go *******************************/

// CLASS TYPE

// This private class type defines the structure associated with the constants,
// constructors and functions for the association class.
type associationClass_[K Key, V Value] struct {
	// This class type has no constants.
}

// CLASS CONSTRUCTORS

// This class constructor creates a new association from the specified key and value.
func (c *associationClass_[K, V]) FromPair(key K, value V) AssociationLike[K, V] {
	var association AssociationLike[K, V]
	association = &association_[K, V]{key, value}
	return association
}

// CLASS METHODS

// This private concrete type encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the association-like abstract type.
type association_[K Key, V Value] struct {
	key   K
	value V
}

// Binding Interface

// This class method returns the key for this association.
func (v *association_[K, V]) GetKey() K {
	return v.key
}

// This class method returns the value for this association.
func (v *association_[K, V]) GetValue() V {
	return v.value
}

// This class method sets the value of this association to a new value.
func (v *association_[K, V]) SetValue(value V) {
	v.value = value
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve a specific association class type.
	var Association = Association[string, int]()

	// Create a new association.
	var key string = "answer"
	var value int = 42
	var association = Association.FromPair(key, value)
	fmt.Printf("key: %q, value: %v\n", association.GetKey(), association.GetValue())

	// Change the value of the association.
	association.SetValue(25)
	fmt.Printf("key: %q, value: %v\n", association.GetKey(), association.GetValue())
}

