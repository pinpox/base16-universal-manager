package main

import (
	"fmt"
	// "gopkg.in/yaml.v2"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"

	"text/template"
)

//Master sources
var (
	schemesSourceURL   = "https://raw.githubusercontent.com/chriskempson/base16-schemes-source/master/list.yaml"
	templatesSourceURL = "https://raw.githubusercontent.com/chriskempson/base16-templates-source/master/list.yaml"
)

//Paths
var (
	schemesCachePath   = "cache/colorschemes/"
	templatesCachePath = "cache/templates/"
	schemesListFile    = schemesCachePath + "schemeslist.yml"
	templatesListFile  = templatesCachePath + "templateslist.yml"
)

//Flags
var (
	updateFlag         = kingpin.Flag("update-list", "Update the list of templates and colorschemes").Bool()
	clearListFlag      = kingpin.Flag("clear-list", "Delete local master list caches").Bool()
	clearSchemesFlag   = kingpin.Flag("clear-templates", "Delete local scheme caches").Bool()
	clearTemplatesFlag = kingpin.Flag("clear-schemes", "Delete local template caches").Bool()
)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()

	p1 := filepath.Join(".", schemesCachePath)
	os.MkdirAll(p1, os.ModePerm)

	p2 := filepath.Join(".", templatesCachePath)
	os.MkdirAll(p2, os.ModePerm)
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
	userInputThemeName := "metal"
	userInputTemplateName := "i3"

	schemeList := LoadBase16ColorschemeList()
	templateList := LoadBase16TemplateList()

	if *updateFlag {
		UpdateSchemes()
		UpdateTemplates()
	}

	schemeList = LoadBase16ColorschemeList()
	templateList = LoadBase16TemplateList()

	templ := templateList.Find(userInputTemplateName)
	fmt.Println("Selected template: ", templ.Name)

	scheme := schemeList.Find(userInputThemeName)
	fmt.Println("Selected scheme: ", scheme.Name)

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
