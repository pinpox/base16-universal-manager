package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Base16Colorscheme struct {
	Name   string `yaml:"scheme"`
	Author string `yaml:"author"`

	Color00 string `yaml:"base00"`
	Color01 string `yaml:"base01"`
	Color02 string `yaml:"base02"`
	Color03 string `yaml:"base03"`
	Color04 string `yaml:"base04"`
	Color05 string `yaml:"base05"`
	Color06 string `yaml:"base06"`
	Color07 string `yaml:"base07"`
	Color08 string `yaml:"base08"`
	Color09 string `yaml:"base09"`
	Color10 string `yaml:"base0A"`
	Color11 string `yaml:"base0B"`
	Color12 string `yaml:"base0C"`
	Color13 string `yaml:"base0D"`
	Color14 string `yaml:"base0E"`
	Color15 string `yaml:"base0F"`

	RepoURL string

	RawBaseURL string
	FileName   string
}

func (s Base16Colorscheme) MustacheContext(ext string) (map[string]interface{}, error) {
	bases := map[string]string{
		"00": s.Color00,
		"01": s.Color01,
		"02": s.Color02,
		"03": s.Color03,
		"04": s.Color04,
		"05": s.Color05,
		"06": s.Color06,
		"07": s.Color07,
		"08": s.Color08,
		"09": s.Color09,
		"0A": s.Color10,
		"0B": s.Color11,
		"0C": s.Color12,
		"0D": s.Color13,
		"0E": s.Color14,
		"0F": s.Color15,
	}
	slug := strings.Replace(strings.ToLower(s.FileName), " ", "-", -1)
	ret := map[string]interface{}{
		"scheme-name":             s.Name,
		"scheme-author":           s.Author,
		"scheme-slug":             slug,
		"scheme-slug-underscored": strings.Replace(slug, "-", "_", -1),
	}

	for base, color := range bases {
		baseKey := "base" + base

		if len(color) != 6 {
			return nil, fmt.Errorf("color %q for base %s has incorrect length, must be 6 hex digits", color, base)
		}

		// TODO
		rVal := color[:2]
		gVal := color[2:4]
		bVal := color[4:6]

		rValf, err := strconv.ParseUint(rVal, 16, 32)
		if err != nil {
			return nil, fmt.Errorf("bad hex red component for base %s: %q", base, rVal)
		}
		gValf, err := strconv.ParseUint(gVal, 16, 32)
		if err != nil {
			return nil, fmt.Errorf("bad hex green component for base %s: %q", base, gVal)
		}
		bValf, err := strconv.ParseUint(bVal, 16, 32)
		if err != nil {
			return nil, fmt.Errorf("bad hex blue component for base %s: %q", base, bVal)
		}

		ret[baseKey+"-hex"] = rVal + gVal + bVal

		ret[baseKey+"-hex-r"] = rVal
		ret[baseKey+"-hex-g"] = gVal
		ret[baseKey+"-hex-b"] = bVal

		ret[baseKey+"-rgb-r"] = rValf
		ret[baseKey+"-rgb-g"] = gValf
		ret[baseKey+"-rgb-b"] = bValf

		ret[baseKey+"-dec-r"] = rValf / 255
		ret[baseKey+"-dec-g"] = gValf / 255
		ret[baseKey+"-dec-b"] = bValf / 255

	}

	return ret, nil
}

func (l *Base16ColorschemeList) GetBase16Colorscheme(name string) (Base16Colorscheme, error) {
	if len(name) == 0 {
		return Base16Colorscheme{}, errors.New("colorscheme name was empty")
	}

	schemePath := path.Join(appConf.SchemesCachePath, name)

	parts := strings.Split(l.colorschemes[name], "/")
	yamlURL := strings.Join([]string{"https://raw.githubusercontent.com", parts[3], parts[4], parts[6], parts[7]}, "/")
	// Create local schemes file, if not present
	if _, err := os.Stat(schemePath); os.IsNotExist(err) {

		fmt.Println("downloading theme from: ", yamlURL)

		schemeData, err := DownloadFileToString(yamlURL)
		if err != nil {
			return Base16Colorscheme{}, fmt.Errorf("downloading %s: %w", yamlURL, err)
		}
		err = os.WriteFile(schemePath, []byte(schemeData), os.ModePerm)
		if err != nil {
			return Base16Colorscheme{}, fmt.Errorf("saving scheme to file: %w", err)
		}
	}

	colorscheme, err := os.ReadFile(schemePath)
	if err != nil {
		return Base16Colorscheme{}, fmt.Errorf("reading scheme from file %s: %w", schemePath, err)
	}

	scheme, err := NewBase16Colorscheme(string(colorscheme))
	if err != nil {
		return Base16Colorscheme{}, fmt.Errorf("creating scheme from file %s: %w", schemePath, err)
	}
	scheme.RawBaseURL = yamlURL
	scheme.FileName = parts[7][:len(parts[7])-5]
	return scheme, err
}

