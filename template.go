package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

type Base16TemplateFile struct {
	Extension string `yaml:"extension"`
	Output    string `yaml:"output"`
}

type Base16Template struct {
	// The actual template
	Files map[string]Base16TemplateFile

	// Name (of the application)
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

func (l *Base16TemplateList) GetBase16Template(name string, remoteBranch string) (Base16Template, error) {
	// yamlURL := "https://raw.githubusercontent.com/" + parts[3] + "/" + parts[4] + "/master/templates/config.yaml"
	if len(name) == 0 {
		return Base16Template{}, errors.New("template name was empty")
	}

	var newTemplate Base16Template
	newTemplate.RepoURL = l.templates[name]
	rawBaseURL, err := GetRawBaseURL(l.templates[name], remoteBranch)
	if err != nil {
		return Base16Template{}, fmt.Errorf("generating template URL: %w", err)
	}
	newTemplate.RawBaseURL = rawBaseURL
	newTemplate.Name = name

	templatePath := path.Join(appConf.TemplatesCachePath, name+".yaml")

	// Create local template file, if not present
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		templateURL := newTemplate.RawBaseURL + "templates/config.yaml"
		templateData, err := DownloadFileToString(templateURL)
		if err != nil {
			return Base16Template{}, fmt.Errorf("downloading template from %s: %w", templateURL, err)
		}
		saveFile, err := os.Create(templatePath)
		// TODO delete old file?
		defer saveFile.Close()
		saveFile.Write([]byte(templateData))
		saveFile.Close()
	}

	template, err := os.ReadFile(templatePath)
	if err != nil {
		return Base16Template{}, fmt.Errorf("reading template file: %w", err)
	}

	// TODO cache actual templates
	var files map[string]Base16TemplateFile

	err = yaml.Unmarshal(template, &files)
	if err != nil {
		return Base16Template{}, fmt.Errorf("parsing template file %q: %w", templatePath, err)
	}

	newTemplate.Files = files
	return newTemplate, nil
}

type Base16TemplateList struct {
	templates map[string]string
}

func (l *Base16TemplateList) UpdateTemplates() error {
	// Get all repos from master source
	var templRepos map[string]string

	templatesYAML, err := DownloadFileToString(appConf.TemplatesMasterURL)
	if err != nil {
		return fmt.Errorf("downloading file from %s: %w", appConf.TemplatesMasterURL, err)
	}

	err = yaml.Unmarshal([]byte(templatesYAML), &templRepos)
	if err != nil {
		return fmt.Errorf("parsing file downloaded from %s: %w", appConf.TemplatesMasterURL, err)
	}

	fmt.Println("Found template repos: ", len(templRepos))
	for k, v := range templRepos {
		l.templates[k] = v
	}

	if err := SaveBase16TemplateList(Base16TemplateList{l.templates}); err != nil {
		return fmt.Errorf("saving template list: %w", err)
	}

	return nil
}

func LoadBase16TemplateList() (Base16TemplateList, error) {
	colorschemes, err := LoadStringMap(appConf.TemplatesListFile)
	return Base16TemplateList{colorschemes}, err
}

func SaveBase16TemplateList(l Base16TemplateList) error {
	return SaveStringMap(l.templates, appConf.TemplatesListFile)
}

func (c *Base16TemplateList) Find(input string) (Base16Template, error) {
	if _, err := os.Stat(appConf.TemplatesListFile); os.IsNotExist(err) {
		fmt.Println("Templates list not found, pulling new one...")
		if err := c.UpdateTemplates(); err != nil {
			return Base16Template{}, fmt.Errorf("updating templates: %w", err)
		}
	} else if err != nil {
		return Base16Template{}, fmt.Errorf("checking existance of template list: %w", err)
	}

	if len(c.templates) == 0 {
		fmt.Println("No templates in list, pulling new one... ")
		if err := c.UpdateTemplates(); err != nil {
			return Base16Template{}, fmt.Errorf("updating templates: %w", err)
		}
	}

	templateName, err := FindMatchInMap(c.templates, input)
	if err != nil {
		return Base16Template{}, fmt.Errorf("finding template in list: %w", err)
	}
	appConfig := appConf.Applications[input]
	remoteBranch := "master"
	if appConfig.DefaultRemoteBranch != "" {
		remoteBranch = appConfig.DefaultRemoteBranch
	}
	return c.GetBase16Template(templateName, remoteBranch)
}
