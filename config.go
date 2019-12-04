package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// SetterConfig is the applicaton's configuration.
type SetterConfig struct {
	GithubToken        string                     `yaml:"GithubToken"`
	SchemesMasterURL   string                     `yaml:"SchemesMasterURL"`
	TemplatesMasterURL string                     `yaml:"TemplatesMasterURL"`
	DryRun             bool                       `yaml:"DryRun"`
	Colorscheme        string                     `yaml:"Colorscheme"`
	Applications       map[string]SetterAppConfig `yaml:"applications"`
	SchemesCachePath   string
	SchemesListFile    string
	TemplatesCachePath string
	TemplatesListFile  string
}

// SetterAppConfig is the configuration for a particular application being themed.
type SetterAppConfig struct {
	Enabled bool              `yaml:"enabled"`
	Hook    string            `yaml:"hook"`
	Mode    string            `yaml:"mode"`
	Files   map[string]string `yaml:"files"`
}

// NewConfig parses the provided configuration file and returns the app configuration.
func NewConfig(path string) SetterConfig {
	if path == "" {
		fmt.Fprintf(os.Stderr, "no config file found\n")
		os.Exit(1)
	}

	var conf SetterConfig
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	check(err)
	err = yaml.Unmarshal((file), &conf)
	check(err)

	conf.SchemesCachePath = filepath.Join(xdgDirs.CacheHome(), "schemes")
	conf.SchemesListFile = filepath.Join(xdgDirs.CacheHome(), "schemeslist.yaml")
	conf.TemplatesCachePath = filepath.Join(xdgDirs.CacheHome(), "templates")
	conf.TemplatesListFile = filepath.Join(xdgDirs.CacheHome(), "templateslist.yaml")

	return conf
}

// Show prints the app configuration
func (c SetterConfig) Show() {
	fmt.Println("GithubToken: ", c.GithubToken)
	fmt.Println("SchemesListFile: ", c.SchemesListFile)
	fmt.Println("TemplatesListFile: ", c.TemplatesListFile)
	fmt.Println("SchemesCachePath: ", c.SchemesCachePath)
	fmt.Println("TemplatesCachePath: ", c.TemplatesCachePath)
	fmt.Println("DryRun: ", c.DryRun)

	for app, appConfig := range c.Applications {
		fmt.Println("  App: ", app)
		fmt.Println("    Enabled: ", appConfig.Enabled)
		fmt.Println("    Hook: ", appConfig.Hook)
		for k, v := range appConfig.Files {
			fmt.Println("      ", k, "  ", v)
		}
	}
}
