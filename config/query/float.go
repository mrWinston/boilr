package query

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func AskFloat(valueName string, in io.Reader, out io.Writer) (interface{}, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s' [float]: ", valueName)

	var input string

	var err error
	input, err = reader.ReadString('\n')

	if err != nil {
		return 0.0, err
	}

	input = strings.Replace(input, "\n", "", -1)

	value, err := strconv.ParseFloat(input, 32)

	if err != nil {
		return 0.0, err
	}
	return value, nil
}
