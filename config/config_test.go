package config

import (
	"reflect"
	"testing"
)

func TestLoadTemplateConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Load Test Config",
			args: args{
				path: "../test_files/test.plate",
			},
			want: map[string]string{
				"TEMPLATE_ROOT": "./test1",
				"one":           "list",
				"two":           "string",
			},
			wantErr: false,
		},
		{
			name: "Load Invalid Config",
			args: args{
				path: "../test_files/test_invalid.plate",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Load Nonexistent Config",
			args: args{
				path: "./alshflkasfhlkfsh",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadTemplateConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTemplateConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadTemplateConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
