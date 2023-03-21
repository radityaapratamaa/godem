package auth

import (
	"net/http"

	"github.com/kodekoding/phastos/go/router"
)

type (
	Validates interface {
		ValidateGroup(next http.Handler) http.Handler
		Validate(next router.WrapperFunc) router.WrapperFunc
	}
	Middlewares interface {
		JWT() Validates
	}
)
