package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want SetterConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetterConfig_Show(t *testing.T) {
	type fields struct {
		GithubToken        string
		SchemesMasterURL   string
		TemplatesMasterURL string
		SchemesListFile    string
		TemplatesListFile  string
		SchemesCachePath   string
		TemplatesCachePath string
		DryRun             bool
		Colorscheme        string
		Applications       map[string]SetterAppConfig
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := SetterConfig{
				GithubToken:        tt.fields.GithubToken,
				SchemesMasterURL:   tt.fields.SchemesMasterURL,
				TemplatesMasterURL: tt.fields.TemplatesMasterURL,
				SchemesListFile:    tt.fields.SchemesListFile,
				TemplatesListFile:  tt.fields.TemplatesListFile,
				SchemesCachePath:   tt.fields.SchemesCachePath,
				TemplatesCachePath: tt.fields.TemplatesCachePath,
				DryRun:             tt.fields.DryRun,
				Colorscheme:        tt.fields.Colorscheme,
				Applications:       tt.fields.Applications,
			}
			c.Show()
		})
	}
}

func TestDefaultMasterURLs(t *testing.T) {
	// Create a temporary, empty config file
	tmpFile, err := ioutil.TempFile(os.TempDir(), "base16-universal-manager-")
	if err != nil {
		t.Fatalf("Cannot create temporary file\n")
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	config := NewConfig(tmpFile.Name())
	if config.SchemesMasterURL == "" {
		t.Fatalf("SchemesMasterURL should default to %s\n", defaultSchemesMasterURL)
	}
	if config.TemplatesMasterURL == "" {
		t.Fatalf("TemplatesMasterURL should default to %s\n", defaultTemplatesMasterURL)
	}
}
