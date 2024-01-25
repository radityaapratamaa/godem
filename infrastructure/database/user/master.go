package user

import (
	"github.com/kodekoding/phastos/v2/go/common"
	"github.com/kodekoding/phastos/v2/go/database"
	"github.com/kodekoding/phastos/v2/go/database/action"
)

type Masters interface {
	common.RepoCRUD
}

type master struct {
	*action.Base
	db database.ISQL
}

func newMaster(db database.ISQL) *master {
	return &master{
		db:   db,
		Base: action.NewBase(db, "users"),
	}
}
