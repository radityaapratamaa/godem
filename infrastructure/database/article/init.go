package article

import "article-test/domain/models"

type Repositories interface {
	article
}

type Repository struct {
	db *models.Database
}

func New(db *models.Database) *Repository {
	return &Repository{
		db: db,
	}
}
