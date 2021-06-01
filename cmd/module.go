package main

import (
	"godem/domain/models"
	"godem/infrastructure/database/user"
	"godem/infrastructure/middleware"
	"godem/infrastructure/service/_internal/http/api"
	jwthandler "godem/infrastructure/service/_internal/http/api/jwt"
	userhandler "godem/infrastructure/service/_internal/http/api/user"
	"godem/lib/util/database"
	"godem/usecase/jwt"
	useruc "godem/usecase/user"
	"log"
)

func loadModules(cfg *models.Config) (*api.ModuleHandler, *middleware.Auth) {
	mainDB, err := database.Connect(cfg)
	if err != nil {
		log.Fatalln("Cannot connect to DB", err.Error())
	}

	// Repo Section
	userRepo := user.NewMaster(mainDB)
	loginRepo := user.NewLogin(mainDB)

	// Usecase Section
	userUsecase := useruc.New(userRepo, loginRepo, cfg.JWT.PublicKey)
	jwtUsecase := jwt.New(cfg.JWT.PublicKey)

	// Handler Section
	userHandler := userhandler.NewHandler(userUsecase.Master(), userUsecase.Login())
	jwtHandler := jwthandler.NewHandler(jwtUsecase.JWT())

	return &api.ModuleHandler{
		User: userHandler,
		JWT:  jwtHandler,
	}, middleware.NewAuth(jwtUsecase.JWT())
}
