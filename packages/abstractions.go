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

