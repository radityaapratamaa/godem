package user

import (
	"context"
	usermodel "godem/domain/models/user"

	"github.com/pkg/errors"

	"godem/domain/models"
	"godem/infrastructure/common"
	"godem/infrastructure/database/user"
)

type Masters interface {
	common.CRUD
}

type master struct {
	repo user.Masters
}

func newMaster(repo user.Masters) *master {
	return &master{repo: repo}
}

func (m *master) GetList(ctx context.Context, requestData interface{}) (interface{}, error) {
	list, err := m.repo.GetList(ctx, requestData)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.GetList")
	}
	result := new(models.SelectResponse)
	result.RequestParam = requestData
	result.Data = list
	return result, nil
}

func (m *master) GetDetailByID(ctx context.Context, id int64) (interface{}, error) {
	data, err := m.repo.GetDetailByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.GetDetailByID")
	}

	return data, err
}

func (m *master) CreateNew(ctx context.Context, requestData interface{}) (*models.CUDResponse, error) {
	return m.repo.CreateNew(ctx, requestData)
}

func (m *master) UpdateData(ctx context.Context, requestData interface{}, id int64) (*models.CUDResponse, error) {
	reqData, valid := requestData.(*usermodel.Users)
	if !valid {

		return nil, errors.Wrap(errors.New("Failed Parse Interface"), "usecase.user.master.ParseInterface")
	}
	reqData.ID = id
	return m.repo.UpdateData(ctx, requestData, id)
}

func (m *master) DeleteData(ctx context.Context, id int64) (*models.CUDResponse, error) {
	return m.repo.DeleteData(ctx, id)
}
