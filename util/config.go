package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	Modules []Module `yaml:"modules"`
}

type Module struct {
	Name string         `yaml:"name"`
	Data map[string]any `yaml:"data"`
}

func GetConfig() (*config, error) {
	path := os.Getenv("HOME") + "/.config/utilishonk/shonk.yml"

	// Check if path exists
	if _, err := os.Stat(path); err != nil {
		return nil, errors.New("config file not found")
	}

	// Check if path points to a yaml file
	if !(filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") {
		return nil, errors.New("config path does not lead to a yaml file")
	}

	// Read file
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf *config

	// Deserialize config gile
	if err := yaml.Unmarshal(f, &conf); err != nil {
		return nil, err
	}

	fmt.Println("Succesfully loaded condig:")
	fmt.Println(conf)

	return conf, nil
}
