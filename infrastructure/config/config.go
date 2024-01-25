package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"godem/domain/models"
	"godem/lib/helper"
	"gopkg.in/yaml.v2"
)

var Cfg models.Config

func New() error {
	var (
		env      = helper.GetEnv()
		filename = fmt.Sprintf("%s.yaml", env)
	)

	pwd, _ := os.Getwd()
	configFilePath := filepath.Join(pwd, "files/config", filename)
	log.Println("reading config file: ", configFilePath)

	filePath := filepath.Clean(configFilePath)
	configFile, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "infrastructure.config.New.OpenFile")
	}
	defer configFile.Close()

	if err := yaml.NewDecoder(configFile).Decode(&Cfg); err != nil {
		return errors.Wrap(err, "infrastructure.config.New.DecodeYaml")
	}

	Cfg.Server.Environment = env
	return nil
}
