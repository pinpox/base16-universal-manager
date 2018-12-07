package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

	RepoURL string

	RawBaseURL string
}

func (s Base16Colorscheme) MustacheContext() map[string]interface{} {
	var bases = []string{
		"00", "01", "02", "03", "04", "05", "06", "07",
		"08", "09", "0A", "0B", "0C", "0D", "0E", "0F",
	}
	slug := "base16-test-slug"
	ret := map[string]interface{}{
		"scheme-name":   s.Name,
		"scheme-author": s.Author,
		//TODO correct this slug
		"scheme-slug": slug,

		"scheme-slug-underscored": strings.Replace(slug, "-", "_", -1),
	}

	for _, base := range bases {
		baseKey := "base" + base

		//TODO
		rVal := "33"
		gVal := "ff"
		bVal := "00"

		rValf, err := strconv.ParseUint(rVal, 16, 32)
		check(err)
		gValf, err := strconv.ParseUint(gVal, 16, 32)
		check(err)
		bValf, err := strconv.ParseUint(bVal, 16, 32)
		check(err)

		ret[baseKey+"-hex"] = rVal + gVal + bVal

		ret[baseKey+"-hex-r"] = rVal
		ret[baseKey+"-hex-g"] = gVal
		ret[baseKey+"-hex-b"] = bVal

		ret[baseKey+"-rgb-r"] = rValf
		ret[baseKey+"-rgb-g"] = gValf
		ret[baseKey+"-rgb-b"] = bValf

		ret[baseKey+"-dec-r"] = rValf / 255
		ret[baseKey+"-dec-g"] = gValf / 255
		ret[baseKey+"-dec-b"] = bValf / 255

	}

	return ret
}

func (l *Base16ColorschemeList) GetBase16Colorscheme(name string) (Base16Colorscheme, error) {

	if len(name) == 0 {
		panic("Colorscheme name was empty")
	}

	path := schemesCachePath + name

	// Create local schemes file, if not present
	if _, err := os.Stat(path); os.IsNotExist(err) {

		parts := strings.Split(l.colorschemes[name], "/")

		yamlURL := "https://raw.githubusercontent.com/" + parts[3] + "/" + parts[4] + "/master/" + parts[7]

		fmt.Println("downloading theme from: ", yamlURL)

		schemeData, err := DownloadFileToStirng(yamlURL)
		check(err)
		saveFile, err := os.Create(path)
		defer saveFile.Close()
		check(err)
		saveFile.Write([]byte(schemeData))
		saveFile.Close()
	}

	colorscheme, err := ioutil.ReadFile(path)
	check(err)

	return NewBase16Colorscheme(string(colorscheme)), err

}

func NewBase16Colorscheme(yamlData string) Base16Colorscheme {
	var scheme Base16Colorscheme

	err := yaml.Unmarshal([]byte(yamlData), &scheme)
	check(err)
	return scheme
}

func LoadBase16ColorschemeList() Base16ColorschemeList {
	colorschemes := LoadStringMap(schemesListFile)
	return Base16ColorschemeList{colorschemes}
}

func SaveBase16ColorschemeList(l Base16ColorschemeList) {
	SaveStringMap(l.colorschemes, schemesListFile)
}

type Base16ColorschemeList struct {
	colorschemes map[string]string
}

func (l *Base16ColorschemeList) UpdateSchemes() {

	//Get all repos from master source
	schemeRepos := make(map[string]string)

	schemesYAML, err := DownloadFileToStirng(schemesSourceURL)
	check(err)

	err = yaml.Unmarshal([]byte(schemesYAML), &schemeRepos)
	check(err)

	fmt.Println("Found colorscheme repos: ", len(schemeRepos))

	for k, v := range schemeRepos {
		schemeRepos[k] = v
	}

	for _, v1 := range schemeRepos {
		fmt.Println("Getting schemes from: " + v1)

		for _, v2 := range findYAMLinRepo(v1) {
			l.colorschemes[v2.Name] = v2.HTMLURL
		}

	}

	fmt.Println("Found colorschemes: ", len(l.colorschemes))
	SaveBase16ColorschemeList(Base16ColorschemeList{l.colorschemes})
}

func (c *Base16ColorschemeList) Find(input string) Base16Colorscheme {

	if _, err := os.Stat(schemesListFile); os.IsNotExist(err) {
		check(err)
		fmt.Println("Colorschemes list not found, pulling new one...")
		c.UpdateSchemes()
	}

	if len(c.colorschemes) == 0 {
		fmt.Println("No templates in list, pulling new one... ")
		c.UpdateSchemes()
	}

	colorschemeName := FindMatchInMap(c.colorschemes, input)
	scheme, err := c.GetBase16Colorscheme(colorschemeName)
	check(err)
	return scheme
}
