package transaction

import (
	"bcg-test/infrastructure/cache"
	"bcg-test/infrastructure/database/goods"
	"bcg-test/infrastructure/database/transaction"
)

type Usecases interface {
	purchase
}

type Usecase struct {
	db      transaction.Repositories
	goodsDB goods.Repositories
	cache   cache.Caches
}

func New(db transaction.Repositories, cache cache.Caches) *Usecase {
	return &Usecase{
		db:    db,
		cache: cache,
	}
}
