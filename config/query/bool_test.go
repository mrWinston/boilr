package query

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestAskBool(t *testing.T) {
	type args struct {
		valueName string
		in        io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantOut string
		wantErr bool
	}{
		{
			name: "Test True",
			args: args{
				valueName: "test",
				in:        strings.NewReader("y\n"),
			},
			want:    true,
			wantOut: "Please enter a value for 'test' [y,n]: ",
			wantErr: false,
		},
		{
			name: "Test False",
			args: args{
				valueName: "test",
				in:        strings.NewReader("n\n"),
			},
			want:    false,
			wantOut: "Please enter a value for 'test' [y,n]: ",
			wantErr: false,
		},
		{
			name: "Test Case",
			args: args{
				valueName: "test",
				in:        strings.NewReader("N\n"),
			},
			want:    false,
			wantOut: "Please enter a value for 'test' [y,n]: ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := AskBool(tt.args.valueName, tt.args.in, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("AskBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AskBool() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("AskBool() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
