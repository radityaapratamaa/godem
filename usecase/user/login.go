package user

import (
	"context"
	"encoding/base64"
	"godem/lib/helper"

	"github.com/pkg/errors"

	"godem/domain/models/user"
	userdb "godem/infrastructure/database/user"
)

type Logins interface {
	Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.LoginResponse, error)
}

type login struct {
	jwtPubKey string
	repo      userdb.Logins
}

var (
	generateJWTToken   = helper.GenerateJWTToken
	decodeBase64String = (base64.StdEncoding).DecodeString
)

func newLogin(repo userdb.Logins, jwtPubKey string) *login {
	return &login{repo: repo, jwtPubKey: jwtPubKey}
}

func (l *login) Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.LoginResponse, error) {
	realPass, err := decodeBase64String(requestData.Password)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.login.Authenticate.DecodePassword")
	}
	requestData.Password = string(realPass)
	data, err := l.repo.Authenticate(ctx, requestData)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.login.Authenticate.DBCheck")
	}

	jwtToken, expired, err := generateJWTToken(l.jwtPubKey, data)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.login.Authenticate.GenerateJWTToken")
	}

	response := new(user.LoginResponse)
	response.Token = jwtToken
	response.Expired = expired
	return response, nil
}
