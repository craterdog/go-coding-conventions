// This function returns a reference to the <className> class type singleton.
func <ClassName>() *<className>Class_ {
	return <className>ClassSingleton
}
...

// This function returns a reference to a specific <className> class type
// singleton and initializes any class constants.
func <ClassName>[<parameterTypes>]() *<className>Class_[<parameters>] {
	var class *<className>Class_[<parameters>]
	var key = fmt.Sprintf("%T", class)
	var value = <className>ClassSingletons[key]
	switch actual := value.(type) {
	case *<className>Class_[<parameters>]:
		class = actual
	default:
		class = &<className>Class_[<parameters>]{
			<classConstantValue>,
			...
		}
		<className>ClassSingletons[key] = class
	}
	return class
}
...

