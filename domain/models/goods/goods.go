package goods

type Goods struct {
	ID        int64   `json:"id"`
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

type GoodsRequest struct {
	Query  string `json:"query" schema:"query"`
	Author string `json:"author" schema:"author"`
}
