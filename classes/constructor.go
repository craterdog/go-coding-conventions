// This public class constructor creates a new <ClassName> from the specified
// <abstractType> value.
func (c *<className>Class_[<parameters>]) From<AbstractType>(value <AbstractType>) <ClassName>Like[<parameters>] {
	...
	var <className> <ClassName>Like[<parameters>] = &<className>_[<parameters>]{
		<privateAttributeName>: <privateAttributeValue>,
		...
	}
	return <className>
}
