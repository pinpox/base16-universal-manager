package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type Base16Colorscheme struct {
	Name   string `yaml:"scheme"`
	Author string `yaml:"author"`

	Color00 string `yaml:"base00"`
	Color01 string `yaml:"base01"`
	Color02 string `yaml:"base02"`
	Color03 string `yaml:"base03"`
	Color04 string `yaml:"base04"`
	Color05 string `yaml:"base05"`
	Color06 string `yaml:"base06"`
	Color07 string `yaml:"base07"`
	Color08 string `yaml:"base08"`
	Color09 string `yaml:"base09"`
	Color10 string `yaml:"base0A"`
	Color11 string `yaml:"base0B"`
	Color12 string `yaml:"base0C"`
	Color13 string `yaml:"base0D"`
	Color14 string `yaml:"base0E"`
	Color15 string `yaml:"base0F"`
}

func NewBase16Colorscheme(yaml string) Base16Colorscheme {
	return Base16Colorscheme{}
}

func (c *Base16Colorscheme) getColors(url string) *Base16Colorscheme {

	yamlFile, err := DownloadFileToStirng(url)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

//GetBase16Colorscheme returns a Base16Colors strunct containing all colors of
//a given colorscheme
func GetBase16Colorscheme(name string) (Base16Colorscheme, error) {
	//Get the colors
	colorsURL := "https://raw.githubusercontent.com/atelierbram/base16-atelier-schemes/master/atelier-cave-light.yaml"
	var base16Colorscheme Base16Colorscheme
	base16Colorscheme.getColors(colorsURL)
	return base16Colorscheme, nil

}

type Base16ColorschemeList struct {
	colorschemes map[string]string
}

func (c *Base16ColorschemeList) UpdateFromRemote(url ...string) {

	//Set the source for templates
	var masterRepo string
	if len(url) == 1 {
		masterRepo = url[0]
	} else {
		masterRepo = "https://raw.githubusercontent.com/chriskempson/base16-schemes-source/master/list.yaml"
	}

	//Get all repos from master source
	schemeRepos := make(map[string]string)

	schemesYAML, err := DownloadFileToStirng(masterRepo)
	check(err)

	err = yaml.Unmarshal([]byte(schemesYAML), &schemeRepos)
	check(err)

	fmt.Println("Found colorscheme repos: ", len(schemeRepos))

	for k, v := range schemeRepos {
		schemeRepos[k] = v
	}

	c.colorschemes = make(map[string]string)

	// c.colorschemes[k] = v

	fmt.Println("Found colorschemes: ", len(c.colorschemes))

}

func (c *Base16ColorschemeList) Find(input string) (Base16Colorscheme, error) {
	return Base16Colorscheme{}, nil
}
