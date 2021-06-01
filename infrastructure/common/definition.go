package common

import (
	"context"
	"godem/domain/models"
	"net/http"
)

type CRUD interface {
	GetList(ctx context.Context, requestData interface{}) (interface{}, error)
	GetDetailByID(ctx context.Context, id int64) (interface{}, error)
	CreateNew(ctx context.Context, requestData interface{}) (*models.CUDResponse, error)
	UpdateData(ctx context.Context, requestData interface{}, id int64) (*models.CUDResponse, error)
	DeleteData(ctx context.Context, id int64) (*models.CUDResponse, error)
}

type CRUDHandler interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetailByID(w http.ResponseWriter, r *http.Request)
	CreateNew(w http.ResponseWriter, r *http.Request)
	UpdateData(w http.ResponseWriter, r *http.Request)
	DeleteData(w http.ResponseWriter, r *http.Request)
}
