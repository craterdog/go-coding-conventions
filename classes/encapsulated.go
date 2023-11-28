// Encapsulated Type

// This private class type encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the <abstractType> abstract type.
type <className>_ struct {
	<privateAttributeName> <AbstractType>
	...
}

// <InterfaceName> Interface

// This class method...
func (v *<className>_) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...
