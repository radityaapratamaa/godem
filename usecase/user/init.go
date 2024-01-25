package user

import (
	"embed"
	"github.com/kodekoding/phastos/v2/go/cache"
	"github.com/kodekoding/phastos/v2/go/database"
	"github.com/kodekoding/phastos/v2/go/mail"
	"godem/infrastructure/database/user"
)

type Usecases interface {
	Master() Masters
	Auth() Auths
}

type Usecase struct {
	master Masters
	auth   Auths
}

func New(
	userRepo user.Repositories,
	smtp mail.SMTPs,
	templateFolder embed.FS,
	trx database.Transactions,
	cache cache.Caches,
) *Usecase {
	return &Usecase{
		master: newMaster(userRepo.Master(), trx),
		auth:   newAuth(userRepo.Auth(), smtp, templateFolder, cache),
	}
}

func (u *Usecase) Master() Masters {
	return u.master
}

func (u *Usecase) Auth() Auths {
	return u.auth
}
