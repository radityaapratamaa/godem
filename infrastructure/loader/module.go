package loader

import (
	"log"

	"github.com/kodekoding/phastos/go/database"

	"godem/domain/models"
	"godem/infrastructure/database/user"
	"godem/infrastructure/middleware/auth"
	"godem/infrastructure/service/_internal/http/api"
	userhandler "godem/infrastructure/service/_internal/http/api/v1/user"
	"godem/usecase/jwt"
	useruc "godem/usecase/user"
)

func InitModules(cfg *models.Config) *api.Modules {
	// load main DB
	db, err := database.Connect(&cfg.Databases)
	if err != nil {
		log.Fatalln("error when connect to DB: ", err.Error())
	}

	// load external source

	// load repo section
	userRepo := user.New(db)

	// load usecase section
	userUsecase := useruc.New(userRepo, cfg.JWT.PublicKey)
	jwtUc := jwt.New(cfg.JWT.SigningKey)

	// load handler section
	userHandler := userhandler.New(userUsecase)
	// sent Modules to router
	return &api.Modules{
		User: userHandler,
		Auth: auth.New(jwtUc),
	}
}
