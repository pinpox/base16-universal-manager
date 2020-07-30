package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"path"
)

type Base16TemplateFile struct {
	Extension string `yaml:"extension"`
	Output    string `yaml:"output"`
}

type Base16Template struct {

	//The actual template
	Files map[string]Base16TemplateFile

	//Name (of the application)
	Name string

	RepoURL string

	RawBaseURL string
}

func (l *Base16TemplateList) GetBase16Template(name string) Base16Template {

	// yamlURL := "https://raw.githubusercontent.com/" + parts[3] + "/" + parts[4] + "/master/templates/config.yaml"
	if len(name) == 0 {
		panic("Template name was empty")
	}

	var newTemplate Base16Template
	newTemplate.RepoURL = l.templates[name]
	parts := strings.Split(l.templates[name], "/")
	newTemplate.RawBaseURL = "https://raw.githubusercontent.com/" + parts[3] + "/" + parts[4] + "/master/"
	newTemplate.Name = name

	templatePath := path.Join(appConf.TemplatesCachePath, name+".yaml")

	// Create local template file, if not present
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		templateData, err := DownloadFileToString(newTemplate.RawBaseURL + "templates/config.yaml")
		check(err)
		saveFile, err := os.Create(templatePath)
		//TODO delete old file?
		defer saveFile.Close()
		saveFile.Write([]byte(templateData))
		saveFile.Close()
	}

	template, err := ioutil.ReadFile(templatePath)
	check(err)

	//TODO cache actual templates
	var files map[string]Base16TemplateFile

	err = yaml.Unmarshal(template, &files)
	check(err)

	newTemplate.Files = files
	return newTemplate
}

type Base16TemplateList struct {
	templates map[string]string
}

func (l *Base16TemplateList) UpdateTemplates() {

	//Get all repos from master source
	var templRepos map[string]string

	templatesYAML, err := DownloadFileToString(appConf.TemplatesMasterURL)
	check(err)

	err = yaml.Unmarshal([]byte(templatesYAML), &templRepos)
	check(err)

	fmt.Println("Found template repos: ", len(templRepos))
	for k, v := range templRepos {
		l.templates[k] = v
	}

	SaveBase16TemplateList(Base16TemplateList{l.templates})

}

func LoadBase16TemplateList() Base16TemplateList {
	colorschemes := LoadStringMap(appConf.TemplatesListFile)
	return Base16TemplateList{colorschemes}
}

func SaveBase16TemplateList(l Base16TemplateList) {
	SaveStringMap(l.templates, appConf.TemplatesListFile)
}

func (c *Base16TemplateList) Find(input string) Base16Template {

	if _, err := os.Stat(appConf.TemplatesListFile); os.IsNotExist(err) {
		check(err)
		fmt.Println("Templates list not found, pulling new one...")
		c.UpdateTemplates()
	}

	if len(c.templates) == 0 {
		fmt.Println("No templates in list, pulling new one... ")
		c.UpdateTemplates()
	}

	templateName := FindMatchInMap(c.templates, input)
	return c.GetBase16Template(templateName)
}
