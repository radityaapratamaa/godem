package user

import (
	"context"
	"fmt"
	"github.com/kodekoding/phastos/v2/go/api"
	"github.com/kodekoding/phastos/v2/go/middlewares"
	"godem/domain/models/user"
	useruc "godem/usecase/user"
)

type Auths interface {
	Authenticate(req api.Request, ctx context.Context) *api.Response
	api.Controller
}

type auth struct {
	uc useruc.Auths
}

func newAuth(loginUc useruc.Auths) *auth {
	return &auth{uc: loginUc}
}

func (l *auth) Authenticate(req api.Request, ctx context.Context) *api.Response {
	var requestData user.LoginRequest
	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	result, err := l.uc.Authenticate(ctx, &requestData)

	if err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetData(result)
}

func (l *auth) ForgotPassword(req api.Request, ctx context.Context) *api.Response {
	var requestData user.LoginRequest
	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	if err := l.uc.ForgotPassword(ctx, requestData.Email); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage(fmt.Sprintf("Email sent to %s", requestData.Email))
}

func (l *auth) ResetPassword(req api.Request, ctx context.Context) *api.Response {
	var requestData user.ResetPasswordRequest
	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	if err := l.uc.ResetPassword(ctx, &requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage("password was successfully reset")
}

func (l *auth) UpdatePassword(req api.Request, ctx context.Context) *api.Response {
	var requestData user.ChangePasswordRequest
	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	if err := l.uc.UpdatePassword(ctx, &requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage("password was successfully changed")
}

func (l *auth) ResendToken(req api.Request, ctx context.Context) *api.Response {
	email := req.GetParams("email")
	if email == "" {
		return api.NewResponse().SetError(api.BadRequest("email cannot be blank", "BLANK_EMAIL"))
	}

	if err := l.uc.ResendToken(ctx, email); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage(fmt.Sprintf("Email sent to %s", email))

}

func (l *auth) VerifyToken(req api.Request, ctx context.Context) *api.Response {
	var requestedData user.VerifyTokenRequest
	if err := req.GetBody(&requestedData); err != nil {
		return api.NewResponse().SetError(err)
	}

	if err := l.uc.VerifyToken(ctx, requestedData.Email, requestedData.OTP); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage("Success")

}

func (l *auth) Logout(_ api.Request, ctx context.Context) *api.Response {

	if _, err := l.uc.Logout(ctx); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage("logged out")

}

func (l *auth) GetConfig() api.ControllerConfig {
	return api.ControllerConfig{
		Path: "/auth",
		Routes: []api.Route{
			api.NewRoute("POST", l.Authenticate),
			api.NewRoute("POST", l.ForgotPassword, api.WithPath("/forgot")),
			api.NewRoute("PUT", l.ResetPassword, api.WithPath("/reset")),
			api.NewRoute("GET", l.ResendToken, api.WithPath("/token/resend/{email}")),
			api.NewRoute("POST", l.VerifyToken, api.WithPath("/token/verify")),
			api.NewRoute("POST", l.Logout, api.WithPath("/logout"), api.WithMiddleware(middlewares.JWTAuth)),
			api.NewRoute("PUT", l.UpdatePassword, api.WithMiddleware(middlewares.JWTAuth)),
		},
	}
}
