package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Bind9 struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Queries string `yaml:"queries"`
	} `yaml:"bind9"`
	Categories struct {
		ToCheck []string `yaml:"toCheck"`
		ToProtect []string `yaml:"toProtect"`
	} `yaml:"categories"`
}

func loadConfig() Config {
	var config Config
	f, err := os.Open("/etc/go-bind9masq/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
