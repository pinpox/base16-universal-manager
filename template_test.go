package main

import (
	"reflect"
	"testing"
)

func TestBase16TemplateList_GetRawBaseURL(t *testing.T) {
	type args struct {
		url        string
		mainBranch string
	}
	type result struct {
		result string
		error  string
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{
			"sourcehut",
			args{"https://git.sr.ht/foo/bar.git", "master"},
			result{"https://git.sr.ht/foo/bar.git/blob/master/", ""},
		},
		{
			"github",
			args{"https://github.com/foo/bar.git", "master"},
			result{"https://raw.githubusercontent.com/foo/bar.git/master/", ""},
		},
		{
			"gitlab",
			args{"https://gitlab.com/foo/bar.git", "master"},
			result{"https://gitlab.com/foo/bar.git/-/raw/master/", ""},
		},
		{
			"unsupported",
			args{"https://baz.com/foo/bar.git", "master"},
			result{"", "git host \"baz.com\" for \"foo/bar.git\" not supported yet"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRawBaseURL(tt.args.url, tt.args.mainBranch)
			if got != tt.want.result {
				t.Errorf("GetRawBaseURL() result = %v, want %v", got, tt.want.result)
			}
			if err != nil && err.Error() != tt.want.error {
				t.Errorf("GetRawBaseURL() error = '%v', want '%v'", err, tt.want.error)
			}
		})
	}
}

func TestBase16TemplateList_GetBase16Template(t *testing.T) {
	type fields struct {
		templates map[string]string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base16Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Base16TemplateList{
				templates: tt.fields.templates,
			}
			got, err := l.GetBase16Template(tt.args.name, "master")
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base16TemplateList.GetBase16Template() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase16TemplateList_UpdateTemplates(t *testing.T) {
	type fields struct {
		templates map[string]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Base16TemplateList{
				templates: tt.fields.templates,
			}
			l.UpdateTemplates()
		})
	}
}

func TestLoadBase16TemplateList(t *testing.T) {
	tests := []struct {
		name string
		want Base16TemplateList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadBase16TemplateList()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadBase16TemplateList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveBase16TemplateList(t *testing.T) {
	type args struct {
		l Base16TemplateList
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveBase16TemplateList(tt.args.l); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestBase16TemplateList_Find(t *testing.T) {
	type fields struct {
		templates map[string]string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base16Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Base16TemplateList{
				templates: tt.fields.templates,
			}
			got, err := c.Find(tt.args.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base16TemplateList.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
