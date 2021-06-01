package context

import (
	"context"
	"godem/domain/models/jwt"
	"net/http"
)

type jwtContext struct{}

func SetJWT(req *http.Request, jwtData *jwt.Claim) {
	ctx := context.WithValue(req.Context(), jwtContext{}, jwtData)
	*req = *req.WithContext(ctx)
}

func GetJWT(ctx context.Context) *jwt.Claim {
	jwtData, ok := ctx.Value(jwtContext{}).(*jwt.Claim)
	if !ok {
		return nil
	}
	return jwtData
}
