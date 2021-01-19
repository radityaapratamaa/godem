package article

import (
	"article-test/infrastructure/cache"
	articledb "article-test/infrastructure/database/article"
)

type Usecases interface {
}

type Usecase struct {
	db    articledb.Repositories
	cache cache.Caches
}

func New(db articledb.Repositories, cache cache.Caches) *Usecase {
	return &Usecase{
		db:    db,
		cache: cache,
	}
}
