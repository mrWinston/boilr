package input

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestAskFloat(t *testing.T) {
	type args struct {
		valueName string
		in        io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantOut string
		wantErr bool
	}{
		{
			name: "Test 0",
			args: args{
				valueName: "zero",
				in:        strings.NewReader("0\n"),
			},
			want:    0.0,
			wantOut: "Please enter a value for 'zero' [float]: ",
			wantErr: false,
		},
		{
			name: "Test Invalid",
			args: args{
				valueName: "zero",
				in:        strings.NewReader("b\nkasj\nkj\n0\n"),
			},
			want:    0.0,
			wantOut: "Please enter a value for 'zero' [float]: ",
			wantErr: true,
		},
		{
			name: "Test EOF",
			args: args{
				valueName: "zero",
				in:        strings.NewReader(""),
			},
			want:    0.0,
			wantOut: "Please enter a value for 'zero' [float]: ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := AskFloat(tt.args.valueName, tt.args.in, out)
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
