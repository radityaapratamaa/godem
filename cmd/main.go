package main

import (
	"log"

	"github.com/pkg/errors"

	"godem/infrastructure/config"
	"godem/infrastructure/service/_internal/http/api"
)

func main() {
	// read config file
	cfg, err := config.New()
	if err != nil {
		log.Fatalln("cannot read config: ", errors.Cause(err).Error())
	}
	log.Printf("%#v", cfg)
	modules, authMiddleware := loadModules(cfg)

	handler := api.NewHandler(modules, authMiddleware)
	log.Fatalln(handler.RegisterAndStartServer())
}
