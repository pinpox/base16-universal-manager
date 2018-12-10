package main

import (
	"reflect"
	"testing"
)

func TestDownloadFileToStirng(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadFileToStirng(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFileToStirng() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadFileToStirng() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findYAMLinRepo(t *testing.T) {
	type args struct {
		repoURL string
	}
	tests := []struct {
		name string
		args args
		want []GitHubFile
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findYAMLinRepo(tt.args.repoURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findYAMLinRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadStringMap(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadStringMap(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveStringMap(t *testing.T) {
	type args struct {
		data map[string]string
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveStringMap(tt.args.data, tt.args.path)
		})
	}
}

func TestFindMatchInMap(t *testing.T) {
	type args struct {
		choices map[string]string
		input   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMatchInMap(tt.args.choices, tt.args.input); got != tt.want {
				t.Errorf("FindMatchInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exe_cmd(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exe_cmd(tt.args.cmd)
		})
	}
}
