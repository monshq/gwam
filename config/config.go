package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Addon struct {
	Name string
	Path string
	Url  string
}

type Config struct {
	Addons []Addon
}

func ReadConfig() Config {
	yfile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	data := Config{}
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
