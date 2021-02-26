package main

import "testing"

func Test_main(t *testing.T) {
	/*
		tests := []struct {
			name string
		}{
			// TODO: Add test cases.
		}
			for range tests {
				t.Run(tt.name, func(t *testing.T) {
					main()
				})
			}
	*/
}

func TestBase16Render(t *testing.T) {
	type args struct {
		templ  Base16Template
		scheme Base16Colorscheme
		app string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Base16Render(tt.args.templ, tt.args.scheme, tt.args.app)
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.e)
		})
	}
}
