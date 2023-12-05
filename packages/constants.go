// Public Constants

// This public constant represents...
const <PublicPackageConstantName> = <publicPackageConstantValue>
...

// Private Constants

// This private constant represents...
const <privatePackageConstantName> = <privatePackageConstantValue>
...

// These private constants implement the singleton pattern to provide a single
// reference to each class type structure.  They also initialize any class
// constants.
var (
	<className>ClassSingleton = &<className>Class_{
		<className>_(...),
		...
	}
	...
)

// These private constants define maps to reference multiple singletons for each
// generic class type structure.
var (
	<genericClassName>ClassSingletons = map[string]any{}
	...
)

