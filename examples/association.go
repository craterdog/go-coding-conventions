package main

import (
	fmt "fmt"
)

/********************************* package.go *********************************/

// PACKAGE TYPES

// Specialized Types

// This specialized type definition represents a specialization of the primitive
// Go any data type.  This type is used generically to represent elements used
// as keys.
type Key any

// This specialized type definition represents a specialization of the primitive
// Go any data type.  This type is used generically to represent components used
// as values.
type Value any

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all binding associations.  It binds a read-only key with a
// setable value.
type Binding[K Key, V Value] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// Abstract Types

// This abstract type defines the set of class constants, constructors and
// functions that must be supported by all association-class-like types.
type AssociationClassLike[K Key, V Value] interface {
	// constructors
	FromPair(key K, value V) AssociationLike[K, V]
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all association-like types.
type AssociationLike[K Key, V Value] interface {
	Binding[K, V]
}

/******************************* association.go *******************************/

// CLASS NAMESPACE

// Private Class Namespace Type

type associationClass_[K Key, V Value] struct {
	// This class has no class constants.
}

// Private Class Namespace References

var associationClass = map[string]any{}

// Public Class Namespace Access

func AssociationClass[K Key, V Value]() AssociationClassLike[K, V] {
	var class *associationClass_[K, V]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = associationClass[key]
	switch actual := value.(type) {
	case *associationClass_[K, V]:
		// This bound class type already exists.
		class = actual
	default:
		// Create a new bound class type.
		class = &associationClass_[K, V]{
			// This class has no class constants.
		}
		associationClass[key] = class
	}
	return class
}

// Public Class Constructors

func (c *associationClass_[K, V]) FromPair(key K, value V) AssociationLike[K, V] {
	var association = &association_[K, V]{
		key:   key,
		value: value,
	}
	return association
}

// CLASS TYPE

// Private Class Type Definition

type association_[K Key, V Value] struct {
	key   K
	value V
}

// Binding Interface

func (v *association_[K, V]) GetKey() K {
	return v.key
}

func (v *association_[K, V]) GetValue() V {
	return v.value
}

func (v *association_[K, V]) SetValue(value V) {
	v.value = value
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve a specific association class namespace.
	var Association = AssociationClass[string, int]()

	// Create a new association.
	var key string = "answer"
	var value int = 42
	var association = Association.FromPair(key, value)
	fmt.Printf("key: %q, value: %v\n", association.GetKey(), association.GetValue())

	// Change the value of the association.
	association.SetValue(25)
	fmt.Printf("key: %q, value: %v\n", association.GetKey(), association.GetValue())
}
