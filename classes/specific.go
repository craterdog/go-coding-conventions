// Specific Namespace

// This private constant defines the single reference to the <ClassName>
// class namespace.  It also initializes any class constants as needed.
var <className>Class = &<className>Class_{
	<classConstantName>: <classConstantValue>,
	...
}

// This public function returns the single reference to the <ClassName>
// class namespace.
func <ClassName>() *<className>Class_ {
	return <className>Class
}
