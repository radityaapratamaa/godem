package user

import (
	"context"
	"godem/domain/models/user"
	"godem/lib/util/database"

	"github.com/pkg/errors"
)

type Logins interface {
	Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.Users, error)
}

type Login struct {
	db *database.DB
}

func NewLogin(db *database.DB) *Login {
	return &Login{db: db}
}

func (l *Login) Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.Users, error) {
	query := `SELECT * FROM users WHERE username = ? and passwd = md5(?) and deleted_at IS NULL`
	query = l.db.Rebind(query)

	var result user.Users
	if err := l.db.Follower.GetContext(ctx, &result, query, requestData.Username, requestData.Password); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.user.login.Authenticate")
	}
	return &result, nil
}
