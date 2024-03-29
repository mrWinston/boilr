package input

import (
	"fmt"
	"io"
	"reflect"
)

// InputType represents an Input from a plate file.
type InputType struct {
	Name string
	Kind reflect.Kind
	Ask  func(string, io.Reader, io.Writer) (interface{}, error)
}

// AvailableInputs holds all inputs that are enabled right now
var AvailableInputs = []*InputType{
	&InputType{
		Name: "bool",
		Kind: reflect.Bool,
		Ask:  AskBool,
	},
	&InputType{
		Name: "string",
		Kind: reflect.String,
		Ask:  AskString,
	},
	&InputType{
		Name: "int",
		Kind: reflect.Int,
		Ask:  AskInt,
	},
	&InputType{
		Name: "list",
		Kind: reflect.Slice,
		Ask:  AskList,
	},
}

// GetInputByKind searches the AvailableInputs for one that matches the
// given reflect.Kind. Returns an error if none is found
func GetInputByKind(kind reflect.Kind) (*InputType, error) {
	for _, input := range AvailableInputs {
		if input.Kind == kind {
			return input, nil
		}
	}
	return nil, fmt.Errorf("Type %v not supported", kind)
}

// GetInputByName searches the AvailableInputs for one that matches the
// given Name. Returns an error if none is found
func GetInputByName(name string) (*InputType, error) {
	for _, input := range AvailableInputs {
		if input.Name == name {
			return input, nil
		}
	}
	return nil, fmt.Errorf("Type %s not supported", name)
}
