package query

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func AskString(valueName string, in io.Reader, out io.Writer) (interface{}, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter a value for '%s': ", valueName)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Replace(input, "\n", "", -1), nil
}
