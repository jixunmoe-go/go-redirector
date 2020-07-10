package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	DeployHost   string `yaml:"DeployHost"`
	ImportURL    string `yaml:"ImportURL"`
	BaseWebURL   string `yaml:"BaseWebURL"`
	DirectoryURL string `yaml:"DirectoryURL"`
	FileURL      string `yaml:"FileURL"`
	SourceURL    string `yaml:"SourceURL"`
	HomepageURL  string `yaml:"HomepageURL"`
	ListenAddr   string `yaml:"ListenAddr"`
}

var config Config

func readConfig() Config {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	return config
}
