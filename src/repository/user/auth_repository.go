package user

import (
	"context"
	common2 "godem/src/common"
	user2 "godem/src/domain/models/user"
	"os"

	"github.com/kodekoding/phastos/v2/go/common"
	"github.com/kodekoding/phastos/v2/go/database"
	"github.com/kodekoding/phastos/v2/go/database/action"
	"github.com/pkg/errors"
)

type Auths interface {
	common.RepoCRUD
	Authenticate(ctx context.Context, requestData *user2.LoginRequest) (*user2.Data, error)
	ResetPassword(ctx context.Context, requestData *user2.ResetPasswordRequest) error
	UpdatePassword(ctx context.Context, requestData *user2.ChangePasswordRequest) error
}

type auth struct {
	*action.Base
	db     database.ISQL
	encKey string
}

func newAuth(db database.ISQL) *auth {
	return &auth{
		db:     db,
		Base:   action.NewBase(db, "users"),
		encKey: os.Getenv(common2.EnvServerEncryptionKey),
	}
}

func (l *auth) Authenticate(ctx context.Context, requestData *user2.LoginRequest) (*user2.Data, error) {
	query := `select id, email, name, device_id, firebase_id from users where LOWER(email) = LOWER(?) and pgp_sym_decrypt(password::bytea, ?) = ? and deleted_at IS NULL AND active_at IS NOT NULL`
	query = l.db.Rebind(query)

	var result user2.Data
	if err := l.db.GetContext(ctx, &result, query, requestData.Email, l.encKey, requestData.Password); err != nil {
		return nil, errors.Wrap(err, "repository.user.auth.Authenticate")
	}
	return &result, nil
}

func (l *auth) ResetPassword(ctx context.Context, requestData *user2.ResetPasswordRequest) error {
	query := `UPDATE users SET activation_code = null, activation_expired_at = null, active_at = now(), updated_at = now(), password = pgp_sym_encrypt(?, ?) WHERE LOWER(email) = LOWER(?) AND activation_code = ?`
	query = l.db.Rebind(query)

	if _, err := l.db.ExecContext(ctx, query, requestData.NewPassword, l.encKey, requestData.Email, requestData.OTP); err != nil {
		return errors.Wrap(err, "repository.user.auth.ResetPassword.ExecContext")
	}
	return nil
}

func (l *auth) UpdatePassword(ctx context.Context, requestData *user2.ChangePasswordRequest) error {
	query := `UPDATE users SET updated_at = now(), password = pgp_sym_encrypt(?, ?) WHERE LOWER(email) = LOWER(?) AND pgp_sym_decrypt(password::bytea, ?) = ? AND active_at IS NOT NULL`
	query = l.db.Rebind(query)

	if _, err := l.db.ExecContext(ctx, query, requestData.NewPassword, l.encKey, requestData.Email, l.encKey, requestData.OldPassword); err != nil {
		return errors.Wrap(err, "repository.user.auth.UpdatePassword.ExecContext")
	}
	return nil
}
