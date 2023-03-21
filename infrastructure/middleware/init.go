package middleware

import (
	"godem/infrastructure/middleware/auth"
)

type Definitions struct {
	auth auth.Middlewares
}

func New(authMiddleware auth.Middlewares) *Definitions {
	return &Definitions{auth: authMiddleware}
}

func (a *Definitions) Auth() auth.Middlewares {
	return a.auth
}
