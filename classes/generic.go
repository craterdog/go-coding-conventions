// Generic Namespace

// This private constant defines a map to hold all the singleton references to
// the type specific <ClassName> class namespaces.
var <className>ClassSingletons = map[string]any{}

// This public function returns the singleton reference to a type specific
// <ClassName> class namespace.  It also initializes any class constants as
// needed.
func <ClassName>[<parameterTypes>]() *<className>Class_[<parameters>] {
	var class *<className>Class_[<parameters>]
	var key = fmt.Sprintf("%T", class) // The name of the bound class type.
	var value = <className>ClassSingletons[key]
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
		<className>ClassSingletons[key] = class
	}
	return class
}
