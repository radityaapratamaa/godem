package main

import (
	"log"

	"github.com/pkg/errors"

	"bcg-test/infrastructure/config"
	"bcg-test/infrastructure/service/_internal/http/api"
)

func main() {
	// read config file
	cfg, err := config.New()
	if err != nil {
		log.Fatalln("cannot read config: ", errors.Cause(err).Error())
	}
	log.Printf("%#v", cfg)
	modules := loadModules(cfg)

	handler := api.NewHandler(modules)
	log.Fatalln(handler.RegisterAndStartServer())
}
