package main

import (
	"fmt"
	"github.com/hoisie/mustache"
	// "gopkg.in/yaml.v2"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"

	_ "text/template"
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
	configFile         = "config.yaml"
	schemesListFile    = schemesCachePath + "schemeslist.yaml"
	templatesListFile  = templatesCachePath + "templateslist.yaml"
	outputPath         = "out"
)

//Flags
var (
	updateFlag         = kingpin.Flag("update-list", "Update the list of templates and colorschemes").Bool()
	clearListFlag      = kingpin.Flag("clear-list", "Delete local master list caches").Bool()
	clearSchemesFlag   = kingpin.Flag("clear-templates", "Delete local scheme caches").Bool()
	clearTemplatesFlag = kingpin.Flag("clear-schemes", "Delete local template caches").Bool()
)

//Configuration
var appConf = NewConfig(configFile)

func main() {

	//Pase Flags
	kingpin.Version("1.0.0")
	kingpin.Parse()

	appConf.Show()

	//Create cache paths, if missing
	p1 := filepath.Join(".", schemesCachePath)
	os.MkdirAll(p1, os.ModePerm)
	p2 := filepath.Join(".", templatesCachePath)
	os.MkdirAll(p2, os.ModePerm)
	p3 := filepath.Join(".", outputPath)
	os.MkdirAll(p3, os.ModePerm)

	// TODO Get this two vars from flags
	userInputThemeName := "flat.yaml"
	userInputTemplateName := "termite"

	schemeList := LoadBase16ColorschemeList()
	templateList := LoadBase16TemplateList()

	if *updateFlag {
		schemeList.UpdateSchemes()
		templateList.UpdateTemplates()
	}

	//TODO delete caches, if user wants to
	schemeList = LoadBase16ColorschemeList()
	templateList = LoadBase16TemplateList()

	templ := templateList.Find(userInputTemplateName)
	fmt.Println("Selected template: ", templ.Name)

	scheme := schemeList.Find(userInputThemeName)
	fmt.Println("Selected scheme: ", scheme.Name)

	Base16Render(templ, scheme)

}

func Base16Render(templ Base16Template, scheme Base16Colorscheme) {

	fmt.Println("Rendering template: "+templ.Name+" with colorscheme: "+scheme.Name+" Files: ", len(templ.Files))

	for k, v := range templ.Files {
		templFileData, err := DownloadFileToStirng(templ.RawBaseURL + "templates/" + k + ".mustache")
		check(err)
		renderedFile := mustache.Render(templFileData, scheme.MustacheContext())

		p3 := filepath.Join(".", outputPath)
		os.MkdirAll(p3, os.ModePerm)

		saveBasePath := outputPath + "/" + templ.Name + "/"

		p4 := filepath.Join(".", saveBasePath)
		os.MkdirAll(p4, os.ModePerm)
		saveFilename := v.Output + v.Extension

		savePath := saveBasePath + saveFilename

		fmt.Println("writing file to: ", savePath)
		saveFile, err := os.Create(savePath)
		defer saveFile.Close()
		check(err)
		saveFile.Write([]byte(renderedFile))
		saveFile.Close()

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}

}
