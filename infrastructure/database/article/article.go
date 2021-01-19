package article

import (
	articlemodel "article-test/domain/models/article"
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type article interface {
	GetArticleList(ctx context.Context, requestData *articlemodel.ArticleRequest) ([]*articlemodel.Article, error)
}

func generateLikeParams(data interface{}) string {
	return fmt.Sprintf("%%%%s%%", data)
}
func (repo *Repository) GetArticleList(ctx context.Context, requestData *articlemodel.ArticleRequest) ([]*articlemodel.Article, error) {

	query := "SELECT * FROM article"
	var params []interface{}
	if requestData.Query != "" {
		keyword := requestData.Query
		query = fmt.Sprintf("%s WHERE (title LIKE ? or body LIKE ?)", query)
		params = append(params, generateLikeParams(keyword), generateLikeParams(keyword))
	}
	if requestData.Author != "" {
		operand := "WHERE"
		if strings.Contains(query, "WHERE") {
			operand = "AND"
		}
		query = fmt.Sprintf("%s %s author = ?", query, operand)
		params = append(params, requestData.Author)
	}

	query = repo.db.Slave.Rebind(query)
	var result []*articlemodel.Article
	if err := repo.db.Slave.SelectContext(ctx, &result, query, params...); err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.article.GetList")
	}
	return result, nil
}
