package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	JiraServerDomain string `yaml:"jiraServerDomain"`
	JiraAuth         string `yaml:"JiraAuth"`
}

func GetConfig(configPath string) (*Config, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal yaml file: %s", err)
	}

	return &config, nil
}

func CreateConfig(path string, filename string, config Config) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create directory: %s", err)
	}

	file, err := os.Create(path + filename)
	if err != nil {
		return fmt.Errorf("could not create file: %s", err)
	}
	defer file.Close()

	data, err := yaml.Marshal(config)

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("could not write to file: %s", err)
	}

	return nil
}
