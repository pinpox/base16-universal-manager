package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestDownloadFileToString(t *testing.T) {
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
			got, err := DownloadFileToString(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFileToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadFileToString() = %v, want %v", got, tt.want)
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
		data []byte
		want string
	}{
		{
			name: "Write simple string",
			path: "testdata/output1",
			data: []byte("A simple string"),
			want: "testdata/writefile/expect1",
		},
		{
			name: "Write emtpy string",
			path: "testdata/output2",
			data: []byte(""),
			want: "testdata/writefile/expect2",
		},
		{
			name: "Write string with linebreaks",
			path: "testdata/output3",
			data: []byte("A string\nwith \nlinebreaks"),
			want: "testdata/writefile/expect3",
		},
		{
			name: "Re-Write string with linebreaks",
			path: "testdata/output4",
			data: []byte("A string\nwith \nlinebreaks"),
			want: "testdata/writefile/expect3",
		},
	}
	for _, tt := range tests {
		os.Remove(tt.path)
		t.Run(tt.name, func(t *testing.T) {
			err := WriteFile(tt.path, tt.data)
			if err != nil {
				t.Errorf("Error during WriteFile() %q: %q", tt.name, err.Error())
			}
			if !deepCompareFiles(tt.path, tt.want) {
				t.Errorf("WriteFile() %q files differ", tt.name)
			}
		})
	}
}

func TestReplaceMultiline(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name        string
		path        string
		replacement string
		blockStart  string
		blockEnd    string
		want        string
	}{
		{
			name:        "full example",
			path:        "testdata/replacefile/input1",
			replacement: "qux\n",
			blockStart:  "beginmarker",
			blockEnd:    "endmarker",
			want:        "testdata/replacefile/expect1",
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simple way to copy input file into one that can be tested with
			// since ReplaceMultiline will overwrite the file in-place
			exe_cmd(fmt.Sprintf("cp %s testdata/output1", tt.path))

			err := ReplaceMultiline("testdata/output1", tt.replacement, tt.blockStart, tt.blockEnd)
			if err != nil {
				t.Errorf("Error during ReplaceMultiline() %q: %q", tt.name, err.Error())
			}
			if !deepCompareFiles("testdata/output1", tt.want) {
				t.Errorf("ReplaceMultiline() %q files differ", tt.name)
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
