/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

/*
This package defines...
For detailed documentation on this package refer to the wiki:
 * https://github.com/craterdog/<repository-name>/wiki

All UpperCase type, constant, abstraction, class and function names are visible
to other projects that import this package. All lowercase names are private to
this package.
*/
package <packageName>

import (
	fmt "fmt"
	...
)

// PACKAGE TYPES

// Primitive Types

// This type definition represents...
type <TypeName> <primitiveType>
...

// Structured Types

// This structured type defines the attributes for a <StructureName> entity.
type <StructureName> struct {
	<AttributeName> <AttributeType>
	...
}
...

// PACKAGE CONSTANTS

// Public Constants

// This constant represents...
const <ConstantName> = <ConstantValue>
...

// Private Constants

// This private constant represents...
const <privateConstantName> = <PrivateConstantValue>
...

// PACKAGE ABSTRACTIONS

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all <TypeName>Like types.
type <TypeName>Like interface {
	<InterfaceName>
	...
}
...

// Abstract Interfaces

type <InterfaceName> interface {
	<MethodName>(<arguments>) <ResultType>
	...
}
...

// PACKAGE CLASSES

// Primitive Classes

// This singleton defines a unique name space for the <ClassName> primitive
// class.
var <ClassName> = &<className>s_{
	<className>_(...), // This is a class constant.
	...
}
...

// Structured Classes

// This singleton defines a unique name space for the <ClassName> structured
// class.
var <ClassName> = &<className>s_{
	<className>_(...), // This is a class constant.
	...
}
...

// PACKAGE FUNCTIONS

// Public Functions

// This function returns...
func <FunctionName>(<arguments>) <ResultType> {
	var <result> <ResultType>
	...
	return <result>
}
...

// Private Functions

// This private function returns...
func <privateFunctionName>(<arguments>) <ResultType> {
	var <result> <ResultType>
	...
	return <result>
}
...
