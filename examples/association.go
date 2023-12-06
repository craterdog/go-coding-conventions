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

// Function Types

// This function type defines the signature for any function that can determine
// the relative ordering of two specified values. The result must be one of
// the following:
//   - -1: The first value is less than the second value.
//   - 0: The first value is equal to the second value.
//   - 1: The first value is more than the second value.
type RankingFunction func(first Value, second Value) int

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all binding associations.  It binds a readonly key with a
// setable value.
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

/******************************* association.go *******************************/

// CLASS NAMESPACE

// This private type defines the namepace structure associated with the
// constants, constructors and functions for the association class namespace.
type associationClass_[K Key, V Value] struct {
	// This class has no class constants.
}

// This private constant defines a map to hold all the singleton references to
// the type specific association class namespaces.
var associationClassSingletons = map[string]any{}

// This public function returns the singleton reference to a type specific
// association class namespace.  It also initializes any class constants as
// needed.
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

// CLASS CONSTRUCTORS

// This public class constructor creates a new association from the specified
// key and value.
func (c *associationClass_[K, V]) FromPair(key K, value V) AssociationLike[K, V] {
	var association = &association_[K, V]{key, value}
	return association
}

// CLASS TYPE

// Encapsulated Type

// This private class structure encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the association-like abstract type.  The attributes maintain the
// information about a key-value pair. This type is parameterized as follows:
//   - K is a primitive type of key.
//   - V is any type of value.
//
// This structure is used by the catalog class to maintain its associations.
type association_[K Key, V Value] struct {
	key   K
	value V
}

// Binding Interface

// This public class method returns the key for this association.
func (v *association_[K, V]) GetKey() K {
	return v.key
}

// This public class method returns the value for this association.
func (v *association_[K, V]) GetValue() V {
	return v.value
}

// This public class method sets the value of this association to a new value.
func (v *association_[K, V]) SetValue(value V) {
	v.value = value
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve a specific association class namespace.
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

