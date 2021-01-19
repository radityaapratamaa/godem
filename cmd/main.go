package main

import (
	"article-test/infrastructure/config"
	"log"

	"github.com/pkg/errors"
)

func main() {
	// read config file
	cfg, err := config.New()
	if err != nil {
		log.Fatalln("cannot read config: ", errors.Cause(err).Error())
	}
	log.Printf("%#v", cfg)
}
