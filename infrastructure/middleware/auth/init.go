package auth

import "godem/usecase/jwt"

type (
	Middleware struct {
		jwt Validates
	}
)

func New(jwtUc jwt.Usecases) *Middleware {
	return &Middleware{jwt: NewJWT(jwtUc)}
}

func (m *Middleware) JWT() Validates {
	return m.jwt
}
