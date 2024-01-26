package loader

import (
	userctrl "godem/controller/user"
	"godem/infrastructure/database/user"

	useruc "godem/usecase/user"
)

func (a *app) loadModules() {

	// get resources from api-app phastos
	db := a.DB()
	trx := a.Trx()
	redis := a.Cache()

	// init repo
	userRepo := user.New(db)

	// init usecase
	userUsecase := useruc.New(
		userRepo,
		smtp,
		a.templateFolder,
		trx,
		redis,
	)

	// init controller
	userController := userctrl.New(userUsecase)

	// register controller to app
	a.App.AddController(userController.Master())
	a.App.AddController(userController.Auth())
}
