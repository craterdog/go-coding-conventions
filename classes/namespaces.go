// Private Class Namespace Type

type <className>Class_[<parameterTypes>] struct {
	<classConstantName> <AbstractType>
	...
}

// Public Class Namespace Access

var <className>Class = map[string]any{}

func <ClassName>Class[<parameterTypes>]() <ClassName>ClassLike[<parameters>] {
	var class *<className>Class_[<parameters>]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = <className>Class[key]
	switch actual := value.(type) {
	case *<className>Class_[<parameters>]:
		// This bound class type already exists.
		class = actual
	default:
		// Create a new bound class type.
		class = &<className>Class_[<parameters>]{
			<classConstantName>: <classConstantValue>,
			...
		}
		<className>Class[key] = class
	}
	return class
}

