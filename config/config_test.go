package config

import (
	"reflect"
	"testing"
)

func TestLoadPlateFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *PlateFile
		wantErr bool
	}{
		{
			name: "Complete 1",
			args: args{
				path: "../test_files/plate_files/complete_1.yml",
			},
			want: &PlateFile{
				Vars: map[string]string{
					"string":  "string",
					"integer": "int",
					"list":    "list",
				},
				Config: Config{
					TemplateRoot: "./complete_1",
				},
				Commands: []CommandConfig{
					CommandConfig{
						Name:    "show me some stuff",
						Command: "ls",
						Workdir: ".",
					},
					CommandConfig{
						Name:    "run a thing",
						Command: "./run.sh",
						Workdir: "./bin",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Minimal",
			args: args{
				path: "../test_files/plate_files/minimal.yml",
			},
			want: &PlateFile{
				Vars: nil,
				Config: Config{
					TemplateRoot: "./minimal",
				},
				Commands: nil,
			},
			wantErr: false,
		},
		{
			name: "Nonexistant Root Entries",
			args: args{
				path: "../test_files/plate_files/nonexistant_root_entries.yml",
			},
			want:    &PlateFile{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadPlateFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPlateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadPlateFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
