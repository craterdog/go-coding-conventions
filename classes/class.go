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

// CLASS TYPES

// Extended Types

// This extended class type extends the primitive Go <primitiveType> data type
// whose value can be accessed and manipulated by the methods that implement one
// or more abstract interfaces.
type <className>_ <primitiveType>
...

// Encapsulated Types

// This encapsulated class type defines a set of primitive attributes that can
// be accessed and manipulated by the methods that implement one or more
// abstract interfaces.
type <className>_ struct {
	<privateAttributeName> <AbstractType>
	...
}
...

// CLASS METHODS

// <InterfaceName> Interface

// This method returns...
func (v <className>_) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// CLASS NAMESPACES

// This namespace class type defines the structure associated with the class
// constants, constructors and functions for the <className> class.
type <className>Class_ struct {
	<classConstantName> <AbstractType>
	...
}
...

// CLASS CONSTANTS

// This class constant represents...
func (c *<className>Class_) <ClassConstantName>() <AbstractType> {
	return c.<classConstantName>
}
...

// CLASS CONSTRUCTORS

// This constructor creates a new instance of a <className> from the specified
// <abstractType> value.
func (c *<className>Class_) From<AbstractType>(value <AbstractType>) <ClassName>Like {
	var <className> <ClassName>Like
	...
	return <className>
}
...

// CLASS FUNCTIONS

// This function returns...
func (c *<className>Class_) <FunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...
