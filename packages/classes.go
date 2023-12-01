// This function returns a reference to the <className> class type and
// initializes any class constants.
func <ClassName>() *<className>Class_ {
	var class = &<className>Class_{
		<className>_(...), // Initialize the class constant <privateClassConstantName>.
		...
	}
	return class
}
...
