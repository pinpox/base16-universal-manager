package main

import (
	"encoding/json"

	"fmt"
	"github.com/agnivade/levenshtein"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//DownloadFileToStirng downloads a file from a given URL and returns it's
//contents as a string if successful
func DownloadFileToStirng(url string) (string, error) {

	fmt.Println("Downloading ", url)

	var client http.Client
	resp, err := client.Get(url + "?access_token=" + appConf.GithubToken)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	}

	return "", err
}

type GitHubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type GitHubFilesCollection struct {
	Collection []GitHubFile
}

func findYAMLinRepo(repoURL string) []GitHubFile {

	parts := strings.Split(repoURL, "/")
	ApiUrl := ("https://api.github.com/repos/" + parts[3] + "/" + parts[4] + "/contents/")
	// fmt.Println("generated api URL: ", ApiUrl)

	// Get all files from repo
	// repoFiles, err := DownloadFileToStirng("https://api.github.com/repos/atelierbram/base16-atelier-schemes/contents/")
	repoFiles, err := DownloadFileToStirng(ApiUrl)
	check(err)
	keys := make([]GitHubFile, 0)
	json.Unmarshal([]byte(repoFiles), &keys)

	// Create a list of .yaml files
	var colorSchemes []GitHubFile
	for _, v := range keys {
		re := regexp.MustCompile(".*yaml")
		if re.MatchString(v.Name) {
			colorSchemes = append(colorSchemes, v)
		}
	}

	// fmt.Println("Found ", len(colorSchemes), "in repo ", repoFiles)
	return colorSchemes
}

func LoadStringMap(path string) map[string]string {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	check(err)
	yamlFile, err := ioutil.ReadAll(f)
	check(err)
	data := make(map[string]string)
	err = yaml.Unmarshal(yamlFile, data)
	check(err)
	return data
}

func SaveStringMap(data map[string]string, path string) {

	yamlData, err := yaml.Marshal(data)
	check(err)
	saveFile, err := os.Create(path)
	defer saveFile.Close()
	saveFile.Write(yamlData)
	saveFile.Close()
	fmt.Println("wrote to: ", saveFile.Name())
}

func FindMatchInMap(choices map[string]string, input string) string {

	var match string
	distance := 1000

	for k := range choices {
		tempDistance := levenshtein.ComputeDistance(input, k)
		if tempDistance < distance {
			match = k
			distance = tempDistance
		}
	}

	return match
}
