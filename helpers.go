package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

//DownloadFileToStirng downloads a file from a given URL and returns it's
//contents as a string if successful
func DownloadFileToStirng(url string) (string, error) {

	var client http.Client
	resp, err := client.Get(url)
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

	// Get all files from repo
	//TODO create and use proper URL form variable
	repoFiles, err := DownloadFileToStirng("https://api.github.com/repos/atelierbram/base16-atelier-schemes/contents/")
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
	return colorSchemes
}
