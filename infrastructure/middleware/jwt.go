package middleware

import (
	"godem/infrastructure/context"
	"godem/lib/helper"
	"godem/lib/util/response"
	"godem/usecase/jwt"
	"net/http"
)

type JWT struct {
	uc jwt.JWTs
}

func NewJWT(jwtUc jwt.JWTs) *JWT {
	return &JWT{uc: jwtUc}
}

func (j *JWT) ValidateGroup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse := j.validateProcess(r)
		if jsonResponse != nil {
			jsonResponse.Send(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (j *JWT) Validate(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonResponse := j.validateProcess(request)
		if jsonResponse != nil {
			jsonResponse.Send(writer)
			return
		}

		next(writer, request)
	}
}

func (j *JWT) validateProcess(r *http.Request) *response.JSON {
	token := helper.GetBearerToken(r)
	reqCtx := r.Context()
	resp := response.NewJSON()
	claimData, err := j.uc.DecodeToken(reqCtx, token)
	if err != nil {
		return resp.ForbiddenResource(err)
	}
	context.SetJWT(r, claimData)
	return nil
}
