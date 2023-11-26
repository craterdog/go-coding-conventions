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

// Specialized Types

// This specialized type definition represents a specialization of the primitive
// Go <primitiveType> data type.
type <TypeName> <primitiveType>
...

// Structured Types

// This structured type defines the attributes for a <StructureName> entity.
type <StructureName> struct {
	<AttributeName> <AbstractType>
	...
}
...

// Function Types

// This function type defines the signature for a <FunctionName> function.
type <FunctionName>Function func(<arguments>) <AbstractType>
...

// PACKAGE CONSTANTS

// Public Constants

// This constant represents...
const <PackageConstantName> = <packageConstantValue>
...

// Private Constants

// This private constant represents...
const <privatePackageConstantName> = <privatePackageConstantValue>
...

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

type <InterfaceName> interface {
	<MethodName>(<arguments>) <AbstractType>
	...
}
...

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all <TypeName>Like types.
type <TypeName>Like interface {
	<InterfaceName>
	...
}
...

// PACKAGE CLASSES

// This singleton exports a unique namespace for the <className> class type.
var <ClassName> = &<className>Class_{
	<className>_(...), // Initialize the class constant <privateClassConstantName>.
	...
}
...

// PACKAGE FUNCTIONS

// Public Functions

// This function returns...
func <PackageFunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// Private Functions

// This private function returns...
func <privatePackageFunctionName>(<arguments>) <ResultType> {
	var <result> <ResultType>
	...
	return <result>
}
...
