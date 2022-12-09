package user

import (
	"context"
	"github.com/kodekoding/phastos/go/common"
	"github.com/kodekoding/phastos/go/database"
	"github.com/kodekoding/phastos/go/database/action"
	"godem/domain/models/user"

	"github.com/pkg/errors"
)

type Logins interface {
	common.RepoCRUD
	Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.Users, error)
}

type login struct {
	*action.Base
	db *database.SQL
}

func newLogin(db *database.SQL) *login {
	return &login{
		db:   db,
		Base: action.NewBase(db, "login"),
	}
}

func (l *login) Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.Users, error) {
	query := `SELECT * FROM users WHERE username = ? and passwd = md5(?) and deleted_at IS NULL`
	query = l.db.Rebind(query)

	var result user.Users
	if err := l.db.Follower.GetContext(ctx, &result, query, requestData.Username, requestData.Password); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.user.login.Authenticate")
	}
	return &result, nil
}
