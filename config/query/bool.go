package query

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var BOOL_TRUE = map[string]bool{
	"y":    true,
	"yes":  true,
	"true": true,
}
var BOOL_FALSE = map[string]bool{
	"n":     true,
	"no":    true,
	"false": true,
}

func AskBool(valueName string, in io.Reader, out io.Writer) (interface{}, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s' [y,n]: ", valueName)
	var input string

	input, err := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, " ", "", -1)
	input = strings.ToLower(input)
	if err != nil {
		return false, err
	}
	if BOOL_TRUE[input] {
		return true, nil
	}
	if BOOL_FALSE[input] {
		return false, nil
	}
	fmt.Fprint(out, "Please enter only 'y' or 'n': ")

	return false, errors.New(fmt.Sprintf("Input string '%s' could not be parsed as bool", input))
}
