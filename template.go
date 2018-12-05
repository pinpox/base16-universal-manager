package main

import (
	"io/ioutil"
)

func GetBase16Template(name string) (string, error) {
	//TODO get from the internets instead (if possible)
	template, err := ioutil.ReadFile("./templates/" + name)
	return string(template), err
}
