package typeconv

import (
	"errors"
	"fmt"
)

// StringifySlice accepts any slice and return a string version of that.
// It supports conversion of bool, int, float and string values. Complex values
// lists, structs and maps are not supported. In that case, error is returned
func StringifySlice(in []interface{}) ([]string, error) {
	var out []string
	for _, v := range in {
		switch v.(type) {
		case bool:
			out = append(out, fmt.Sprintf("%t", v.(bool)))
		case string:
			out = append(out, v.(string))
		case int:
			out = append(out, fmt.Sprintf("%d", v.(int)))
		case float32:
			out = append(out, fmt.Sprintf("%g", v.(float32)))
		case float64:
			out = append(out, fmt.Sprintf("%g", v.(float64)))
		default:
			return nil, errors.New("Type is neither bool, string, int, or float")
		}

	}

	return out, nil
}
