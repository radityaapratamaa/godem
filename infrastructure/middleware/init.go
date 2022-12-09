package middleware

import (
	"godem/infrastructure/middleware/auth"
	"godem/usecase/jwt"
)

type Definitions struct {
	auth auth.Middlewares
}

type Modules struct {
	JWTUc jwt.Usecases
}

func New(middlewareModules *Modules) *Definitions {
	return &Definitions{auth: auth.New(middlewareModules.JWTUc)}
}

func (a *Definitions) Auth() auth.Middlewares {
	return a.auth
}
