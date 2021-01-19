package article

type Article struct {
	ID        int64  `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type ArticleRequest struct {
	Query  string `json:"query" schema:"query"`
	Author string `json:"author" schema:"author"`
}
