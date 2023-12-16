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

    https://github.com/craterdog/<repository-name>/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions
located here:

	https://github.com/craterdog/go-coding-conventions/wiki
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

// This function type defines the signature for any function that...
type <FunctionName>Function func(<arguments>) <AbstractType>
...

// PACKAGE CONSTANTS

// This constant represents...
const <PackageConstantName> = <packageConstantValue>
...

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all <InterfaceName> types.
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

// PACKAGE FUNCTIONS

// This function returns...
func <PackageFunctionName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...
