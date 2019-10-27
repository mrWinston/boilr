package query

import (
	"bytes"
	"io"
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
		want    int
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
			name: "Test Invalid",
			args: args{
				valueName: "zero",
				in:        strings.NewReader("b\nkasj\nkj\n0\n"),
			},
			want:    0,
			wantOut: "Please enter a value for 'zero' [number]: ",
			wantErr: true,
		},
		{
			name: "Test EOF",
			args: args{
				valueName: "zero",
				in:        strings.NewReader("bkasjkj0"),
			},
			want:    0,
			wantOut: "Please enter a value for 'zero' [number]: ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := AskInt(tt.args.valueName, tt.args.in, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("AskInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AskInt() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("AskInt() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
