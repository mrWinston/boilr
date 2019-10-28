package input

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestAskList(t *testing.T) {
	type args struct {
		valueName string
		in        io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantOut string
		wantErr bool
	}{
		{
			name: "Test empty",
			args: args{
				valueName: "empty",
				in:        strings.NewReader("\n"),
			},
			want:    []string{},
			wantOut: "Please enter the values for 'empty' [list]. Items separated by <return>. End input with an empty line: \n",
			wantErr: false,
		},
		{
			name: "Test one Element",
			args: args{
				valueName: "list",
				in:        strings.NewReader("one\n\n"),
			},
			want:    []string{"one"},
			wantOut: "Please enter the values for 'list' [list]. Items separated by <return>. End input with an empty line: \n",
			wantErr: false,
		},
		{
			name: "Test Some Elements",
			args: args{
				valueName: "list",
				in:        strings.NewReader("one\ntwo\nthree\nfour\nfive\nsix\n\n"),
			},
			want:    []string{"one", "two", "three", "four", "five", "six"},
			wantOut: "Please enter the values for 'list' [list]. Items separated by <return>. End input with an empty line: \n",
			wantErr: false,
		},
		{
			name: "Test No Input",
			args: args{
				valueName: "list",
				in:        strings.NewReader(""),
			},
			want:    nil,
			wantOut: "Please enter the values for 'list' [list]. Items separated by <return>. End input with an empty line: \n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := AskList(tt.args.valueName, tt.args.in, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("AskList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if isStringSliceEqualToInterface(got, tt.want) {
				t.Errorf("AskList() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("AskList() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func isStringSliceEqualToInterface(inter interface{}, slice []string) bool {
	if inter == nil && slice == nil {
		return true
	}
	if inter == nil || slice == nil {
		return false
	}

	switch reflect.TypeOf(inter).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(inter)
		if s.Len() != len(slice) {
			return false
		}

		for i := 0; i < len(slice); i++ {
			if s.Index(i).Interface() != slice[i] {
				return false
			}
		}
		return true

	default:
		return false
	}

}
