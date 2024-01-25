package user

import "github.com/kodekoding/phastos/v2/go/database"

type Repositories interface {
	Auth() Auths
	Master() Masters
}

type Repository struct {
	auth   Auths
	master Masters
}

func New(db database.ISQL) Repositories {
	return &Repository{auth: newAuth(db), master: newMaster(db)}
}

func (r Repository) Auth() Auths {
	return r.auth
}

func (r Repository) Master() Masters {
	return r.master
}
