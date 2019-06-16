package query

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
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

func AskBool(valueName string, loop bool, in io.Reader, out io.Writer) (bool, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s' [y,n]: ", valueName)
	var input string
	for ok := true; ok; ok = loop {
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
	}

	return false, errors.New(fmt.Sprintf("Input string '%s' could not be parsed as bool", input))
}

func AskString(valueName string, in io.Reader, out io.Writer) (string, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s': ", valueName)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Replace(input, "\n", "", -1), nil
}

func AskInt(valueName string, loop bool, in io.Reader, out io.Writer) (int, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s' [number]: ", valueName)

	var input string

	for ok := true; ok; ok = loop {
		var err error
		input, err = reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		input = strings.Replace(input, "\n", "", -1)

		value, err := strconv.Atoi(input)
		if err == nil {
			return value, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("Input string '%s' could not be parsed as int", input))
}