func (l *Base16ColorschemeList) GetBase16ColorschemeFile(name string) (Base16Colorscheme, error) {
	if len(name) == 0 {
		return Base16Colorscheme{}, fmt.Errorf("colorscheme path was empty")
	}

	colorscheme, err := os.ReadFile(name)
	if err != nil {
		return Base16Colorscheme{}, fmt.Errorf("reading scheme from file %s: %w", name, err)
	}

	scheme, err := NewBase16Colorscheme(string(colorscheme))
	if err != nil {
		return Base16Colorscheme{}, fmt.Errorf("creating scheme from file %s: %w", name, err)
	}
	_, fileName := path.Split(name)
	scheme.FileName = fileName[:len(fileName)-5]

	return scheme, err
}

func NewBase16Colorscheme(yamlData string) (Base16Colorscheme, error) {
	var scheme Base16Colorscheme

	err := yaml.Unmarshal([]byte(yamlData), &scheme)
	if err != nil {
		return Base16Colorscheme{}, fmt.Errorf("unmarshalling scheme from YAML: %w", err)
	}
	return scheme, nil
}

func LoadBase16ColorschemeList() (Base16ColorschemeList, error) {
	colorschemes, err := LoadStringMap(appConf.SchemesListFile)
	return Base16ColorschemeList{colorschemes}, err
}

func SaveBase16ColorschemeList(l Base16ColorschemeList) error {
	return SaveStringMap(l.colorschemes, appConf.SchemesListFile)
}

type Base16ColorschemeList struct {
	colorschemes map[string]string
}

func (l *Base16ColorschemeList) UpdateSchemes() error {
	// Get all repos from master source
	schemeRepos := make(map[string]string)

	schemesYAML, err := DownloadFileToString(appConf.SchemesMasterURL)
	if err != nil {
		return fmt.Errorf("downloading scheme from %s: %w", appConf.SchemesMasterURL, err)
	}

	err = yaml.Unmarshal([]byte(schemesYAML), &schemeRepos)
	if err != nil {
		return fmt.Errorf("unmarshalling YAML: %w", err)
	}

	fmt.Println("Found colorscheme repos: ", len(schemeRepos))

	for k, v := range schemeRepos {
		schemeRepos[k] = v
	}

	for _, v1 := range schemeRepos {
		fmt.Println("Getting schemes from: " + v1)

		for _, v2 := range findYAMLinRepo(v1) {
			l.colorschemes[v2.Name] = v2.HTMLURL
		}

	}

	fmt.Println("Found colorschemes: ", len(l.colorschemes))
	if err := SaveBase16ColorschemeList(Base16ColorschemeList{l.colorschemes}); err != nil {
		return fmt.Errorf("saving colorscheme list: %w", err)
	}

	return nil
}

func (c *Base16ColorschemeList) Find(input string) (Base16Colorscheme, error) {
	if strings.Contains(input, "/") {
		// a local scheme path, not a name
		colorschemeName := input
		scheme, err := c.GetBase16ColorschemeFile(colorschemeName)
		if err != nil {
			return Base16Colorscheme{}, fmt.Errorf("getting colorscheme file for %s: %w", colorschemeName, err)
		}
		return scheme, nil
	} else {
		// a name, look it up
		_, err := os.Stat(appConf.SchemesListFile)
		if os.IsNotExist(err) {
			fmt.Println("Colorschemes list not found, pulling new one...")
			if err := c.UpdateSchemes(); err != nil {
				return Base16Colorscheme{}, fmt.Errorf("updating schemes: %w", err)
			}
		} else if err != nil {
			return Base16Colorscheme{}, fmt.Errorf("checking for existing colorschems list: %w", err)
		}

		if len(c.colorschemes) == 0 {
			fmt.Println("No templates in list, pulling new one... ")
			if err := c.UpdateSchemes(); err != nil {
				return Base16Colorscheme{}, fmt.Errorf("updating schemes: %w", err)
			}
		}

		colorschemeName, err := FindMatchInMap(c.colorschemes, input)
		if err != nil {
			return Base16Colorscheme{}, fmt.Errorf("finding colorscheme in list: %w", err)
		}
		scheme, err := c.GetBase16Colorscheme(colorschemeName)
		if err != nil {
			return Base16Colorscheme{}, fmt.Errorf("getting colorscheme %s: %w", colorschemeName, err)
		}
		return scheme, nil
	}
}
