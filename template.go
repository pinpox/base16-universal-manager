package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Base16Template struct {

	//The actual template
	Template string

	//Name (of the application)
	Name string
}

func (*Base16Template) Render() (string, error) {

	return "", nil
}

func GetBase16Template(name string) (Base16Template, error) {
	//TODO get from the internets instead (if possible)
	template, err := ioutil.ReadFile("./templates/" + name)

	return Base16Template{
		Template: string(template),
		Name:     "Test",
	}, err
}

type Base16TemplateList struct {
	URLs []string
}

func (c *Base16TemplateList) add(url string) {
	c.URLs = append(c.URLs, url)
}

func (c *Base16TemplateList) UpdateFromRemote(url ...string) {

	//Set the source for templates
	var masterRepo string
	if len(url) == 1 {
		masterRepo = url[0]
	} else {
		masterRepo = "https://raw.githubusercontent.com/chriskempson/base16-templates-source/master/list.yaml"
	}

	//Get all repos from master source
	var templRepos map[string]string

	templatesYAML, err := DownloadFileToStirng(masterRepo)
	check(err)

	err = yaml.Unmarshal([]byte(templatesYAML), &templRepos)
	check(err)

	fmt.Println("Found template repos: ", len(templRepos))
	for k, v := range templRepos {
		fmt.Printf("%s: %s\n", k, v)
		c.add(v)
	}
}

func (c *Base16TemplateList) Find(input string) (Base16Template, error) {
	return Base16Template{}, nil
}