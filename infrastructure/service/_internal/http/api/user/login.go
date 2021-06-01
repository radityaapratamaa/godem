package user

import (
	"encoding/json"
	"godem/domain/models/user"
	"godem/lib/util/response"
	useruc "godem/usecase/user"
	"net/http"
)

type Logins interface {
	Authenticate(w http.ResponseWriter, r *http.Request)
}

type Login struct {
	uc useruc.Logins
}

func NewLogin(loginUc useruc.Logins) *Login {
	return &Login{uc: loginUc}
}

func (l *Login) Authenticate(w http.ResponseWriter, r *http.Request) {
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	var requestData user.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		resp.BadRequest(err)
		return
	}

	result, err := l.uc.Authenticate(r.Context(), &requestData)
	if err != nil {
		resp.InternalServerError(err)
		return
	}

	resp.Success(result)
	return
}
