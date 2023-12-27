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

// Private Class Namespace Type

type <className>Class_ struct {
	<classConstantName> <AbstractType>
	...
}

// Public Class Namespace Access

var <className>Class = &<className>Class_{
	<classConstantName>: <classConstantValue>,
	...
}

func <ClassName>() <ClassName>ClassLike {
	return <className>Class
}

// Public Class Constants

func (c *<className>Class_) Get<ClassConstantName>() <AbstractType> {
	return c.<classConstantName>
}
...

// Public Class Constructors

func (c *<className>Class_) <ConstructorName>(<arguments>) <ClassName>Like {
	...
	var <className> = &<className>_{
		<privateAttributeName>: <privateAttributeValue>,
		...
	}
	return <className>
}
...

// Public Class Functions

func (c *<className>Class_) <FunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// CLASS TYPE

// Private Class Type Definition

type <className>_ struct {
	<privateAttributeName> <AbstractType>
	...
}

// <InterfaceName> Interface

func (v *<className>_) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// Private Interface

func (v *<className>_) <methodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

