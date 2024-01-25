package main

import (
	"embed"
	"log"

	"github.com/joho/godotenv"

	"godem/infrastructure/loader"
)

//go:embed files/*
var templateFolder embed.FS

func main() {
	// read config file
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln("cannot read config: ", err.Error())
	}

	app := loader.NewApp(loader.WithFolderTemplate(templateFolder))
	if err != nil {
		log.Fatalln(err)
	}

	if err = app.Start(); err != nil {
		log.Fatalln(err.Error())
	}
}
