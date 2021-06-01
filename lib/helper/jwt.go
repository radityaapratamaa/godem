package helper

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func GenerateJWTToken(signingKey string, data interface{}) (string, int64, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claimData := token.Claims.(jwt.MapClaims)
	byteData, err := json.Marshal(data)
	if err != nil {
		return "", 0, errors.Wrap(err, "usecase.user.login.setJWTClaims.MarshalData")
	}

	if err := json.Unmarshal(byteData, &claimData); err != nil {
		return "", 0, errors.Wrap(err, "usecase.auth.jwt.setJWTClaims.UnmarshalToClaims")
	}

	claimData["auth_type"] = "login"
	expired := time.Now().Add(time.Hour * 168).Unix()

	claimData["expired"] = expired
	jwtToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", 0, errors.Wrap(err, "usecase.auth.jwt.setJWTClaims.SignedString")
	}
	return jwtToken, expired, nil
}
