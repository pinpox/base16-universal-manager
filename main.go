package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/OpenPeeDeeP/xdg"
	"github.com/hoisie/mustache"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Configuration file
var configFile string

var xdgDirs = xdg.New("base16-universal-manager", "")

//Flags
var (
	updateFlag         = kingpin.Flag("update-list", "Update the list of templates and colorschemes").Bool()
	clearListFlag      = kingpin.Flag("clear-list", "Delete local template and colorscheme list caches").Bool()
	clearTemplatesFlag = kingpin.Flag("clear-templates", "Delete local template caches").Bool()
	clearSchemesFlag   = kingpin.Flag("clear-schemes", "Delete local schemes caches").Bool()
	configFileFlag     = kingpin.Flag("config", "Specify configuration file to use").Default(xdgDirs.QueryConfig("config.yaml")).String()
	printConfigFlag    = kingpin.Flag("print-config", "Print current configuration").Bool()
	schemeFlag         = kingpin.Flag("scheme", "Specify scheme to use (Overrides config)").String()
)

//Configuration
var appConf SetterConfig

func main() {
	//Parse Flags
	kingpin.Version("1.0.0")
	kingpin.Parse()

	appConf = NewConfig(*configFileFlag)

	if *printConfigFlag {
		appConf.Show()
	}

	if *clearListFlag {
		err := os.Remove(appConf.SchemesListFile)
		if err == nil {
			fmt.Printf("Deleted cached colorscheme list %s\n", appConf.SchemesListFile)
		} else {
			fmt.Fprintf(os.Stderr, "Error deleting cached colorscheme list: %v\n", err)
		}
		err = os.Remove(appConf.TemplatesListFile)
		if err == nil {
			fmt.Printf("Deleted cached template list %s\n", appConf.TemplatesListFile)
		} else {
			fmt.Fprintf(os.Stderr, "Error deleting cached template list: %v\n", err)
		}
	}

	if *clearSchemesFlag {
		err := os.RemoveAll(appConf.SchemesCachePath)
		if err == nil {
			fmt.Printf("Deleted cached colorscheme list %s\n", appConf.SchemesCachePath)
		} else {
			fmt.Fprintf(os.Stderr, "Error deleting cached colorschemes: %v\n", err)
		}
	}

	if *clearTemplatesFlag {
		err := os.RemoveAll(appConf.TemplatesCachePath)
		if err == nil {
			fmt.Printf("Deleted cached templates %s\n", appConf.TemplatesCachePath)
		} else {
			fmt.Fprintf(os.Stderr, "Error deleting cached templates: %v\n", err)
		}
	}

	// Create cache paths, if missing
	os.MkdirAll(appConf.SchemesCachePath, os.ModePerm)
	os.MkdirAll(appConf.TemplatesCachePath, os.ModePerm)

	schemeList := LoadBase16ColorschemeList()
	templateList := LoadBase16TemplateList()

	if *updateFlag {
		schemeList.UpdateSchemes()
		templateList.UpdateTemplates()
	}

	var scheme Base16Colorscheme
	if *schemeFlag == "" {
	    // Scheme from config
	    scheme = schemeList.Find(appConf.Colorscheme)
	} else {
	    // Scheme from flag
	    scheme = schemeList.Find(*schemeFlag)
	}
	fmt.Println("[CONFIG]: Selected scheme: ", scheme.Name)

	templateEnabled := false
	for app, appConfig := range appConf.Applications {
        if appConfig.Template == "" {
            appConfig.Template = app
        }
		if appConfig.Enabled {
			err := Base16Render(templateList.Find(appConfig.Template), scheme, app)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when rendering file: %v\n", err)
			}
			templateEnabled = true
		}
	}

	if !templateEnabled {
		fmt.Println("No templates enabled")
	}

}

// Base16Render takes an application-specific template and renders a config file
// implementing the provided colorscheme.
func Base16Render(templ Base16Template, scheme Base16Colorscheme, app string) error {
	fmt.Println("[RENDER]: Rendering template \"" + templ.Name + "\"")

	for k, v := range templ.Files {
		templFileData, err := DownloadFileToString(templ.RawBaseURL + "templates/" + k + ".mustache")
		if err != nil {
			return fmt.Errorf("could not download template file: %w", err)
		}
		renderedFile := mustache.Render(templFileData, scheme.MustacheContext(v.Extension))

		savePath, err := getSavePath(appConf.Applications[app].Files[k].Path, k+v.Extension)
		if err != nil {
			return fmt.Errorf("could not get location for save path: %w", err)
		}
		if savePath == "" {
			continue
		}

		//If DryRun is enabled, just print the output location for debugging
		if appConf.DryRun {
			fmt.Println("    - (dryrun) file would be written to: ", savePath)
		} else {
			switch appConf.Applications[app].Files[k].Mode {
			case "rewrite":
				fmt.Println("     - writing: ", savePath)
				if err = WriteFile(savePath, []byte(renderedFile)); err != nil {
					return err
				}
			case "replace":
				fmt.Println("     - replacing in: ", savePath)
				startMarker := appConf.Applications[app].Files[k].StartMarker
				endMarker := appConf.Applications[app].Files[k].EndMarker
				if err = ReplaceMultiline(savePath, renderedFile, startMarker, endMarker); err != nil {
					return err
				}
			}
		}
	}

	if appConf.DryRun {
		fmt.Println("Not running hook, DryRun enabled: ", appConf.Applications[app].Hook)
	} else {
		exe_cmd(appConf.Applications[app].Hook)
	}

	return nil
}

func getSavePath(path, defaultFilename string) (string, error) {
	if path == "" {
		return "", nil
	}

	var homeDir string
	if path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("could not get user home directory: %w", err)
		}
		homeDir = usr.HomeDir
	}

	var savePath string
	if path == "~" {
		savePath = homeDir
	} else if strings.HasPrefix(path, "~/") {
		savePath = filepath.Join(homeDir, path[2:])
	} else if path[0] != '/' {
		savePath = filepath.Join(".", path)
	} else {
		savePath = path
	}

	if strings.HasSuffix(path, "/") {
		err := os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("could not create folder: %w", err)
		}
		savePath = filepath.Join(savePath, defaultFilename)
	}

	return savePath, nil
}

//TODO proper error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}
