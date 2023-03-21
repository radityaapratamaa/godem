package user

import "github.com/kodekoding/phastos/go/database"

type Repositories interface {
	Login() Logins
	Master() Masters
}

type Repository struct {
	login  Logins
	master Masters
}

func New(db *database.SQL) Repositories {
	return &Repository{login: newLogin(db), master: newMaster(db)}
}

func (r Repository) Login() Logins {
	return r.login
}

func (r Repository) Master() Masters {
	return r.master
}
