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
// the relative ordering of two values. The result must be one of the following:
//
//	-1: The first value is less than the second value.
//	 0: The first value is equal to the second value.
//	 1: The first value is more than the second value.
//
// The meaning of "less" and "more" is determined by the specific function that
// implements this signature.
type RankingFunction func(first Value, second Value) int

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

// This abstract interface defines the set of method signatures that must be
// supported by all sequences of values.
type Sequential[V Value] interface {
	AsArray() []V
	GetSize() int
	IsEmpty() bool
}

// This abstract interface defines the set of method signatures that must be
// supported by all associative sequences whose values consist of key-value
// pair associations.
type Associative[K Key, V Value] interface {
	AddAssociation(association Binding[K, V])
	GetValue(key K) V
	RemoveValue(key K) V
	SetValue(key K, value V)
}

// This abstract interface defines the set of method signatures that must be
// supported by all sortable sequences whose values may be sorted using various
// ranking algorithms.
type Sortable[V Value] interface {
	SortValues() // Sort the values using a canonical ranking function.
	SortValuesWithRanker(rank RankingFunction)
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

// This abstract type defines the set of class constants, constructors and
// functions that must be supported by all catalog-class-like types.
type CatalogClassLike[K Key, V Value] interface {
	// constants
	GetAssociationClass() AssociationClassLike[K, V]
	// constructors
	Empty() CatalogLike[K, V]
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all catalog-like types.
type CatalogLike[K Key, V Value] interface {
	Sequential[Binding[K, V]]
	Associative[K, V]
	//Sortable[Binding[K, V]] // This is commented out to simplify the example.
}

/******************************* association.go *******************************/

// CLASS NAMESPACE

// Private Class Namespace Type

type associationClass_[K Key, V Value] struct {
	// This class has no class constants.
}

// Public Class Namespace Access

var associationClass = map[string]any{}

func Association[K Key, V Value]() AssociationClassLike[K, V] {
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

/********************************* catalog.go *********************************/

// CLASS NAMESPACE

// Private Class Namespace Type

type catalogClass_[K Key, V Value] struct {
	associationClass AssociationClassLike[K, V]
}

// Public Class Namespace Access

var catalogClass = map[string]any{}

func Catalog[K Key, V Value]() CatalogClassLike[K, V] {
	var class *catalogClass_[K, V]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = catalogClass[key]
	switch actual := value.(type) {
	case *catalogClass_[K, V]:
		// This bound class type already exists.
		class = actual
	default:
		// Create a new bound class type.
		class = &catalogClass_[K, V]{
			associationClass: Association[K, V](),
		}
		catalogClass[key] = class
	}
	return class
}

// Public Class Constants

func (c *catalogClass_[K, V]) GetAssociationClass() AssociationClassLike[K, V] {
	return c.associationClass
}

// Public Class Constructors

func (c *catalogClass_[K, V]) Empty() CatalogLike[K, V] {
	var associations = []Binding[K, V]{}
	var keys = map[Key]Binding[K, V]{}
	var catalog CatalogLike[K, V] = &catalog_[K, V]{
		associations: associations,
		keys:         keys,
	}
	return catalog
}

// CLASS TYPE

// Private Class Type Definition

type catalog_[K Key, V Value] struct {
	associations []Binding[K, V]
	keys         map[Key]Binding[K, V]
}

// Sequential Interface

func (v *catalog_[K, V]) AsArray() []Binding[K, V] {
	var length = len(v.associations)
	var result = make([]Binding[K, V], length)
	copy(result, v.associations) // Return a copy of the private attribute.
	return result
}

func (v *catalog_[K, V]) GetSize() int {
	return len(v.associations)
}

func (v *catalog_[K, V]) IsEmpty() bool {
	return len(v.associations) == 0
}

// Associative Interface

func (v *catalog_[K, V]) AddAssociation(association Binding[K, V]) {
	var key = association.GetKey()
	var value = association.GetValue()
	v.SetValue(key, value)
}

func (v *catalog_[K, V]) GetValue(key K) V {
	var value V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		// Extract the value.
		value = association.GetValue()
	}
	return value
}

func (v *catalog_[K, V]) RemoveValue(key K) V {
	var old V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		for index, candidate := range v.associations {
			if fmt.Sprint(candidate.GetKey()) == fmt.Sprint(key) {
				v.associations = append(
					v.associations[:index],
					v.associations[index+1:]...,
				)
				break
			}
		}
		old = association.GetValue()
		delete(v.keys, key)
	}
	return old
}

func (v *catalog_[K, V]) SetValue(key K, value V) {
	var association, exists = v.keys[key]
	if exists {
		// Set the value of an existing association.
		association.SetValue(value)
	} else {
		// Add a new association.
		association = &association_[K, V]{
			key:   key,
			value: value,
		}
		v.associations = append(v.associations, association)
		v.keys[key] = association
	}
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve a specific catalog class namespace.
	var Catalog = Catalog[string, float64]()

	// Retrieve the corresponding association class namespace.
	var Association = Catalog.GetAssociationClass()

	// Create a new empty catalog class instance.
	var catalog = Catalog.Empty()

	// Add a new association to the catalog.
	var association = Association.FromPair("answer", 42.0)
	catalog.AddAssociation(association)

	// Set a new value associated with a key to the catalog.
	catalog.SetValue("pi", 3.1415)
	fmt.Printf("empty: %v\n", catalog.IsEmpty())
	fmt.Printf("size: %v\n", catalog.GetSize())
	fmt.Printf("answer: %v\n", catalog.GetValue("answer"))
	fmt.Printf("pi: %v\n", catalog.GetValue("pi"))

	// Remove the value associated with a key from the catalog.
	catalog.RemoveValue("answer")
	fmt.Printf("size: %v\n", catalog.GetSize())
}
