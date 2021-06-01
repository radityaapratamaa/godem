package jwt

import "godem/usecase/jwt"

type Handler struct {
	claim Claims
}

func NewHandler(jwtClaimUc jwt.JWTs) *Handler {
	return &Handler{claim: NewClaimHandler(jwtClaimUc)}
}

func (h *Handler) Claim() Claims {
	return h.claim
}
