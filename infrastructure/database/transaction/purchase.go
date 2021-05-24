package transaction

import (
	"bcg-test/domain/models"
	"bcg-test/domain/models/transaction"
	"bcg-test/infrastructure/database"
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type purchase interface {
	database.CRUD
}

func generateLikeParams(data interface{}) string {
	return fmt.Sprintf("%%%%s%%", data)
}

func (repo *Repository) GetList(ctx context.Context, requestData interface{}) (interface{}, error) {
	purchaseRequestData, valid := requestData.(*transaction.PurchaseRequest)
	if !valid {
		errMsg := errors.New("The Request Data must be PurchaseRequest struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.purchase.GetList.ParsingInterface")
	}

	query := "SELECT * FROM purchase WHERE deleted_at IS NULL"
	var params []interface{}
	if purchaseRequestData.Query != "" {
		keyword := purchaseRequestData.Query
		query = fmt.Sprintf("%s AND (title LIKE ? or body LIKE ?)", query)
		params = append(params, generateLikeParams(keyword), generateLikeParams(keyword))
	}

	query = repo.db.Slave.Rebind(query)
	var result []*transaction.Purchase
	if err := repo.db.Slave.SelectContext(ctx, &result, query, params...); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.transaction.purchase.GetList")
	}
	return result, nil
}

func (repo *Repository) GetDetailByID(ctx context.Context, id int64) (interface{}, error) {
	query := "SELECT * FROM purchase WHERE id = ?"
	query = repo.db.Slave.Rebind(query)
	var result transaction.Purchase
	if err := repo.db.Slave.SelectContext(ctx, &result, query, id); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.transaction.purchase.GetList")
	}
	return &result, nil
}

func (repo *Repository) CreateNew(ctx context.Context, requestData interface{}) (*models.CUDResponse, error) {
	purchaseData, valid := requestData.(*transaction.Purchase)
	if !valid {
		errMsg := errors.New("The Request Data must be Purchase struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.purchase.GetList.ParsingInterface")
	}

	query := "INSERT INTO purchase (id, goods_id, qty, purchase_date, created_at, updated_at, deleted_at) values (?, ?, ?, ?, ?, ?, ?)"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, purchaseData.ID, purchaseData.GoodsID, purchaseData.Qty, purchaseData.PurchaseDate, purchaseData.CreatedAt, purchaseData.UpdatedAt, purchaseData.DeletedAt)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.transaction.purchase.GetList")
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
	purchaseData, valid := requestData.(*transaction.Purchase)
	if !valid {
		errMsg := errors.New("The Request Data must be Purchase struct")
		return nil, errors.Wrap(errMsg, "infrastructure.database.purchase.GetList.ParsingInterface")
	}

	query := "UPDATE purchase SET qty=?, updated_at=? WHERE id=?"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, purchaseData.Qty, purchaseData.UpdatedAt, purchaseData.ID)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.transaction.purchase.UpdateData")
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

	query := "UPDATE purchase SET deleted_at=now() WHERE id=?"
	query = repo.db.Slave.Rebind(query)

	sqlResult, err := repo.db.Master.ExecContext(ctx, query, id)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.transaction.purchase.DeleteData")
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
