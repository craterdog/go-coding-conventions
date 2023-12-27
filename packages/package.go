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
This package file defines the INTERFACE to this package.  Any additions to the
types defined in this file require a MINOR version change.  Any deletions from,
or changes to, the types defined in this file require a MAJOR version change.

The package...

For detailed documentation on this package refer to the wiki:

    https://github.com/<repository-name>/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:

	https://github.com/craterdog/go-coding-conventions/wiki

Additional implementations of the classes provided by this package can be
developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types; and the class implementations only depend
on interfaces, not on each other.
*/
package <packageName>

// PACKAGE CONSTANTS

// This package constant represents...
const <PackageConstantName> = <packageConstantValue>
...

// PACKAGE TYPES

// Specialized Types

// This specialized type definition represents a specialization of the primitive
// Go <primitive type> data type.
type <TypeName> <primitiveType>
...

// Function Types

// This function type defines the signature for any function that...
type <FunctionName>Function func(<arguments>) <AbstractType>
...

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all <interface name> types.
type <InterfaceName> interface {
	<MethodName>(<arguments>) <AbstractType>
	...
}
...

// Abstract Types

// This abstract type defines the set of class constants, constructors and
// functions that must be supported by all <type-name>-class-like types.
type <TypeName>ClassLike interface {
	Get<ConstantName>() <AbstractType>
	...
	<ConstructorName>(<arguments>) <TypeName>Like
	...
	<FunctionName>(<arguments>) <AbstractType>
	...
}
...

// This abstract type defines the set of abstract interfaces that must be
// supported by all <type-name>-like types.
type <TypeName>Like interface {
	<InterfaceName>
	...
}
...

