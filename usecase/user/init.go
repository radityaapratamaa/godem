package user

import "godem/infrastructure/database/user"

type Usecase struct {
	master Masters
}

func New(masterRepo user.Masters) *Usecase {
	return &Usecase{
		master: NewMaster(masterRepo),
	}
}

func (u *Usecase) Master() Masters {
	return u.master
}
