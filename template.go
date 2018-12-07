package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Base16Template struct {

	//The actual template
	Template string

	//Name (of the application)
	Name string
}

func (l *Base16TemplateList) GetBase16Template(name string) (Base16Template, error) {

	path := templatesCachePath + name

	// Create local template file, if not present
	if _, err := os.Stat(path); os.IsNotExist(err) {
		templateData, err := DownloadFileToStirng(l.templates[name])
		check(err)
		saveFile, err := os.Create(path)
		defer saveFile.Close()
		saveFile.Write([]byte(templateData))
		saveFile.Close()
	}

	template, err := ioutil.ReadFile(path)

	return NewBase16TemplateFromYAML(string(template)), err
}

func NewBase16TemplateFromYAML(yamlData string) Base16Template {
	//TODO
	return Base16Template{
		Template: yamlData,
		Name:     "Test",
	}
}

type Base16TemplateList struct {
	templates map[string]string
}

func UpdateTemplates() {

	//Get all repos from master source
	var templRepos map[string]string

	templatesYAML, err := DownloadFileToStirng(templatesSourceURL)
	check(err)

	err = yaml.Unmarshal([]byte(templatesYAML), &templRepos)
	check(err)

	templates := make(map[string]string)

	fmt.Println("Found template repos: ", len(templRepos))
	for k, v := range templRepos {
		templates[k] = v
	}

	SaveBase16TemplateList(Base16TemplateList{templates})

}

func LoadBase16TemplateList() Base16TemplateList {
	colorschemes := LoadStringMap(templatesListFile)
	return Base16TemplateList{colorschemes}
}

func SaveBase16TemplateList(l Base16TemplateList) {
	SaveStringMap(l.templates, templatesListFile)
}

func (c *Base16TemplateList) Find(input string) (Base16Template, error) {
	templateName := FindMatchInMap(c.templates, input)
	return c.GetBase16Template(templateName)
}
