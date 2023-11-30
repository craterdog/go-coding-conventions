package main

import "fmt"

/********************************* package.go *********************************/

// PACKAGE TYPES

// Specialized Types

// This specialized type definition represents a specialization of the primitive
// Go `any` data type to represent the key type in a key-value pair.
type Key any

// This specialized type definition represents a specialization of the primitive
// Go `any` data type to represent the value type in a key-value pair.
type Value any

// Function Types

// This function type defines the signature for any function that can determine
// the relative ordering of two specified values. The result must be one of
// the following:
//   - -1: The first value comes before the second value.
//   - 0: The first value is equal to the second value.
//   - 1: The first value comes after the second value.
type RankingFunction func(first, second any) int

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all binding association types.
type Binding[K Key, V Value] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// This abstract interface defines the set of method signatures that must be
// supported by all sequences of values.
type Sequential[V Value] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []V
}

// This abstract interface defines the set of method signatures that must be
// supported by all associative sequences whose values consist of key-value
// pair associations.
type Associative[K Key, V Value] interface {
	GetValue(key K) V
	SetValue(key K, value V)
	RemoveValue(key K) V
	AddAssociation(association Binding[K, V])
}

// This abstract interface defines the set of method signatures that must be
// supported by all sortable sequences whose values may be sorted using various
// ranking algorithms.
type Sortable[V Value] interface {
	SortValues() // Sort the values using a canonical ranking function.
	SortValuesWithRanker(rank RankingFunction)
}

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all association-like types.
type AssociationLike[K Key, V Value] interface {
	Binding[K, V]
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all catalog-like types.
type CatalogLike[K Key, V Value] interface {
	Sequential[Binding[K, V]]
	Associative[K, V]
	//Sortable[Binding[K, V]] // This is commented out to simplify the example.
}

// PACKAGE CLASSES

// This function returns a reference to a specific association class type and
// initializes any class constants.
func Association[K Key, V Value]() *associationClass_[K, V] {
	var class = &associationClass_[K, V]{
		// This class has no class constants.
	}
	return class
}

// This function returns a reference to a specific catalog class type and
// initializes any class constants.
func Catalog[K Key, V Value]() *catalogClass_[K, V] {
	var class = &catalogClass_[K, V]{
		Association[K, V](), // The corresponding association class type.
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

/********************************* catalog.go *********************************/

// CLASS TYPE

// This private class type defines the structure associated with the constants,
// constructors and functions for the catalog class.
type catalogClass_[K Key, V Value] struct {
	association *associationClass_[K, V]
}

// CLASS CONSTANTS

// This class constant is a reference to a specific association class type.
func (c *catalogClass_[K, V]) Association() *associationClass_[K, V] {
	return c.association
}

// CLASS CONSTRUCTORS

// This class constructor creates a new empty catalog.
func (c *catalogClass_[K, V]) FromNothing() CatalogLike[K, V] {
	var catalog CatalogLike[K, V]
	var keys = map[Key]Binding[K, V]{}
	var associations = []Binding[K, V]{}
	catalog = &catalog_[K, V]{keys, associations}
	return catalog
}

// CLASS METHODS

// This private concrete type encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the association-like abstract type.
type catalog_[K Key, V Value] struct {
	keys         map[Key]Binding[K, V]
	associations []Binding[K, V]
}

// Sequential Interface

// This class method determines whether or not this sequence is empty.
func (v *catalog_[K, V]) IsEmpty() bool {
	return len(v.associations) == 0
}

// This class method returns the number of values contained in this sequence.
func (v *catalog_[K, V]) GetSize() int {
	return len(v.associations)
}

// This class method returns an array of the values in this sequence.
func (v *catalog_[K, V]) AsArray() []Binding[K, V] {
	var length = len(v.associations)
	var result = make([]Binding[K, V], length)
	copy(result, v.associations) // Return a copy of the private attribute.
	return result
}

// Associative Interface

// This class method returns from this catalog the value that is associated
// with the specified key.
func (v *catalog_[K, V]) GetValue(key K) V {
	var value V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		// Extract the value.
		value = association.GetValue()
	}
	return value
}

// This class method sets in this catalog the value that is associated with the
// specified key to the specified new value.
func (v *catalog_[K, V]) SetValue(key K, value V) {
	var association, exists = v.keys[key]
	if exists {
		// Set the value of an existing association.
		association.SetValue(value)
	} else {
		// Add a new association.
		association = &association_[K, V]{key, value}
		v.associations = append(v.associations, association)
		v.keys[key] = association
	}
}

// This class method removes from this catalog the value associated with the
// specified key and returns it.
func (v *catalog_[K, V]) RemoveValue(key K) V {
	var old V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		for index, candidate := range v.associations {
			if fmt.Sprint(candidate.GetKey()) == fmt.Sprint(key) {
				v.associations = append(v.associations[:index], v.associations[index+1:]...)
				break
			}
		}
		old = association.GetValue()
		delete(v.keys, key)
	}
	return old
}

// This class method adds to this catalog the specified key-value pair
// association.  If an association with that key already exists in the
// catalog, the value for the association is updated to the new value.
func (v *catalog_[K, V]) AddAssociation(association Binding[K, V]) {
	var key = association.GetKey()
	var value = association.GetValue()
	v.SetValue(key, value)
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	var Catalog = Catalog[string, float64]() // Create a specific catalog class type.
	var Association = Catalog.Association()  // Create a specific association class type.
	var catalog = Catalog.FromNothing()      // Create an empty catalog class instance.
	var association = Association.FromPair("answer", 42.0)
	catalog.AddAssociation(association)
	catalog.SetValue("pi", 3.1415)
	fmt.Printf("empty: %v\n", catalog.IsEmpty())
	fmt.Printf("size: %v\n", catalog.GetSize())
	fmt.Printf("answer: %v\n", catalog.GetValue("answer"))
	fmt.Printf("pi: %v\n", catalog.GetValue("pi"))
	catalog.RemoveValue("answer")
	fmt.Printf("size: %v\n", catalog.GetSize())
}
