package main

import (
	"reflect"
	"testing"
)

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
			if got := l.GetBase16Template(tt.args.name); !reflect.DeepEqual(got, tt.want) {
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
			if got := LoadBase16TemplateList(); !reflect.DeepEqual(got, tt.want) {
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
			SaveBase16TemplateList(tt.args.l)
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
			if got := c.Find(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base16TemplateList.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
