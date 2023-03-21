package context

import (
	"context"
	"github.com/kodekoding/phastos/go/apps"
	"godem/domain/models/jwt"
	"net/http"
)

type (
	jwtContext       struct{}
	slackAppsContext struct{}
)

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

func SetSlackApps(req *http.Request, slack apps.Slacks) {
	ctx := context.WithValue(req.Context(), slackAppsContext{}, slack)
	*req = *req.WithContext(ctx)
}

func GetSlackApps(ctx context.Context) apps.Slacks {
	slackPlatform, valid := ctx.Value(slackAppsContext{}).(apps.Slacks)
	if !valid {
		return nil
	}
	return slackPlatform
}
