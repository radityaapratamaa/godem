package jwt

import (
	"github.com/kodekoding/phastos/go/response"

	"godem/usecase/jwt"
	"net/http"
)

type Claims interface {
	GetJWTClaim(w http.ResponseWriter, r *http.Request) *response.JSON
}

type Claim struct {
	uc jwt.JWTs
}

func NewClaimHandler(uc jwt.JWTs) *Claim {
	return &Claim{uc: uc}
}

func (c *Claim) GetJWTClaim(w http.ResponseWriter, r *http.Request) *response.JSON {

	data, err := c.uc.GetJWTClaim(r.Context())
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}
	return response.NewJSON().Success(data)
}
