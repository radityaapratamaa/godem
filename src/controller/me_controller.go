package controller

import (
	"context"
	"github.com/kodekoding/phastos/v2/go/api"
	"github.com/kodekoding/phastos/v2/go/middlewares"
)

type me struct {
	*api.ControllerImpl
}

func NewMe() *me {
	return &me{}
}

func (m *me) GetData(_ api.Request, ctx context.Context) *api.Response {
	return api.NewResponse().SetMessage("its me")
}

func (m *me) GetConfig() api.ControllerConfig {
	return api.ControllerConfig{
		Path:        "/me",
		Middlewares: m.JoinMiddleware(middlewares.JWTAuth),
		Routes: []api.Route{
			api.NewRoute("GET", m.GetData),
		},
	}
}
