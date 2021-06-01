package user

import (
	"godem/infrastructure/database/mocks/user"
)

var (
	master *user.Masters
	login  *user.Logins
	uc     *Usecase
)

func initTest() {
	master = new(user.Masters)
	login = new(user.Logins)

	uc = New(master, login, "")
}
