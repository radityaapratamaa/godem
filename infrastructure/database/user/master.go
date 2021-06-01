package user

import (
	"context"
	"fmt"
	"godem/lib/util/database"

	"github.com/pkg/errors"

	"godem/domain/models"
	"godem/domain/models/user"
	"godem/infrastructure/common"
)

type Masters interface {
	common.CRUD
}

type Master struct {
	db *database.DB
}

func NewMaster(db *database.DB) *Master {
	return &Master{db: db}
}

func generateLikeParams(data interface{}) string {
	return fmt.Sprintf("%%%s%%", data)
}

func (repo *Master) GetList(ctx context.Context, requestData interface{}) (interface{}, error) {
	goodsRequestData, valid := requestData.(*user.UsersRequest)
	if !valid {
		errMsg := errors.New("The Request Data must be goodsRequest struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.goods.GetList.ParsingInterface")
	}

	query := "SELECT id, username, address, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL"
	var params []interface{}
	if goodsRequestData.Query != "" {
		keyword := goodsRequestData.Query
		query = fmt.Sprintf("%s AND (username LIKE ? or address LIKE ?)", query)
		params = append(params, generateLikeParams(keyword), generateLikeParams(keyword))
	}

	query = repo.db.Rebind(query)
	var result []*user.Users
	if err := repo.db.Follower.SelectContext(ctx, &result, query, params...); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.users.GetList")
	}
	return result, nil
}

func (repo *Master) GetDetailByID(ctx context.Context, id int64) (interface{}, error) {
	query := "SELECT id, username, address, created_at, updated_at, deleted_at FROM users WHERE id = ? and deleted_at is NULL"
	query = repo.db.Rebind(query)
	var result user.Users
	if err := repo.db.Follower.GetContext(ctx, &result, query, id); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.users.GetList")
	}
	return &result, nil
}

func (repo *Master) CreateNew(ctx context.Context, requestData interface{}) (*models.CUDResponse, error) {
	userData, valid := requestData.(*user.Users)
	if !valid {
		errMsg := errors.New("The Request Data must be Goods struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.goods.GetList.ParsingInterface")
	}

	query := "INSERT INTO users (id, username, passwd, address, created_at, updated_at, deleted_at) values (?, ?, md5(?), ?, now(), ?, ?)"
	query = repo.db.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, userData.ID, userData.Username, userData.Password, userData.Address, userData.UpdatedAt, userData.DeletedAt)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.users.GetList")
	}

	result := new(models.CUDResponse)
	if lastInsertId, err := sqlResult.LastInsertId(); err == nil {
		result.LastInsertID = lastInsertId
	}
	if rowsAff, err := sqlResult.RowsAffected(); err == nil {
		result.RowsAffected = rowsAff
	}

	result.Status = true
	return result, nil
}

func (repo *Master) UpdateData(ctx context.Context, requestData interface{}, id int64) (*models.CUDResponse, error) {
	userData, valid := requestData.(*user.Users)
	if !valid {
		errMsg := errors.New("The Request Data must be Goods struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.goods.GetList.ParsingInterface")
	}

	query := "UPDATE users SET username=?, passwd=?, address=?, updated_at=now() WHERE id=?"
	query = repo.db.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, userData.Username, userData.Password, userData.Address, userData.ID)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.users.GetList")
	}

	result := new(models.CUDResponse)
	if lastInsertId, err := sqlResult.LastInsertId(); err == nil {
		result.LastInsertID = lastInsertId
	}
	if rowsAff, err := sqlResult.RowsAffected(); err == nil {
		result.RowsAffected = rowsAff
	}

	result.Status = true
	return result, nil
}

func (repo *Master) DeleteData(ctx context.Context, id int64) (*models.CUDResponse, error) {

	query := "UPDATE users SET deleted_at=now() WHERE id=?"
	query = repo.db.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, id)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.users.GetList")
	}

	result := new(models.CUDResponse)
	if lastInsertId, err := sqlResult.LastInsertId(); err == nil {
		result.LastInsertID = lastInsertId
	}
	if rowsAff, err := sqlResult.RowsAffected(); err == nil {
		result.RowsAffected = rowsAff
	}

	result.Status = true
	return result, nil
}
