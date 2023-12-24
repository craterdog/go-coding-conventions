/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package <packageName>

import (
	fmt "fmt"
)

// CLASS NAMESPACE

// This private type defines the namespace structure associated with the
// constants, constructors and functions for the <ClassName> class namespace.
type <className>Class_[<parameterTypes>] struct {
	<classConstantName> <AbstractType>
	...
}

// Specific Namespace

// This private constant defines the single reference to the <ClassName>
// class namespace.  It also initializes any class constants as needed.
var <className>Class = &<className>Class_{
	<classConstantName>: <classConstantValue>,
	...
}

// This public function returns the single reference to the <ClassName>
// class namespace.
func <ClassName>() *<className>Class_ {
	return <className>Class
}

/************************************* OR ************************************/

// Generic Namespace

// This private constant defines a map to hold all the single references to
// the type specific <ClassName> class namespaces.
var <className>Class = map[string]any{}

// This public function returns the single reference to a type specific
// <ClassName> class namespace.  It also initializes any class constants as
// needed.
func <ClassName>[<parameterTypes>]() *<className>Class_[<parameters>] {
	var class *<className>Class_[<parameters>]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = <className>Class[key]
	switch actual := value.(type) {
	case *<className>Class_[<parameters>]:
		// This bound class type already exists.
		class = actual
	default:
		// Create a new bound class type.
		class = &<className>Class_[<parameters>]{
			<classConstantName>: <classConstantValue>,
			...
		}
		<className>Class[key] = class
	}
	return class
}

// CLASS CONSTANTS

// This public class constant represents...
func (c *<className>Class_[<parameters>]) <ClassConstantName>() <AbstractType> {
	return c.<classConstantName>
}
...

// CLASS CONSTRUCTORS

// This public class constructor creates a new <ClassName> from the specified
// <abstractType> value.
func (c *<className>Class_[<parameters>]) From<AbstractType>(value <AbstractType>) <ClassName>Like[<parameters>] {
	...
	var <className> <ClassName>Like[<parameters>] = &<className>_[<parameters>]{
		<privateAttributeName>: <privateAttributeValue>,
		...
	}
	return <className>
}
...

// CLASS FUNCTIONS

// This public class function returns...
func (c *<className>Class_[<parameters>]) <FunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// CLASS TYPE

// Extended Type

// This private class type extends the primitive Go <primitiveType> data type
// and defines the methods that implement the <ClassName>Like abstract type.
type <className>_ <primitiveType>

// <InterfaceName> Interface

// This public class method...
func (v <className>_) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// Private Interface

// This private class method...
func (v <className>_) <methodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

/************************************* OR ************************************/

// Encapsulated Type

// This private class type encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the <ClassName>Like abstract type.
type <className>_[<parameterTypes>] struct {
	<privateAttributeName> <AbstractType>
	...
}

// <InterfaceName> Interface

// This public class method...
func (v *<className>_[<parameters>]) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// Private Interface

// This private class method...
func (v *<className>_[<parameters>]) <methodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...
