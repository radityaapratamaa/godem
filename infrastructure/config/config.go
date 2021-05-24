package config

import (
	"bcg-test/domain/models"
	"bcg-test/lib/helper"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func New() (*models.Config, error) {
	var (
		env      = helper.GetEnv()
		filename = fmt.Sprintf("%s.yaml", env)
	)

	pwd, _ := os.Getwd()
	filepath := filepath.Join(pwd, "files/", filename)
	log.Println("reading config file: ", filepath)

	configFile, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.config.New.OpenFile")
	}
	defer configFile.Close()

	var configuration models.Config
	if err := yaml.NewDecoder(configFile).Decode(&configuration); err != nil {
		return nil, errors.Wrap(err, "infrastructure.config.New.DecodeYaml")
	}
	return &configuration, nil
}
