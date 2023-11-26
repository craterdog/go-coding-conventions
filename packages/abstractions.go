// PACKAGE ABSTRACTIONS

// Abstract Interfaces

type <InterfaceName> interface {
	<MethodName>(<arguments>) <ResultType>
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
