package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func AskInt(valueName string, in io.Reader, out io.Writer) (interface{}, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s' [number]: ", valueName)

	var input string

	var err error
	input, err = reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.Replace(input, "\n", "", -1)

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return value, nil
}
