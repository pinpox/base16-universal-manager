package main

import (
	"os"
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

func TestWriteFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		data string
		want string
	}{
		{"Write simple string", "testdata/output1", "A simple string", "testdata/writefile/expect1"},
		{"Write emtpy string", "testdata/output2", "", "testdata/writefile/expect2"},
		{"Write string with linebreaks", "testdata/output3", "A string\nwith \nlinebreaks", "testdata/writefile/expect3"},
		{"Re-Write string with linebreaks", "testdata/output4", "A string\nwith \nlinebreaks", "testdata/writefile/expect3"},
	}
	for _, tt := range tests {
		os.Remove(tt.path)
		t.Run(tt.name, func(t *testing.T) {
			WriteFile(tt.path, tt.data)
			if !deepCompareFiles(tt.path, tt.want) {
				t.Errorf("WriteFile() \"%v\" files differ", tt.name)
			}
		})
	}
}

func TestAppendFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		data string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppendFile(tt.path, tt.data)
		})
	}
}

func TestReplaceMultiline(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name        string
		input       string
		replacement string
		blockStart  string
		blockEnd    string
		want        string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMultiline(tt.input, tt.replacement, tt.blockStart, tt.blockEnd); got != tt.want {
				t.Errorf("ReplaceMultiline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deepCompareFiles(t *testing.T) {
	tests := []struct {
		name  string
		file1 string
		file2 string
		want  bool
	}{
		{"Two identical files", "./testdata/compare/fileA1equal", "./testdata/compare/fileA2equal", true},
		{"Two differing compare/files", "./testdata/compare/fileB1diff", "./testdata/compare/fileB2diff", false},
		{"Two emtpy files", "./testdata/compare/fileC1empty", "./testdata/compare/fileC2empty", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepCompareFiles(tt.file1, tt.file2); got != tt.want {
				t.Errorf("deepCompareFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
