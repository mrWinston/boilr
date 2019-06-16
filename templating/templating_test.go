package templating

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/flosch/pongo2"
)

func Test_preprocessTemplateRoot(t *testing.T) {
	simpleTestFolder, err := filepath.Abs("../test_files/template_dirs/simple_test/")
	if err != nil {
		t.Error(err)
	}
	moreElaborateTestFolder, err := filepath.Abs("../test_files/template_dirs/test1/")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		templateRoot   string
		outRoot        string
		defaultContext pongo2.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []renderJob
		wantErr bool
	}{
		{
			name: "Simple Test",
			args: args{
				templateRoot:   simpleTestFolder,
				outRoot:        "/tmp/out/",
				defaultContext: pongo2.Context{},
			},
			want: []renderJob{
				renderJob{
					inFile:  simpleTestFolder + "/test",
					outFile: "/tmp/out/test",
					context: pongo2.Context{},
				},
			},
			wantErr: false,
		},
		{
			name: "More Elaborate Test",
			args: args{
				templateRoot:   moreElaborateTestFolder,
				outRoot:        "/tmp/out/",
				defaultContext: pongo2.Context{},
			},
			want: []renderJob{
				renderJob{
					inFile:  moreElaborateTestFolder + "/bla.j2",
					outFile: "/tmp/out/bla.j2",
					context: pongo2.Context{},
				},
				renderJob{
					inFile:  moreElaborateTestFolder + "/test2.xml",
					outFile: "/tmp/out/test2.xml",
					context: pongo2.Context{},
				},
				renderJob{
					inFile:  moreElaborateTestFolder + "/folder/test3.yml",
					outFile: "/tmp/out/folder/test3.yml",
					context: pongo2.Context{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := preprocessTemplateRoot(tt.args.templateRoot, tt.args.outRoot, tt.args.defaultContext)
			if (err != nil) != tt.wantErr {
				t.Errorf("preprocessTemplateRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("preprocessTemplateRoot(): Length doesn't match. Got %v, want %v", len(got), len(tt.want))
			}

			for _, wanted := range tt.want {
				found := false

				for _, gotted := range got {
					if reflect.DeepEqual(gotted, wanted) {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("preprocessTemplateRoot(): Element Not Found in Output: %v, \n got: %v", wanted, got)
				}
			}

		})
	}
}

func Test_convertToStringArray(t *testing.T) {
	context := pongo2.Context{
		"one":  []string{"1"},
		"five": []string{"1", "2", "3", "4", "5"},
		"no":   "Nope",
	}
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				in: context["one"],
			},
			want:    []string{"1"},
			wantErr: false,
		},
		{
			name: "five",
			args: args{
				in: context["five"],
			},
			want:    []string{"1", "2", "3", "4", "5"},
			wantErr: false,
		},
		{
			name: "no",
			args: args{
				in: context["no"],
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToStringArray(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToStringArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
