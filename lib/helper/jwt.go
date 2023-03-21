package helper

import (
	"github.com/pkg/errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"godem/domain/models/user"
)

type ClaimData struct {
	*user.Users
	jwt.RegisteredClaims
}

func GenerateJWTToken(signingKey string, data *user.Users) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claimData := new(ClaimData)
	claimData.Users = data
	claimData.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Issuer:    "godem-backend",
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claimData)
	token, err := tokenClaims.SignedString(signingKey)
	if err != nil {
		return "", errors.Wrap(err, "lib.helper.jwt.GenerateJWTToken.SignedString")
	}
	return token, nil
}
