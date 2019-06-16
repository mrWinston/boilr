package typeconv

import (
	"reflect"
	"testing"
)

func TestStringifySlice(t *testing.T) {
	type args struct {
		in []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "StringArray",
			args: args{
				in: []interface{}{
					"one",
					"two",
					"three",
				},
			},
			want: []string{
				"one",
				"two",
				"three",
			},
			wantErr: false,
		},
		{
			name: "Bool Array",
			args: args{
				in: []interface{}{
					true,
					false,
					true,
				},
			},
			want: []string{
				"true",
				"false",
				"true",
			},
			wantErr: false,
		},
		{
			name: "Int Array",
			args: args{
				in: []interface{}{
					1,
					2,
					3,
				},
			},
			want: []string{
				"1",
				"2",
				"3",
			},
			wantErr: false,
		},
		{
			name: "Float Array",
			args: args{
				in: []interface{}{
					1.1,
					2.2,
					3.3,
				},
			},
			want: []string{
				"1.1",
				"2.2",
				"3.3",
			},
			wantErr: false,
		},
		{
			name: "Mixed Array",
			args: args{
				in: []interface{}{
					"One",
					false,
					3,
					4.4,
				},
			},
			want: []string{
				"One",
				"false",
				"3",
				"4.4",
			},
			wantErr: false,
		},
		{
			name: "Invalid Type",
			args: args{
				in: []interface{}{
					[]int{1},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringifySlice(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringifySlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringifySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
