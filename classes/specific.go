// Specific Namespace

// This private constant defines the singleton reference to the <ClassName>
// class namespace.  It also initializes any class constants as needed.
var <className>ClassSingleton = &<className>Class_{
	<classConstantName>: <classConstantValue>,
	...
}

// This public function returns the singleton reference to the <ClassName>
// class namespace.
func <ClassName>() *<className>Class_ {
	return <className>ClassSingleton
}
