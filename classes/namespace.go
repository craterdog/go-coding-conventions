// Private Class Namespace Type

type <className>Class_ struct {
	<classConstantName> <AbstractType>
	...
}

// Private Class Namespace Reference

var <className>Class = &<className>Class_{
	<classConstantName>: <classConstantValue>,
	...
}

// Public Class Namespace Access

func <ClassName>Class() <ClassName>ClassLike {
	return <className>Class
}

