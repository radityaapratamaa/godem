package database

import (
	"bcg-test/domain/models"
	"context"
)

type CRUD interface {
	GetList(ctx context.Context, requestData interface{}) (interface{}, error)
	GetDetailByID(ctx context.Context, id int64) (interface{}, error)
	CreateNew(ctx context.Context, requestData interface{}) (*models.CUDResponse, error)
	UpdateData(ctx context.Context, requestData interface{}, id int64) (*models.CUDResponse, error)
	DeleteData(ctx context.Context, id int64) (*models.CUDResponse, error)
}
