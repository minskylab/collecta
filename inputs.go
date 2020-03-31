package collecta

// InputType defines a type of input, that is useful to define the primitives
type InputType string

// Text is an input type
const Text InputType = "text"

// Options is an input type
const Options InputType = "options"

// Satisfaction is an input type
const Satisfaction InputType = "satisfaction"

// YesNo is an input type
const YesNo InputType = "yesno"

// Input is a input interface util to define how an input interact with other types
type Input interface {
	Type() InputType
	Value() ([]string, error)
}
