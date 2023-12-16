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
//
//	-1: The first value is less than the second value.
//	 0: The first value is equal to the second value.
//	 1: The first value is more than the second value.
type RankingFunction func(first Value, second Value) int

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all binding Associations.  It binds a read-only key with a
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
// pair Associations.
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

// This abstract type defines the set of abstract interfaces that must be
// supported by all AssociationLike types.
type AssociationLike[K Key, V Value] interface {
	Binding[K, V]
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all CatalogLike types.
type CatalogLike[K Key, V Value] interface {
	Sequential[Binding[K, V]]
	Associative[K, V]
	//Sortable[Binding[K, V]] // This is commented out to simplify the example.
}

/******************************* association.go *******************************/

// CLASS NAMESPACE

// This private type defines the namespace structure associated with the
// constants, constructors and functions for the Association class namespace.
type associationClass_[K Key, V Value] struct {
	// This class has no class constants.
}

// This private constant defines a map to hold all the singleton references to
// the type specific Association class namespaces.
var associationClassSingletons = map[string]any{}

// This public function returns the singleton reference to a type specific
// Association class namespace.  It also initializes any class constants as
// needed.
func Association[K Key, V Value]() *associationClass_[K, V] {
	var class *associationClass_[K, V]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = associationClassSingletons[key]
	switch actual := value.(type) {
	case *associationClass_[K, V]:
		// This bound class type already exists.
		class = actual
	default:
		// Create a new bound class type.
		class = &associationClass_[K, V]{
			// This class has no class constants.
		}
		associationClassSingletons[key] = class
	}
	return class
}

// CLASS CONSTRUCTORS

// This public class constructor creates a new Association from the specified
// key and value.
func (c *associationClass_[K, V]) FromPair(key K, value V) AssociationLike[K, V] {
	var association = &association_[K, V]{
		key:   key,
		value: value,
	}
	return association
}

// CLASS TYPE

// Encapsulated Type

// This private class structure encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the AssociationLike abstract type.  The attributes maintain the
// information about a key-value pair. This type is parameterized as follows:
//   - K is a primitive type of key.
//   - V is any type of value.
//
// This structure is used by the catalog class to maintain its Associations.
type association_[K Key, V Value] struct {
	key   K
	value V
}

// Binding Interface

// This public class method returns the key for this Association.
func (v *association_[K, V]) GetKey() K {
	return v.key
}

// This public class method returns the value for this Association.
func (v *association_[K, V]) GetValue() V {
	return v.value
}

// This public class method sets the value of this Association to a new value.
func (v *association_[K, V]) SetValue(value V) {
	v.value = value
}

/********************************* catalog.go *********************************/

// CLASS NAMESPACE

// This private type defines the namespace structure associated with the
// constants, constructors and functions for the Catalog class namespace.
type catalogClass_[K Key, V Value] struct {
	association *associationClass_[K, V]
}

// This private constant defines a map to hold all the singleton references to
// the type specific Catalog class namespaces.
var catalogClassSingletons = map[string]any{}

// This public function returns the singleton reference to a type specific
// Catalog class namespace.  It also initializes any class constants as needed.
func Catalog[K Key, V Value]() *catalogClass_[K, V] {
	var class *catalogClass_[K, V]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = catalogClassSingletons[key]
	switch actual := value.(type) {
	case *catalogClass_[K, V]:
		// This bound class type already exists.
		class = actual
	default:
		// Create a new bound class type.
		class = &catalogClass_[K, V]{
			Association[K, V](), // The corresponding Association class namespace.
		}
		catalogClassSingletons[key] = class
	}
	return class
}

// CLASS CONSTANTS

// This public class constant returns a reference to the corresponding specific
// Association class namespace.
func (c *catalogClass_[K, V]) Association() *associationClass_[K, V] {
	return c.association
}

// CLASS CONSTRUCTORS

// This public class constructor creates a new empty Catalog.
func (c *catalogClass_[K, V]) Empty() CatalogLike[K, V] {
	var associations = []Binding[K, V]{}
	var keys = map[Key]Binding[K, V]{}
	var catalog CatalogLike[K, V] = &catalog_[K, V]{
		associations,
		keys,
	}
	return catalog
}

// CLASS TYPE

// Encapsulated Type

// This private class structure encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the CatalogLike abstract type.  This type is parameterized as
// follows:
//   - K is a primitive type of key.
//   - V is any type of value.
type catalog_[K Key, V Value] struct {
	associations []Binding[K, V]
	keys         map[Key]Binding[K, V]
}

// Sequential Interface

// This public class method returns an array of the values in this sequence.
func (v *catalog_[K, V]) AsArray() []Binding[K, V] {
	var length = len(v.associations)
	var result = make([]Binding[K, V], length)
	copy(result, v.associations) // Return a copy of the private attribute.
	return result
}

// This public class method returns the number of values contained in this
// sequence.
func (v *catalog_[K, V]) GetSize() int {
	return len(v.associations)
}

// This public class method determines whether or not this sequence is empty.
func (v *catalog_[K, V]) IsEmpty() bool {
	return len(v.associations) == 0
}

// Associative Interface

// This public class method adds to this Catalog the specified key-value pair
// Association.  If an Association with that key already exists in the
// Catalog, the value for the Association is updated to the new value.
func (v *catalog_[K, V]) AddAssociation(association Binding[K, V]) {
	var key = association.GetKey()
	var value = association.GetValue()
	v.SetValue(key, value)
}

// This public class method returns from this Catalog the value that is
// associated with the specified key.
func (v *catalog_[K, V]) GetValue(key K) V {
	var value V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		// Extract the value.
		value = association.GetValue()
	}
	return value
}

// This public class method removes from this Catalog the value associated with
// the specified key and returns it.
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

// This public class method sets in this Catalog the value that is associated
// with the specified key to the specified new value.
func (v *catalog_[K, V]) SetValue(key K, value V) {
	var association, exists = v.keys[key]
	if exists {
		// Set the value of an existing Association.
		association.SetValue(value)
	} else {
		// Add a new Association.
		association = &association_[K, V]{key, value}
		v.associations = append(v.associations, association)
		v.keys[key] = association
	}
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve a specific Catalog class namespace.
	var Catalog = Catalog[string, float64]()

	// Retrieve the corresponding Association class namespace.
	var Association = Catalog.Association()

	// Create a new empty Catalog class instance.
	var catalog = Catalog.Empty()

	// Add a new Association to the Catalog.
	var association = Association.FromPair("answer", 42.0)
	catalog.AddAssociation(association)

	// Set a new value associated with a key to the Catalog.
	catalog.SetValue("pi", 3.1415)
	fmt.Printf("empty: %v\n", catalog.IsEmpty())
	fmt.Printf("size: %v\n", catalog.GetSize())
	fmt.Printf("answer: %v\n", catalog.GetValue("answer"))
	fmt.Printf("pi: %v\n", catalog.GetValue("pi"))

	// Remove the value associated with a key from the Catalog.
	catalog.RemoveValue("answer")
	fmt.Printf("size: %v\n", catalog.GetSize())
}
