package user

import "godem/infrastructure/database/user"

type Usecases interface {
	Master() Masters
	Login() Logins
}

type Usecase struct {
	master Masters
	login  Logins
}

func New(userRepo user.Repositories, jwtPubKey string) *Usecase {
	return &Usecase{
		master: newMaster(userRepo.Master()),
		login:  newLogin(userRepo.Login(), jwtPubKey),
	}
}

func (u *Usecase) Master() Masters {
	return u.master
}

func (u *Usecase) Login() Logins {
	return u.login
}
