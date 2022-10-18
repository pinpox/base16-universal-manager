package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

func GetRawBaseURL(repoURL string, mainBranch string) (string, error) {

	parts := strings.Split(repoURL, "/")
	rawBaseURL := ""
	repoName := parts[3] + "/" + parts[4]
	switch parts[2] {
	case "git.sr.ht":
		rawBaseURL = "https://git.sr.ht/" + repoName + "/blob/" + mainBranch + "/"
	case "github.com":
		rawBaseURL = "https://raw.githubusercontent.com/" + repoName + "/" + mainBranch + "/"
	case "gitlab.com":
		rawBaseURL = "https://gitlab.com/" + repoName + "/-/raw/" + mainBranch + "/"
	default:
		return "", fmt.Errorf("git host %q for %q not supported yet", parts[2], repoName)
	}
	return rawBaseURL, nil
}

func (l *Base16TemplateList) GetBase16Template(name string, remoteBranch string) Base16Template {

	// yamlURL := "https://raw.githubusercontent.com/" + parts[3] + "/" + parts[4] + "/master/templates/config.yaml"
	if len(name) == 0 {
		panic("Template name was empty")
	}

	var newTemplate Base16Template
	newTemplate.RepoURL = l.templates[name]
	rawBaseURL, err := GetRawBaseURL(l.templates[name], remoteBranch)
	check(err)
	newTemplate.RawBaseURL = rawBaseURL
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
	appConfig := appConf.Applications[input]
	remoteBranch := "master"
	if appConfig.DefaultRemoteBranch != "" {
		remoteBranch = appConfig.DefaultRemoteBranch
	}
	return c.GetBase16Template(templateName, remoteBranch)
}
