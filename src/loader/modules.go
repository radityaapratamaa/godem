package loader

import (
	"godem/src/controller"
	userctrl "godem/src/controller/user"
	"godem/src/repository/user"
	useruc "godem/src/usecase/user"
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
	userControllers := userctrl.New(userUsecase)
	meController := controller.NewMe()

	// register controller to app
	// register multi/group controller
	a.App.AddControllers(userControllers)

	// register single controller
	a.App.AddController(meController)
}
