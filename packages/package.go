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

All UpperCase constant, type, and function names are visible to other projects
that import this package. All lowercase names are private to this package.
*/
package <packageName>

import (
	fmt "fmt"
	...
)

// PACKAGE CONSTANTS

// This constant represents...
const <ConstantName> =...
...

// This private constant represents...
const <privateConstantName> =...
...

// PACKAGE TYPES

// This type alias represents...
type <TypeName> <primitiveType>

// This singleton defines a unique name space for the <className> class.
var <ClassName> = &<className>s_{
	<className>_(...), // This is a class constant.
	...
}

// PACKAGE FUNCTIONS

// This function returns...
func <FunctionName>(<arguments>) <resultType> {
	var <result> <resultType>
	...
	return <result>
}
...

// This private function returns...
func <privateFunctionName>(<arguments>) <resultType> {
	var <result> <resultType>
	...
	return <result>
}
...
