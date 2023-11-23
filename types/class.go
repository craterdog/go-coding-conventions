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

// CLASS DEFINITIONS

// This private concrete type implements the <ClassName>Like interface.
type <className>_ <primitiveType>

           ***or***

type <className>_ struct {
	<attributeName> <AbstractType>
	...
}

// This private class type defines the structure associated with the class
// constants and class functions for the <className> elements.
type <className>s_ struct {
	<classConstantName> <ClassName>Like
	...
}

// CLASS CONSTANTS

// This class constant represents...
func (c *<className>s_) <ClassConstantName>() <ClassName>Like {
	return c.<classConstantName>
}
...

// CLASS CONSTRUCTORS

// This constructor creates a new <className> from the specified string value.
func (c *<className>s_) FromString(string_ string) <ClassName>Like {
	var <className> <ClassName>Like
	...
	return <className>
}
...

// CLASS METHODS

// <InterfaceName> Interface

// This method returns...
func (v <className>_) <MethodName>(<arguments>) <resultType> {
	var <result> <resultType>
	...
	return <result>
}
...

// CLASS FUNCTIONS

// This function returns...
func (c *<className>s_) <FunctionName>(<arguments>) <resultType> {
	var <result> <resultType>
	...
	return <result>
}
...
