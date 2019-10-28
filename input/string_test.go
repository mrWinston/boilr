package input

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestAskString(t *testing.T) {
	type args struct {
		valueName string
		in        io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantOut string
		wantErr bool
	}{
		{
			name: "Test Simple",
			args: args{
				valueName: "test",
				in:        strings.NewReader("test\n"),
			},
			want:    "test",
			wantOut: "Please enter a value for 'test': ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := AskString(tt.args.valueName, tt.args.in, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("AskString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AskString() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("AskString() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
