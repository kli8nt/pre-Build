package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Mohamed-Rafraf/building/utils"
	"github.com/caarlos0/env/v8"
)

var err error
var data utils.Data

func DockerfileExists(folderPath string) bool {
	dockerfilePath := filepath.Join(folderPath, "Dockerfile")
	_, err := os.Stat(dockerfilePath)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// Handle other error cases if needed
		return false
	}
}

func main() {
	if err := env.Parse(&data); err != nil {
		log.Printf("%+v\n", err)
	}

	err = utils.CloneRepo(data.RepositoryURL, data.GithubToken, data.ApplicationName)
	if err != nil {
		log.Println(err)
	}
	if !DockerfileExists(data.ApplicationName) {
		err = utils.Build(&data)
		if err != nil {
			log.Println(err)
		}
	}
}
