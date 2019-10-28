package input

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func AskList(valueName string, in io.Reader, out io.Writer) (interface{}, error) {
	reader := bufio.NewReader(in)
	fmt.Fprintf(out, "Please enter the values for '%s' [list]. Items separated by <return>. End input with an empty line: \n", valueName)
	out_list := []string{}
	var err error

	for input := "placeholder"; input != "\n"; input, err = reader.ReadString('\n') {
		if err != nil {
			return out_list, err
		}

		input = strings.Replace(input, "\n", "", -1)
		out_list = append(out_list, input)
	}
	if err != nil {
		return nil, err
	}
	return out_list, nil
}
