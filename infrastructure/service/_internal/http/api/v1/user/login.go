package user

import (
	"encoding/json"
	"net/http"

	"github.com/kodekoding/phastos/go/response"

	"godem/domain/models/user"
	useruc "godem/usecase/user"
)

type Logins interface {
	Authenticate(w http.ResponseWriter, r *http.Request) *response.JSON
}

type login struct {
	uc useruc.Logins
}

func newLogin(loginUc useruc.Logins) *login {
	return &login{uc: loginUc}
}

func (l *login) Authenticate(w http.ResponseWriter, r *http.Request) *response.JSON {
	var requestData user.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return response.NewJSON().BadRequest(err)
	}

	result, err := l.uc.Authenticate(r.Context(), &requestData)
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}

	return response.NewJSON().Success(result)
}
