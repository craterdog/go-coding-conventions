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

// This private namespace type defines the structure associated with the class
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

// This class constructor creates a new <className> from the specified <abstractType> value.
func (c *<className>Class_) From<AbstractType>(value <AbstractType>) <ClassName>Like {
	var <className> <ClassName>Like
	...
	return <className>
}
...

// CLASS FUNCTIONS

// This class function returns...
func (c *<className>Class_) <FunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// CLASS METHODS

// Extended Type

// This private class type extends the primitive Go <primitiveType> data type
// and defines the methods that implement the <abstractType> abstract type.
type <className>_ <primitiveType>

// <InterfaceName> Interface

// This class method...
func (v <className>_) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

/************************************* OR ************************************/

// Encapsulated Type

// This private class type encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the <abstractType> abstract type.
type <className>_ struct {
	<privateAttributeName> <AbstractType>
	...
}

// <InterfaceName> Interface

// This class method...
func (v *<className>_) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...
