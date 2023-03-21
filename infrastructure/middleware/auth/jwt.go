package auth

import (
	"godem/usecase/jwt"
	"net/http"

	"github.com/kodekoding/phastos/go/response"
	"github.com/kodekoding/phastos/go/router"

	"godem/infrastructure/context"
	"godem/lib/helper"
)

var validateProcess = (*JWT).validateProcess

type JWT struct {
	uc jwt.JWTs
}

func NewJWT(jwtUc jwt.Usecases) *JWT {
	return &JWT{uc: jwtUc.JWT()}
}

func (auth *JWT) ValidateGroup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse := validateProcess(auth, r)
		if jsonResponse != nil {
			jsonResponse.ErrorChecking(r)
			jsonResponse.Send(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (auth *JWT) Validate(next router.WrapperFunc) router.WrapperFunc {
	return func(writer http.ResponseWriter, request *http.Request) *response.JSON {
		jsonResponse := validateProcess(auth, request)
		if jsonResponse != nil {
			return jsonResponse
		}

		return next(writer, request)
	}
}

func (auth *JWT) validateProcess(r *http.Request) *response.JSON {
	token := helper.GetBearerToken(r)
	reqCtx := r.Context()
	claimData, err := auth.uc.DecodeToken(reqCtx, token)
	if err != nil {
		return response.NewJSON().ForbiddenResource(err)
	}
	context.SetJWT(r, claimData)
	return nil
}
