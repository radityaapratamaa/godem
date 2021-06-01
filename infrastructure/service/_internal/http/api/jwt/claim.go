package jwt

import (
	"godem/lib/util/response"
	"godem/usecase/jwt"
	"net/http"
)

type Claims interface {
	GetJWTClaim(w http.ResponseWriter, r *http.Request)
}

type Claim struct {
	uc jwt.JWTs
}

func NewClaimHandler(uc jwt.JWTs) *Claim {
	return &Claim{uc: uc}
}

func (c *Claim) GetJWTClaim(w http.ResponseWriter, r *http.Request) {
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	data, err := c.uc.GetJWTClaim(r.Context())
	if err != nil {
		resp.InternalServerError(err)
		return
	}
	resp.Success(data)
	return
}
