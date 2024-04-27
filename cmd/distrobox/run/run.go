package run

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type runConfig struct {
	command       string `yaml:"command"`
	environmentId string `yaml:"environment-id"`
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a saved command in the right environment",
	Long:  "Run a saved command in the right distrobox or in the host, regardless of the current environment",
	Run: func(cmd *cobra.Command, arg []string) {
		conf, err := getConfig(os.Getenv("HOME") + "/.config/utilishonk/distrobox/run/" + arg[0] + ".yml")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Running this command:")
		fmt.Println(conf.command)

		fmt.Println("In this environment:")
		fmt.Println(conf.environmentId)

		fmt.Println("From this environment:")
		id, err := getEnvironmentId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	},
}

func getEnvironmentId() (string, error) {
	// Get container ID
	// If it's an empty string, we're on the host system
	if containerId := os.Getenv("CONTAINER_ID"); containerId != "" {
		// Check for ambiguous naming
		if containerId == "host" {
			return "", errors.New("`host` is an ambiguous distrobox name. Shonk is confused!")
		}

		return containerId, nil
	} else {
		return "host", nil
	}
}

func getConfig(path string) (runConfig, error) {
	// Check if path exists
	if _, err := os.Stat(path); err != nil {
		return runConfig{}, errors.New("distrobox run: config doesn't exist")
	}

	// Check if path points to a yaml file
	if !(filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") {
		return runConfig{}, errors.New("distrobox run: config is not a yaml file")
	}

	// Read config file
	f, err := os.ReadFile(path)
	if err != nil {
		return runConfig{}, err
	}

	var conf runConfig

	// Deserialize config file
	if err := yaml.Unmarshal(f, &conf); err != nil {
		return runConfig{}, err
	}

	fmt.Println("Succesfully loaded config:")
	fmt.Println(conf)

	return conf, nil
}
