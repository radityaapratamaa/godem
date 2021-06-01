package main

import (
	"godem/domain/models"
	"godem/infrastructure/database/user"
	"godem/infrastructure/service/_internal/http/api"
	userhandler "godem/infrastructure/service/_internal/http/api/user"
	"godem/lib/util/database"
	useruc "godem/usecase/user"
	"log"
)

func loadModules(cfg *models.Config) *api.ModuleHandler {
	mainDB, err := database.Connect(cfg)
	if err != nil {
		log.Fatalln("Cannot connect to DB", err.Error())
	}

	userRepo := user.NewMaster(mainDB)
	userUsecase := useruc.New(userRepo)
	userHandler := userhandler.NewHandler(userUsecase.Master())

	return &api.ModuleHandler{
		User: userHandler,
	}
}
