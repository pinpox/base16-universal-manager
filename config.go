package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SetterConfig struct {
	GithubToken        string                      `yaml:"GithubToken"`
	SchemesListFile    string                      `yaml:"SchemesListFile"`
	TemplatesListFile  string                      `yaml:"TemplatesListFile"`
	SchemesCachePath   string                      `yaml:"SchemesCachePath"`
	TemplatesCachePath string                      `yaml:"TemplatesCachePath"`
	DryRun             bool                        `yaml:"DryRun"`
	Applications       map[string]StetterAppConfig `yaml:"applications"`
}

type StetterAppConfig struct {
	Enabled bool              `yaml:"enabled"`
	Hook    string            `yaml:"hook"`
	Files   map[string]string `yaml:"files"`
}

func NewConfig(path string) SetterConfig {
	var conf SetterConfig
	file, err := ioutil.ReadFile(path)
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

	for k, v := range c.Applications {
		fmt.Println("  App: ", k)
		fmt.Println("    Enabled: ", v.Enabled)
		fmt.Println("    Hook: ", v.Hook)
		for k1, v1 := range v.Files {
			fmt.Println("      ", k1, "  ", v1)
		}
	}
}

type Application1 struct {
}
