package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Action struct {
	Key     string
	Command string
}

type Configuration map[string]map[string]Action

func LoadConfiguration(file string) (*Configuration, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var configuration Configuration
	err = yaml.Unmarshal(source, &configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}
