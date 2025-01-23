package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/agnivade/levenshtein"
	"gopkg.in/yaml.v2"
)

// DownloadFileToString downloads a file from a given URL and returns it's
// contents as a string if successful
func DownloadFileToString(url string) (string, error) {
	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if appConf.GithubToken != "" {
		req.Header.Add("Authorization", "token "+appConf.GithubToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	} else if err == nil {
		err = fmt.Errorf("HTTP code %v", resp.StatusCode)
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
	// repoFiles, err := DownloadFileToString("https://api.github.com/repos/atelierbram/base16-atelier-schemes/contents/")
	repoFiles, err := DownloadFileToString(ApiUrl)
	if err != nil {
		fmt.Println("Failed to get schemes from: " + repoURL + " with error '" + err.Error() + "'")
		return nil
	}
	keys := make([]GitHubFile, 0)
	json.Unmarshal([]byte(repoFiles), &keys)

	// Create a list of .yaml files
	var colorSchemes []GitHubFile
	for _, v := range keys {
		re := regexp.MustCompile(".*ya?ml")
		if re.MatchString(v.Name) {
			colorSchemes = append(colorSchemes, v)
		}
	}

	// fmt.Println("Found ", len(colorSchemes), "in repo ", repoFiles)
	return colorSchemes
}

func LoadStringMap(path string) (map[string]string, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading from file %s: %w", path, err)
	}
	data := make(map[string]string)
	if err := yaml.Unmarshal(yamlFile, data); err != nil {
		return nil, fmt.Errorf("parsing YAML from file %s: %w", path, err)
	}
	return data, nil
}

func SaveStringMap(data map[string]string, path string) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshaling to YAML: %w", err)
	}
	if err := os.WriteFile(path, yamlData, os.ModePerm); err != nil {
		return fmt.Errorf("writing YAML: %w", err)
	}
	fmt.Println("wrote to: ", path)
	return nil
}

func FindMatchInMap(choices map[string]string, input string) (string, error) {
	if len(choices) == 0 {
		return "", errors.New("cannot select from empty choices")
	}
	var matches []string
	var distance int

	for k := range choices {
		tempDistance := levenshtein.ComputeDistance(input, k)
		if tempDistance < distance || len(matches) == 0 {
			distance = tempDistance
			matches = nil
			matches = append(matches, k)
		} else if tempDistance == distance {
			matches = append(matches, k)
		}
	}

	if len(matches) != 1 {
		return "", fmt.Errorf(
			"found %d matches to %q with same similarity score, consider specifying a closer match to a candidate - candidates are: %v",
			len(matches),
			input,
			strings.Join(matches, ", "),
		)
	}

	return matches[0], nil
}

func exe_cmd(cmd string) {
	if len(cmd) == 0 {
		return
	}
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	out, err := exec.Command(head, parts...).Output()

	fmt.Println("[HOOK]: Running: ", cmd)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", out)
}

func WriteFile(path string, contents []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not open file %q: %w", path, err)
	}
	defer file.Close()

	if _, err = file.Write(contents); err != nil {
		return fmt.Errorf("could not write in file %q: %w", path, err)
	}

	if err = file.Sync(); err != nil {
		return fmt.Errorf("could not flush file contents %q: %w", path, err)
	}

	return nil
}

func ReplaceMultiline(filepath, replaceContents, startMarker, endMarker string) error {
	if startMarker == "" {
		return fmt.Errorf("start marker regular expression cannot be empty if file mode is replace")
	}
	if endMarker == "" {
		return fmt.Errorf("end marker regular expression cannot be empty if file mode is replace")
	}
	startMarkerRegex, err := regexp.Compile(startMarker)
	if err != nil {
		return fmt.Errorf("invalid start marker regular expression: %w", err)
	}
	endMarkerRegex, err := regexp.Compile(endMarker)
	if err != nil {
		return fmt.Errorf("invalid end marker regular expression: %w", err)
	}
	newContents, err := getReplacedContents(filepath, replaceContents, startMarkerRegex, endMarkerRegex)
	if err != nil {
		return fmt.Errorf("could not replace in file %q: %w", filepath, err)
	}
	return WriteFile(filepath, newContents)
}

func getReplacedContents(filepath, replaceContents string, startMarkerRegex, endMarkerRegex *regexp.Regexp) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open %q: %w", filepath, err)
	}
	var buffer bytes.Buffer
	var startFound, endFound bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if !startFound && startMarkerRegex.Match(line) {
			startFound = true
			buffer.Write(line)
			buffer.WriteString("\n")
			buffer.Write([]byte(replaceContents))
		} else if !startFound {
			buffer.Write(line)
			buffer.WriteString("\n")
		} else if !endFound && endMarkerRegex.Match(line) {
			endFound = true
			buffer.Write(line)
			buffer.WriteString("\n")
		} else if startFound && endFound {
			buffer.Write(line)
			buffer.WriteString("\n")
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read file %q: %w", filepath, err)
	}
	if !startFound {
		return nil, fmt.Errorf("could not find a line matching start_marker regex in %q", filepath)
	}
	if !endFound {
		return nil, fmt.Errorf("could not find a line matching end_marker regex in %q", filepath)
	}
	return buffer.Bytes(), nil
}

func deepCompareFiles(file1, file2 string) bool {
	sf, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}

	df, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}

	sscan := bufio.NewScanner(sf)
	dscan := bufio.NewScanner(df)

	for sscan.Scan() {
		dscan.Scan()
		if !bytes.Equal(sscan.Bytes(), dscan.Bytes()) {
			return false
		}
	}

	return true
}
