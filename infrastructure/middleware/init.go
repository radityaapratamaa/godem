package middleware

import "godem/usecase/jwt"

type Auth struct {
	jwt *JWT
}

func NewAuth(jwtUc jwt.JWTs) *Auth {
	return &Auth{jwt: NewJWT(jwtUc)}
}

func (a *Auth) JWT() *JWT {
	return a.jwt
}
