package main

import (
	"io/ioutil"
)

func getTemplate(url string) (string, error) {
	b16templateName := "i3"
	template, err := ioutil.ReadFile("./templates/" + b16templateName)
	return string(template), err
}

func GetB16Template(name string) (string, error) {
	return nil
}
