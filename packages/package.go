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

All UpperCase type, constant, abstraction, function and class names are visible
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

// This public constant represents...
const <PublicPackageConstantName> = <publicPackageConstantValue>
...

// Private Constants

// This private constant represents...
const <privatePackageConstantName> = <privatePackageConstantValue>
...

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all <interfaceName> types.
type <InterfaceName> interface {
	<MethodName>(<arguments>) <AbstractType>
	...
}
...

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all <typeName>-like types.
type <TypeName>Like interface {
	<InterfaceName>
	...
}
...

// PACKAGE FUNCTIONS

// Public Functions

// This public function returns...
func <PublicPackageFunctionName>(<arguments>) <AbstractType> {
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

// PACKAGE CLASSES

// This function returns a reference to the <className> class type and
// initializes any class constants.
func <ClassName>() *<className>Class_ {
	var class = &<className>Class_{
		<className>_(...), // Initialize the class constant <privateClassConstantName>.
		...
	}
	return class
}
...
