package user

import (
	"context"
	"godem/src/common"
	usermodel "godem/src/domain/models/user"
	"godem/src/usecase/user"
	"strconv"

	"github.com/kodekoding/phastos/v2/go/api"
	"github.com/kodekoding/phastos/v2/go/middlewares"
)

type Masters interface {
	GetList(req api.Request, ctx context.Context) *api.Response
	GetDetailById(req api.Request, ctx context.Context) *api.Response
	Insert(req api.Request, ctx context.Context) *api.Response
	Update(req api.Request, ctx context.Context) *api.Response
	Delete(req api.Request, ctx context.Context) *api.Response
	ResetDevice(req api.Request, ctx context.Context) *api.Response
	api.Controller
}

type master struct {
	*api.ControllerImpl
	uc user.Masters
}

func newMaster(ucMaster user.Masters) *master {
	return &master{uc: ucMaster}
}

func (m *master) GetList(req api.Request, ctx context.Context) *api.Response {
	var requestData usermodel.Request

	if err := req.GetQuery(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	data, err := m.uc.GetList(ctx, &requestData)
	if err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetData(data)
}

func (m *master) GetDetailById(req api.Request, ctx context.Context) *api.Response {
	id, err := strconv.Atoi(req.GetParams("id"))
	if err != nil {
		return api.NewResponse().SetError(api.UnprocessableEntity(err.Error(), common.ErrInputValidationCode))
	}

	data, err := m.uc.GetDetailById(ctx, id)
	if err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetData(data)
}

func (m *master) Insert(req api.Request, ctx context.Context) *api.Response {
	var requestData usermodel.Data
	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	data, err := m.uc.Insert(ctx, &requestData)
	if err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetData(data)
}

func (m *master) Update(req api.Request, ctx context.Context) *api.Response {
	var requestData usermodel.Data

	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(err)
	}

	id, err := strconv.ParseInt(req.GetParams("id"), 10, 64)

	if err != nil {
		return api.NewResponse().SetError(api.UnprocessableEntity(err.Error(), common.ErrInputValidationCode))
	}

	requestData.Id = id
	data, err := m.uc.Update(ctx, &requestData)
	if err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetData(data)
}

func (m *master) Delete(req api.Request, ctx context.Context) *api.Response {
	id, err := strconv.ParseInt(req.GetParams("id"), 10, 64)

	if err != nil {
		return api.NewResponse().SetError(api.UnprocessableEntity(err.Error(), common.ErrInputValidationCode))
	}

	data, err := m.uc.Delete(ctx, id)
	if err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetData(data)
}

func (m *master) ResetDevice(req api.Request, ctx context.Context) *api.Response {
	var requestData usermodel.ResetDeviceRequest
	if err := req.GetBody(&requestData); err != nil {
		return api.NewResponse().SetError(api.UnprocessableEntity(err.Error(), common.ErrInputValidationCode))
	}

	if _, err := m.uc.ResetDeviceId(ctx, requestData.UserId); err != nil {
		return api.NewResponse().SetError(err)
	}

	return api.NewResponse().SetMessage("device id successful to reset")
}

func (m *master) GetConfig() api.ControllerConfig {
	return api.ControllerConfig{
		Path: "/user",
		Routes: []api.Route{
			api.NewRoute("GET", m.GetList, api.WithMiddleware(middlewares.JWTAuth)),
			api.NewRoute("GET", m.GetDetailById, api.WithPath("/{id}"), api.WithMiddleware(middlewares.JWTAuth)),
			api.NewRoute("POST", m.ResetDevice, api.WithPath("/device/reset"), api.WithMiddleware(middlewares.JWTAuth)),
		},
	}
}
