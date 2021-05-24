package goods

import "bcg-test/domain/models"

type Repositories interface {
	goodsIface
}

type Repository struct {
	db *models.Database
}

func New(db *models.Database) *Repository {
	return &Repository{
		db: db,
	}
}
