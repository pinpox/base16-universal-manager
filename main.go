package main

import (
	"fmt"
	// "gopkg.in/yaml.v2"
	"os"
	"text/template"
)

var schemesSourceURL = "https://raw.githubusercontent.com/chriskempson/base16-schemes-source/master/list.yaml"

func main() {
	/*
		var schemesRepos map[string]string
		var templates map[string]string

		schemesYAML, err := DownloadFileToStirng(schemesSourceURL)
		check(err)
		templatesYAML, err := DownloadFileToStirng(templatesSourceURL)
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


	*/

	///////////////////// Start here
	// TODO Get this two vars from flags
	userInputThemeName := "atelier-cave-light"
	userInputTemplateName := "i3"

	var schemeList Base16ColorschemeList
	var templateList Base16TemplateList

	schemeList.UpdateFromRemote()
	templateList.UpdateFromRemote()

	templ, err := templateList.Find(userInputTemplateName)
	check(err)

	scheme, err := schemeList.Find(userInputThemeName)
	check(err)

	Render(templ, scheme)

}

func Render(templ Base16Template, scheme Base16Colorscheme) {

	////Get a stuct containing all 16 colors
	//Base16ThemeColors, err := GetBase16Colorscheme(userInputThemeName)
	//check(err)
	//fmt.Println("Found colorscheme: " + Base16ThemeColors.Name)
	//fmt.Println(Base16ThemeColors) //TODO remove this

	////Get the template as string
	//Base16Template, err := GetBase16Template(userInputTemplateName)
	//check(err)
	fmt.Println("Rendering template: " + templ.Name + " with colorscheme: " + scheme.Name)

	//Render the template to Stdout
	// TODO use mustache themes instead
	t, err := template.New(templ.Name).Parse(templ.Template)
	check(err)
	err = t.Execute(os.Stdout, scheme)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}

}
