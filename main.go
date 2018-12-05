package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
	"text/template"
)

var templatesSourceURL = "https://raw.githubusercontent.com/chriskempson/base16-templates-source/master/list.yaml"
var schemesSourceURL = "https://raw.githubusercontent.com/chriskempson/base16-schemes-source/master/list.yaml"

type GitHubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type GitHubFilesCollection struct {
	Collection []GitHubFile
}

func findYAMLinRepo(repoURL string) []GitHubFile {

	// Get all files from repo
	repoFiles, err := DownloadYAML("https://api.github.com/repos/atelierbram/base16-atelier-schemes/contents/")
	check(err)
	keys := make([]GitHubFile, 0)
	json.Unmarshal([]byte(repoFiles), &keys)

	// Create a list of .yaml files
	var colorSchemes []GitHubFile
	for _, v := range keys {
		re := regexp.MustCompile(".*yaml")
		if re.MatchString(v.Name) {
			colorSchemes = append(colorSchemes, v)
		}
	}
	return colorSchemes
}

func main() {

	userInputTheme := "atelier-cave-light"
	userInputTemplate := "i3"

	var schemesRepos map[string]string
	var templates map[string]string

	schemesYAML, err := DownloadYAML(schemesSourceURL)
	check(err)
	templatesYAML, err := DownloadYAML(templatesSourceURL)
	check(err)
	err = yaml.Unmarshal([]byte(schemesYAML), &schemesRepos)
	check(err)
	err = yaml.Unmarshal([]byte(templatesYAML), &templates)
	check(err)

	fmt.Println("Found templates: ", len(templates))
	for k, v := range templates {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println("Found colorschemes: ", len(schemesRepos))

	//Get the colors
	colorsURL := "https://raw.githubusercontent.com/atelierbram/base16-atelier-schemes/master/atelier-cave-light.yaml"
	var base16Colorscheme B16Colors
	base16Colorscheme.getColors(colorsURL)

	///////////////////// Start here
	B16Theme, err := GetB16Colorscheme(userInputTheme)
	check(err)
	B16Template, err := GetB16Template(userInputTemplate)
	check(err)

	t, err := template.New(userInputTemplate).Parse(string(B16Template))
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, B16Theme)
	if err != nil {
		panic(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
