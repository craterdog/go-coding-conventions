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
