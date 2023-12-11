// Encapsulated Type

// This private class type encapsulates a Go structure containing private
// attributes that can only be accessed and manipulated using methods that
// implement the <className>-like abstract type.
type <className>_[<parameterTypes>] struct {
	<privateAttributeName> <AbstractType>
	...
}

// <InterfaceName> Interface

// This public class method...
func (v *<className>_[<parameters>]) <MethodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...

// Private Interface

// This private class method...
func (v *<className>_[<parameters>]) <methodName>(<arguments>) <AbstractType> {
	var <result> <AbstractType>
	...
	return <result>
}
...
