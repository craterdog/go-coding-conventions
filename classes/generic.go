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

type <className>Class_[<parameterTypes>] struct {
	<classConstantName> <AbstractType>
	...
}

// Public Class Namespace Access

var <className>Class = map[string]any{}

func <ClassName>Class[<parameterTypes>]() <ClassName>ClassLike[<parameters>] {
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

// Public Class Constants

func (c *<className>Class_[<parameters>]) Get<ClassConstantName>() <AbstractType> {
	return c.<classConstantName>
}
...

// Public Class Constructors

func (c *<className>Class_[<parameters>]) <ConstructorName>(<arguments>) <ClassName>Like[<parameters>] {
	...
	var <className> = &<className>_[<parameters>]{
		<privateAttributeName>: <privateAttributeValue>,
		...
	}
	return <className>
}
...

// Public Class Functions

func (c *<className>Class_[<parameters>]) <FunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// CLASS TYPE

// Private Class Type Definition

type <className>_[<parameterTypes>] struct {
	<privateAttributeName> <AbstractType>
	...
}

// <InterfaceName> Interface

func (v *<className>_[<parameters>]) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// Private Interface

func (v *<className>_[<parameters>]) <methodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

