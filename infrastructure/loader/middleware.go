package loader

import (
	"godem/infrastructure/middleware"
	"godem/infrastructure/middleware/auth"
)

func InitMiddleware(authModules auth.Middlewares) *middleware.Definitions {
	return middleware.New(authModules)
}
