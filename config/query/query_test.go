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
		loop      bool
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
				loop:      false,
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
				loop:      false,
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
				loop:      false,
				in:        strings.NewReader("N\n"),
			},
			want:    false,
			wantOut: "Please enter a value for 'test' [y,n]: ",
			wantErr: false,
		},
		{
			name: "Test Loop",
			args: args{
				valueName: "test",
				loop:      true,
				in:        strings.NewReader("k\nk\nk\ny\n"),
			},
			want:    true,
			wantOut: "Please enter a value for 'test' [y,n]: ",
			wantErr: false,
		},
		{
			name: "Test Wrong Input",
			args: args{
				valueName: "test",
				loop:      false,
				in:        strings.NewReader("k\n"),
			},
			want:    false,
			wantOut: "Please enter a value for 'test' [y,n]: ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := AskBool(tt.args.valueName, tt.args.loop, tt.args.in, out)
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

func TestAskInt(t *testing.T) {
	type args struct {
		valueName string
		loop      bool
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
			name: "Test 0",
			args: args{
				valueName: "zero",
				loop:      false,
				in:        strings.NewReader("0\n"),
			},
			want:    0,
			wantOut: "Please enter a value for 'zero' [number]: ",
			wantErr: false,
		},
		{
			name: "Test 0 Loop",
			args: args{
				valueName: "zero",
				loop:      true,
				in:        strings.NewReader("b\nkasj\nkj\n0\n"),
			},
			want:    0,
			wantOut: "Please enter a value for 'zero' [number]: ",
			wantErr: false,
		},
		{
			name: "Test Invalid",
			args: args{
				valueName: "zero",
				loop:      false,
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
				loop:      false,
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
			got, err := AskInt(tt.args.valueName, tt.args.loop, tt.args.in, out)
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
