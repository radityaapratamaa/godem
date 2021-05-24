package goods

import (
	"bcg-test/domain/models"
	"bcg-test/domain/models/goods"
	goodsmodel "bcg-test/domain/models/goods"
	"bcg-test/infrastructure/database"
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type goodsIface interface {
	database.CRUD
	GetDetailByName(ctx context.Context, name string) (*goodsmodel.Goods, error)
	UpdateStock(ctx context.Context, action string, id int64, qty int) (*models.CUDResponse, error)
}

func generateLikeParams(data interface{}) string {
	return fmt.Sprintf("%%%%s%%", data)
}

func (repo *Repository) UpdateStock(ctx context.Context, action string, id int64, qty int) (*models.CUDResponse, error) {
	query := "UPDATE goods SET "

	switch action {
	case "sold":
		query += "qty=qty-?"
	case "add":
		query += "qty=qty+?"
	default:
		return nil, errors.New("Undefined action, only accept 'sold' and 'add'")
	}

	query += ", updated_at=now() WHERE id=?"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, qty, id)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
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

func (repo *Repository) GetDetailByName(ctx context.Context, name string) (*goodsmodel.Goods, error) {
	query := "SELECT * FROM goods WHERE deleted_at IS NULL and name = ?"

	query = repo.db.Slave.Rebind(query)
	var result *goodsmodel.Goods
	if err := repo.db.Slave.SelectContext(ctx, &result, query, name); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
	}
	return result, nil
}

func (repo *Repository) GetList(ctx context.Context, requestData interface{}) (interface{}, error) {
	goodsRequestData, valid := requestData.(*goods.GoodsRequest)
	if !valid {
		errMsg := errors.New("The Request Data must be goodsRequest struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.goods.GetList.ParsingInterface")
	}

	query := "SELECT * FROM goods WHERE deleted_at IS NULL"
	var params []interface{}
	if goodsRequestData.Query != "" {
		keyword := goodsRequestData.Query
		query = fmt.Sprintf("%s AND (title LIKE ? or body LIKE ?)", query)
		params = append(params, generateLikeParams(keyword), generateLikeParams(keyword))
	}
	if goodsRequestData.Author != "" {
		operand := "WHERE"
		if strings.Contains(query, "WHERE") {
			operand = "AND"
		}
		query = fmt.Sprintf("%s %s author = ?", query, operand)
		params = append(params, goodsRequestData.Author)
	}

	query = repo.db.Slave.Rebind(query)
	var result []*goods.Goods
	if err := repo.db.Slave.SelectContext(ctx, &result, query, params...); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
	}
	return result, nil
}

func (repo *Repository) GetDetailByID(ctx context.Context, id int64) (interface{}, error) {
	query := "SELECT * FROM goods WHERE id = ?"
	query = repo.db.Slave.Rebind(query)
	var result goods.Goods
	if err := repo.db.Slave.SelectContext(ctx, &result, query, id); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
	}
	return &result, nil
}

func (repo *Repository) CreateNew(ctx context.Context, requestData interface{}) (*models.CUDResponse, error) {
	goodsData, valid := requestData.(*goods.Goods)
	if !valid {
		errMsg := errors.New("The Request Data must be Goods struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.goods.GetList.ParsingInterface")
	}

	query := "INSERT INTO goods (id, sku, name, price, qty, created_at, updated_at, deleted_at) values (?, ?, ?, ?, ?, ?, ?, ?)"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, goodsData.ID, goodsData.SKU, goodsData.Name, goodsData.Price, goodsData.Qty, goodsData.CreatedAt, goodsData.UpdatedAt, goodsData.DeletedAt)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
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

func (repo *Repository) UpdateData(ctx context.Context, requestData interface{}, id int64) (*models.CUDResponse, error) {
	goodsData, valid := requestData.(*goods.Goods)
	if !valid {
		errMsg := errors.New("The Request Data must be Goods struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.goods.GetList.ParsingInterface")
	}

	query := "UPDATE goods SET SKU=?, name=?, price=?, qty=?, updated_at=? WHERE id=?"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, goodsData.SKU, goodsData.Name, goodsData.Price, goodsData.Qty, goodsData.UpdatedAt, goodsData.ID)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
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

func (repo *Repository) DeleteData(ctx context.Context, id int64) (*models.CUDResponse, error) {

	query := "UPDATE goods SET deleted_at=now() WHERE id=?"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, id)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
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
