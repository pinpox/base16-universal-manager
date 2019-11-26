package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hoisie/mustache"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Configuration file
var configFile string

//Flags
var (
	updateFlag         = kingpin.Flag("update-list", "Update the list of templates and colorschemes").Bool()
	clearListFlag      = kingpin.Flag("clear-list", "Delete local master list caches").Bool()
	clearTemplatesFlag = kingpin.Flag("clear-templates", "Delete local scheme caches").Bool()
	clearSchemesFlag   = kingpin.Flag("clear-schemes", "Delete local template caches").Bool()
	configFileFlag     = kingpin.Flag("config", "Specify configuration file to use").Default("config.yaml").String()
)

//Configuration
var appConf SetterConfig

func main() {

	//Pase Flags
	kingpin.Version("1.0.0")
	kingpin.Parse()

	appConf = NewConfig(*configFileFlag)

	// appConf.Show()
	//TODO delete caches, if user wants to

	//Create cache paths, if missing
	p1 := filepath.Join(".", appConf.SchemesCachePath)
	os.MkdirAll(p1, os.ModePerm)
	p2 := filepath.Join(".", appConf.TemplatesCachePath)
	os.MkdirAll(p2, os.ModePerm)

	schemeList := LoadBase16ColorschemeList()
	templateList := LoadBase16TemplateList()

	if *updateFlag {
		schemeList.UpdateSchemes()
		templateList.UpdateTemplates()
	}

	scheme := schemeList.Find(appConf.Colorscheme)
	fmt.Println("[CONFIG]: Selected scheme: ", scheme.Name)

	templateEnabled := false
	for app, appConfig := range appConf.Applications {
		if appConfig.Enabled {
			Base16Render(templateList.Find(app), scheme)
			templateEnabled = true
		}
	}

	if !templateEnabled {
		fmt.Println("No templates enabled")
	}

}

func Base16Render(templ Base16Template, scheme Base16Colorscheme) {

	fmt.Println("[RENDER]: Rendering template \"" + templ.Name + "\"")

	for k, v := range templ.Files {
		templFileData, err := DownloadFileToStirng(templ.RawBaseURL + "templates/" + k + ".mustache")
		check(err)
		renderedFile := mustache.Render(templFileData, scheme.MustacheContext())

		saveBasePath := appConf.Applications[templ.Name].Files[k] + "/"
		p4 := filepath.Join(".", saveBasePath)
		os.MkdirAll(p4, os.ModePerm)
		savePath := saveBasePath + k + v.Extension

		//If DryRun is enabled, just print the output location for debugging
		if appConf.DryRun {
			fmt.Println("    - (dryrun) file would be written to: ", savePath)
		} else {
			switch appConf.Applications[templ.Name].Mode {
			case "rewrite":
				fmt.Println("     - writing: ", savePath)
				saveFile, err := os.Create(savePath)
				defer saveFile.Close()
				check(err)
				saveFile.Write([]byte(renderedFile))
				saveFile.Close()
			case "append":
				fmt.Println("     - appending to: ", savePath)
			case "replace":
				fmt.Println("     - replacing in: ", savePath)
			}
		}
	}

	if appConf.DryRun {
		fmt.Println("Not running hook, DryRun enabled: ", appConf.Applications[templ.Name].Hook)
	} else {
		exe_cmd(appConf.Applications[templ.Name].Hook)
	}
}

//TODO proper error handling
func check(e error) {
	if e != nil {
		panic(e)
	}

}
