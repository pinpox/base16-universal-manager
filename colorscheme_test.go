package main

import (
	"reflect"
	"testing"
)

func TestBase16Colorscheme_MustacheContext(t *testing.T) {
	type fields struct {
		Name       string
		Author     string
		Color00    string
		Color01    string
		Color02    string
		Color03    string
		Color04    string
		Color05    string
		Color06    string
		Color07    string
		Color08    string
		Color09    string
		Color10    string
		Color11    string
		Color12    string
		Color13    string
		Color14    string
		Color15    string
		RepoURL    string
		RawBaseURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Base16Colorscheme{
				Name:       tt.fields.Name,
				Author:     tt.fields.Author,
				Color00:    tt.fields.Color00,
				Color01:    tt.fields.Color01,
				Color02:    tt.fields.Color02,
				Color03:    tt.fields.Color03,
				Color04:    tt.fields.Color04,
				Color05:    tt.fields.Color05,
				Color06:    tt.fields.Color06,
				Color07:    tt.fields.Color07,
				Color08:    tt.fields.Color08,
				Color09:    tt.fields.Color09,
				Color10:    tt.fields.Color10,
				Color11:    tt.fields.Color11,
				Color12:    tt.fields.Color12,
				Color13:    tt.fields.Color13,
				Color14:    tt.fields.Color14,
				Color15:    tt.fields.Color15,
				RepoURL:    tt.fields.RepoURL,
				RawBaseURL: tt.fields.RawBaseURL,
			}
			if got := s.MustacheContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base16Colorscheme.MustacheContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase16ColorschemeList_GetBase16Colorscheme(t *testing.T) {
	type fields struct {
		colorschemes map[string]string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Base16Colorscheme
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Base16ColorschemeList{
				colorschemes: tt.fields.colorschemes,
			}
			got, err := l.GetBase16Colorscheme(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base16ColorschemeList.GetBase16Colorscheme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base16ColorschemeList.GetBase16Colorscheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBase16Colorscheme(t *testing.T) {
	type args struct {
		yamlData string
	}
	tests := []struct {
		name string
		args args
		want Base16Colorscheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBase16Colorscheme(tt.args.yamlData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBase16Colorscheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadBase16ColorschemeList(t *testing.T) {
	tests := []struct {
		name string
		want Base16ColorschemeList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadBase16ColorschemeList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadBase16ColorschemeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveBase16ColorschemeList(t *testing.T) {
	type args struct {
		l Base16ColorschemeList
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveBase16ColorschemeList(tt.args.l)
		})
	}
}

func TestBase16ColorschemeList_UpdateSchemes(t *testing.T) {
	type fields struct {
		colorschemes map[string]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Base16ColorschemeList{
				colorschemes: tt.fields.colorschemes,
			}
			l.UpdateSchemes()
		})
	}
}

func TestBase16ColorschemeList_Find(t *testing.T) {
	type fields struct {
		colorschemes map[string]string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base16Colorscheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Base16ColorschemeList{
				colorschemes: tt.fields.colorschemes,
			}
			if got := c.Find(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base16ColorschemeList.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
