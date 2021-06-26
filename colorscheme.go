package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
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
	FileName   string
}

func (s Base16Colorscheme) MustacheContext(ext string) map[string]interface{} {
	var bases = map[string]string{
		"00": s.Color00,
		"01": s.Color01,
		"02": s.Color02,
		"03": s.Color03,
		"04": s.Color04,
		"05": s.Color05,
		"06": s.Color06,
		"07": s.Color07,
		"08": s.Color08,
		"09": s.Color09,
		"0A": s.Color10,
		"0B": s.Color11,
		"0C": s.Color12,
		"0D": s.Color13,
		"0E": s.Color14,
		"0F": s.Color15,
	}
	slug := strings.Replace(strings.ToLower(s.FileName), " ", "-", -1)
	ret := map[string]interface{}{
		"scheme-name":             s.Name,
		"scheme-author":           s.Author,
		"scheme-slug":             slug,
		"scheme-slug-underscored": strings.Replace(slug, "-", "_", -1),
	}

	for base, color := range bases {
		baseKey := "base" + base

		//TODO
		rVal := color[:2]
		gVal := color[2:4]
		bVal := color[4:6]

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

	schemePath := path.Join(appConf.SchemesCachePath, name)

	parts := strings.Split(l.colorschemes[name], "/")
	yamlURL := strings.Join([]string{"https://raw.githubusercontent.com", parts[3], parts[4], parts[6], parts[7]}, "/")
	// Create local schemes file, if not present
	if _, err := os.Stat(schemePath); os.IsNotExist(err) {

		fmt.Println("downloading theme from: ", yamlURL)

		schemeData, err := DownloadFileToString(yamlURL)
		check(err)
		saveFile, err := os.Create(schemePath)
		//TODO delete old file?
		defer saveFile.Close()
		check(err)
		saveFile.Write([]byte(schemeData))
		saveFile.Close()
	}

	colorscheme, err := ioutil.ReadFile(schemePath)
	check(err)

	scheme := NewBase16Colorscheme(string(colorscheme))
	scheme.RawBaseURL = yamlURL
	scheme.FileName = parts[7][:len(parts[7])-5]
	return scheme, err

}

func NewBase16Colorscheme(yamlData string) Base16Colorscheme {
	var scheme Base16Colorscheme

	err := yaml.Unmarshal([]byte(yamlData), &scheme)
	check(err)
	return scheme
}

func LoadBase16ColorschemeList() Base16ColorschemeList {
	colorschemes := LoadStringMap(appConf.SchemesListFile)
	return Base16ColorschemeList{colorschemes}
}

func SaveBase16ColorschemeList(l Base16ColorschemeList) {
	SaveStringMap(l.colorschemes, appConf.SchemesListFile)
}

type Base16ColorschemeList struct {
	colorschemes map[string]string
}

func (l *Base16ColorschemeList) UpdateSchemes() {

	//Get all repos from master source
	schemeRepos := make(map[string]string)

	schemesYAML, err := DownloadFileToString(appConf.SchemesMasterURL)
	check(err)

	err = yaml.Unmarshal([]byte(schemesYAML), &schemeRepos)
	check(err)

	fmt.Println("Found colorscheme repos: ", len(schemeRepos))

	for k, v := range schemeRepos {
		schemeRepos[k] = v
	}

	color := strings.Split(appConf.Colorscheme, "-")[0]

	for _, v1 := range schemeRepos {
		color_url := strings.Join(strings.Split(v1, "-"), " ")
		if strings.Contains(color_url, color) {
			fmt.Println("Getting schemes from: " + v1)

			for _, v2 := range findYAMLinRepo(v1) {
				l.colorschemes[v2.Name] = v2.HTMLURL
			}

		}
	}

	fmt.Println("Found colorschemes: ", len(l.colorschemes))
	SaveBase16ColorschemeList(Base16ColorschemeList{l.colorschemes})
}

func (c *Base16ColorschemeList) Find(input string) Base16Colorscheme {

	if _, err := os.Stat(appConf.SchemesListFile); os.IsNotExist(err) {
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
