package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	defaultSchemesMasterURL   = "https://raw.githubusercontent.com/chriskempson/base16-schemes-source/master/list.yaml"
	defaultTemplatesMasterURL = "https://raw.githubusercontent.com/chriskempson/base16-templates-source/master/list.yaml"
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
	Enabled             bool                  `yaml:"enabled"`
	Hook                string                `yaml:"hook"`
	Template            string                `yaml:"template"`
	Files               map[string]FileConfig `yaml:"files"`
	DefaultRemoteBranch string                `yaml:"remote-branch"`
}

// FileConfig is the configuration for how a particular file should be rendered
type FileConfig struct {
	Path        string `yaml:"path"`
	Mode        string `yaml:"mode"`
	StartMarker string `yaml:"start_marker"`
	EndMarker   string `yaml:"end_marker"`
}

// NewConfig parses the provided configuration file and returns the app configuration.
func NewConfig(path string) (SetterConfig, error) {
	if path == "" {
		return SetterConfig{}, errors.New("no config file specified")
	}

	var conf SetterConfig
	file, err := os.ReadFile(path)
	if err != nil {
		return SetterConfig{}, fmt.Errorf("reading config file: %w", err)
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return SetterConfig{}, fmt.Errorf("parsing YAML from config file %s: %w", path, err)
	}

	if conf.SchemesMasterURL == "" {
		conf.SchemesMasterURL = defaultSchemesMasterURL
	}
	if conf.TemplatesMasterURL == "" {
		conf.TemplatesMasterURL = defaultTemplatesMasterURL
	}

	conf.SchemesCachePath = filepath.Join(xdgDirs.CacheHome(), "schemes")
	conf.SchemesListFile = filepath.Join(xdgDirs.CacheHome(), "schemeslist.yaml")
	conf.TemplatesCachePath = filepath.Join(xdgDirs.CacheHome(), "templates")
	conf.TemplatesListFile = filepath.Join(xdgDirs.CacheHome(), "templateslist.yaml")

	return conf, nil
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
		fmt.Println("    Default remote branch: ", appConfig.DefaultRemoteBranch)
		fmt.Println("    Hook: ", appConfig.Hook)
		for k, v := range appConfig.Files {
			fmt.Println("      ", k, "  ", v)
		}
	}
}
