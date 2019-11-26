package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type SetterConfig struct {
	GithubToken        string                      `yaml:"GithubToken"`
	SchemesMasterURL   string                      `yaml:"SchemesMasterURL"`
	TemplatesMasterURL string                      `yaml:"TemplatesMasterURL"`
	SchemesListFile    string                      `yaml:"SchemesListFile"`
	TemplatesListFile  string                      `yaml:"TemplatesListFile"`
	SchemesCachePath   string                      `yaml:"SchemesCachePath"`
	TemplatesCachePath string                      `yaml:"TemplatesCachePath"`
	DryRun             bool                        `yaml:"DryRun"`
	Colorscheme        string                      `yaml:"Colorscheme"`
	Applications       map[string]StetterAppConfig `yaml:"applications"`
}

type StetterAppConfig struct {
	Enabled bool              `yaml:"enabled"`
	Hook    string            `yaml:"hook"`
	Mode    string            `yaml:"mode"`
	Files   map[string]string `yaml:"files"`
}

func NewConfig(path string) SetterConfig {
	var conf SetterConfig
	file, err := ioutil.ReadFile(path)
	if err != nil {

		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	check(err)
	err = yaml.Unmarshal((file), &conf)
	check(err)
	return conf
}
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

type Application1 struct {
}
