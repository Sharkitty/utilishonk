package run

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"utilishonk/util"
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
		conf, err := util.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		var distroboxConf util.Module
		for _, module := range conf.Modules {
			if module.Name == "distrobox" {
				distroboxConf = module
				break
			}
		}

		fmt.Println("Running this command:")
		fmt.Println(distroboxConf.Data["command"])

		fmt.Println("In this environment:")
		fmt.Println(distroboxConf.Data["environment-id"])

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
