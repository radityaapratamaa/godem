package user

import (
	mockuserdbrepo "godem/infrastructure/database/mocks/user"
)

var (
	masterRepo *mockuserdbrepo.Masters
	loginRepo  *mockuserdbrepo.Logins
	loginUc    *login
	masterUc   *master
)

func initTest() {
	masterRepo = new(mockuserdbrepo.Masters)
	loginRepo = new(mockuserdbrepo.Logins)
	masterUc = newMaster(masterRepo)
	loginUc = newLogin(loginRepo, "")
}
