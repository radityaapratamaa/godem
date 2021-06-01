package user

import "godem/infrastructure/database/user"

type Usecase struct {
	master Masters
	login  Logins
}

func New(masterRepo user.Masters, loginRepo user.Logins, jwtPubKey string) *Usecase {
	return &Usecase{
		master: NewMaster(masterRepo),
		login:  NewLogin(loginRepo, jwtPubKey),
	}
}

func (u *Usecase) Master() Masters {
	return u.master
}

func (u *Usecase) Login() Logins {
	return u.login
}
