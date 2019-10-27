package templating

import (
	"reflect"
	"testing"

	"github.com/flosch/pongo2"
)

func RenderJobSlicesMustBeEqual(one []renderJob, two []renderJob, t *testing.T) {
	if one == nil && two == nil {
		return
	}
	if (one == nil) != (two == nil) {
		t.Errorf("One of them is Nil: One: %v , Two: %v", one, two)
	}

	for _, elFirst := range one {
		found := false

		for _, elTwo := range two {
			if reflect.DeepEqual(elFirst, elTwo) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Can not find Element %v", elFirst)
		}
	}
	return
}

func Test_isProcessingDone(t *testing.T) {
	type args struct {
		jobs []renderJob
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Single Job Done",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/path/1",
						outFile: "/path/1",
					},
				},
			},
			want: true,
		},
		{
			name: "Test Single Job NotDone",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/__for_blub_in_blubs__/1",
						outFile: "/__for_blub_in_blubs__/1",
					},
				},
			},
			want: false,
		},
		{
			name: "Test Multi Job NotDone",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/blabla/1",
						outFile: "/blabla/1",
					},
					renderJob{
						inFile:  "/blublub/1",
						outFile: "/blublub/1",
					},
					renderJob{
						inFile:  "/__for_blub_in_blubs__/1",
						outFile: "/__for_blub_in_blubs__/1",
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isProcessingDone(tt.args.jobs); got != tt.want {
				t.Errorf("isProcessingDone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitJobSingle(t *testing.T) {
	noEnvContext := pongo2.Context{}
	oneEnvContext := pongo2.Context{
		"envs": []string{"one"},
	}
	twoEnvContext := pongo2.Context{
		"envs": []string{"one", "two"},
	}
	//	threeEnvContext := pongo2.Context{
	//		"envs": []interface{}{"one", 2, 3.0},
	//	}
	//	eightEnvContext := pongo2.Context{
	//		"envs": []string{"1", "2", "3", "4", "5", "6", "7", "8"},
	//	}
	type args struct {
		job renderJob
	}
	tests := []struct {
		name    string
		args    args
		want    []renderJob
		wantErr bool
	}{
		{
			name: "NoChange",
			args: args{
				job: renderJob{
					inFile:  "/one/two",
					outFile: "/tmp/one/two",
					context: noEnvContext,
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/one/two",
					outFile: "/tmp/one/two",
					context: noEnvContext,
				}},
			wantErr: false,
		},
		{
			name: "OneEnvOneSubst",
			args: args{
				job: renderJob{
					inFile:  "/one/__for_env_in_envs__/blub",
					outFile: "/tmp/one/__for_env_in_envs__/blub",
					context: oneEnvContext,
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/one/__for_env_in_envs__/blub",
					outFile: "/tmp/one/one/blub",
					context: pongo2.Context{
						"envs": []string{"one"},
						"env":  "one",
					},
				}},
			wantErr: false,
		},
		{
			name: "TwoEnvOneSubst",
			args: args{
				job: renderJob{
					inFile:  "/one/__for_env_in_envs__/blub",
					outFile: "/tmp/one/__for_env_in_envs__/blub",
					context: twoEnvContext,
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/one/__for_env_in_envs__/blub",
					outFile: "/tmp/one/one/blub",
					context: pongo2.Context{
						"envs": []string{"one", "two"},
						"env":  "one",
					},
				},
				renderJob{
					inFile:  "/one/__for_env_in_envs__/blub",
					outFile: "/tmp/one/two/blub",
					context: pongo2.Context{
						"envs": []string{"one", "two"},
						"env":  "two",
					},
				}},
			wantErr: false,
		},
		{
			name: "TwoEnvTwoSubst",
			args: args{
				job: renderJob{
					inFile:  "/one/__for_env_in_envs__/__for_env_in_envs__/blub",
					outFile: "/tmp/one/__for_env_in_envs__/__for_env_in_envs__/blub",
					context: twoEnvContext,
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/one/__for_env_in_envs__/__for_env_in_envs__/blub",
					outFile: "/tmp/one/one/__for_env_in_envs__/blub",
					context: pongo2.Context{
						"envs": []string{"one", "two"},
						"env":  "one",
					},
				},
				renderJob{
					inFile:  "/one/__for_env_in_envs__/__for_env_in_envs__/blub",
					outFile: "/tmp/one/two/__for_env_in_envs__/blub",
					context: pongo2.Context{
						"envs": []string{"one", "two"},
						"env":  "two",
					},
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitJob(tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitJobSingle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitJobSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postProcessJobs(t *testing.T) {
	simpleContext := pongo2.Context{
		"blub": "blub",
	}
	oneSubstContext := pongo2.Context{
		"blub": "blub",
		"envs": []string{"test"},
	}

	threeEnvsContext := pongo2.Context{
		"blub": "blub",
		"envs": []string{"dev", "test", "prod"},
	}

	twoSubstContext := pongo2.Context{
		"envs":  []string{"dev", "test"},
		"sites": []string{"home", "away"},
	}
	type args struct {
		jobs []renderJob
	}
	tests := []struct {
		name      string
		args      args
		want      []renderJob
		wantError bool
	}{
		{
			name: "Check Order",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/template/one",
						outFile: "/tmp/one",
						context: simpleContext,
					},
					renderJob{
						inFile:  "/template/two",
						outFile: "/tmp/two",
						context: simpleContext,
					},
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/template/one",
					outFile: "/tmp/one",
					context: simpleContext,
				},
				renderJob{
					inFile:  "/template/two",
					outFile: "/tmp/two",
					context: simpleContext,
				},
			},
			wantError: false,
		},
		{
			name: "One Substitution",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/template/__for_env_in_envs__/one",
						outFile: "/tmp/__for_env_in_envs__/one",
						context: oneSubstContext,
					},
					renderJob{
						inFile:  "/template/two",
						outFile: "/tmp/two",
						context: oneSubstContext,
					},
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/template/two",
					outFile: "/tmp/two",
					context: oneSubstContext,
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/one",
					outFile: "/tmp/test/one",
					context: pongo2.Context{
						"blub": "blub",
						"envs": []string{"test"},
						"env":  "test",
					},
				},
			},
			wantError: false,
		},
		{
			name: "Three Env Context One Subst",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/template/__for_env_in_envs__/one",
						outFile: "/tmp/__for_env_in_envs__/one",
						context: threeEnvsContext,
					},
					renderJob{
						inFile:  "/template/two",
						outFile: "/tmp/two",
						context: threeEnvsContext,
					},
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/template/two",
					outFile: "/tmp/two",
					context: threeEnvsContext,
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/one",
					outFile: "/tmp/dev/one",
					context: pongo2.Context{
						"blub": "blub",
						"envs": []string{"dev", "test", "prod"},
						"env":  "dev",
					},
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/one",
					outFile: "/tmp/test/one",
					context: pongo2.Context{
						"blub": "blub",
						"envs": []string{"dev", "test", "prod"},
						"env":  "test",
					},
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/one",
					outFile: "/tmp/prod/one",
					context: pongo2.Context{
						"blub": "blub",
						"envs": []string{"dev", "test", "prod"},
						"env":  "prod",
					},
				},
			},
			wantError: false,
		},
		{
			name: "Two Substitutions",
			args: args{
				jobs: []renderJob{
					renderJob{
						inFile:  "/template/__for_env_in_envs__/__for_site_in_sites__.html",
						outFile: "/tmp/__for_env_in_envs__/__for_site_in_sites__.html",
						context: twoSubstContext,
					},
				},
			},
			want: []renderJob{
				renderJob{
					inFile:  "/template/__for_env_in_envs__/__for_site_in_sites__.html",
					outFile: "/tmp/dev/home.html",
					context: pongo2.Context{
						"envs":  []string{"dev", "test"},
						"sites": []string{"home", "away"},
						"env":   "dev",
						"site":  "home",
					},
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/__for_site_in_sites__.html",
					outFile: "/tmp/dev/away.html",
					context: pongo2.Context{
						"envs":  []string{"dev", "test"},
						"sites": []string{"home", "away"},
						"env":   "dev",
						"site":  "away",
					},
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/__for_site_in_sites__.html",
					outFile: "/tmp/test/home.html",
					context: pongo2.Context{
						"envs":  []string{"dev", "test"},
						"sites": []string{"home", "away"},
						"env":   "test",
						"site":  "home",
					},
				},
				renderJob{
					inFile:  "/template/__for_env_in_envs__/__for_site_in_sites__.html",
					outFile: "/tmp/test/away.html",
					context: pongo2.Context{
						"envs":  []string{"dev", "test"},
						"sites": []string{"home", "away"},
						"env":   "test",
						"site":  "away",
					},
				},
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errors := postProcessJobs(tt.args.jobs)

			if (errors == nil) == tt.wantError {
				t.Errorf("Error values didn't Match")
			}

			RenderJobSlicesMustBeEqual(got, tt.want, t)

		})
	}
}
