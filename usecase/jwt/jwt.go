package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	jwtmodel "godem/domain/models/jwt"
	ctxinternal "godem/infrastructure/context"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/pkg/errors"
)

type JWTs interface {
	DecodeToken(ctx context.Context, token string) (*jwtmodel.Claim, error)
	GetJWTClaim(ctx context.Context) (*jwtmodel.Claim, error)
}

type jwtU struct {
	signingKey string
	keyFunc    jwt.Keyfunc
}

func NewJWT(signingKey string) *jwtU {
	return &jwtU{signingKey: signingKey}
}

func (u *jwtU) DecodeToken(ctx context.Context, token string) (*jwtmodel.Claim, error) {
	if u.signingKey == "" {
		err := errors.New("JWT Signing Key is nil")
		return nil, errors.Wrap(err, "usecase.jwt.DecodeToken.CheckSigningKey")
	}

	u.keyFunc = func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
		}
		return []byte(u.signingKey), nil
	}

	tokenClient := strings.TrimSpace(token)
	data := jwt.MapClaims{}
	tokenData := new(jwt.Token)
	var errToken error
	tokenData, errToken = jwt.ParseWithClaims(tokenClient, data, u.keyFunc)
	if errToken != nil {
		return nil, errors.Wrap(errToken, "usecase.jwt.DecodeToken.ParseTokenClaims")
	}

	if !tokenData.Valid {
		invalidTokenError := errors.New("Token is invalid")
		return nil, errors.Wrap(invalidTokenError, "usecase.jwt.DecodeToken.CheckValidity")
	}

	claimByte, _ := json.Marshal(tokenData.Claims)
	var result jwtmodel.Claim
	if err := json.Unmarshal(claimByte, &result); err != nil {
		return nil, errors.Wrap(err, "usecase.jwt.DecodeToken.UnmarshalClaimToStruct")
	}
	result.Token = token
	return &result, nil
}

func (u *jwtU) GetJWTClaim(ctx context.Context) (*jwtmodel.Claim, error) {
	jwtClaimData := ctxinternal.GetJWT(ctx)
	if jwtClaimData == nil {
		notFoundJWT := errors.New("JWT is not valid")
		return nil, errors.Wrap(notFoundJWT, "usecase.jwt.GetJWTDataFromContext")
	}
	return jwtClaimData, nil
}
