package article

import (
	"article-test/domain/models"
	articlemodel "article-test/domain/models/article"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

type article interface{}

func (uc *Usecase) GetArticleList(ctx context.Context, requestData *articlemodel.ArticleRequest) (*models.SelectResponse, error) {
	cacheData, err := uc.cache.Get("article:list")
	if err != nil {
		return nil, err
	}
	var listData []*articlemodel.Article
	if cacheData == "" {
		listData, err = uc.db.GetArticleList(ctx, requestData)
		if err != nil {
			return nil, err
		}

		byteData, _ := json.Marshal(listData)
		if _, err := uc.cache.Set("article:list", string(byteData), 1000); err != nil {
			return nil, errors.Wrap(err, "usecase.article.GetArticleList")
		}
	} else {
		if err := json.Unmarshal([]byte(cacheData), &listData); err != nil {
			return nil, errors.Wrap(err, "usecase.article.GetList.UnmarshalRedis")
		}
	}

	result := new(models.SelectResponse)
	result.RequestParam = requestData
	result.Data = listData
	return result, nil
}
