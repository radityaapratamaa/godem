package transaction

import (
	"bcg-test/domain/models"
)

type Purchase struct {
	ID           int64   `json:"id"`
	GoodsID      int64   `json:"goods_id"`
	Qty          int     `json:"qty"`
	PurchaseDate string  `json:"purchase_date"`
	TotalPrice   float64 `json:"total_price"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"udpated_at"`
	DeletedAt    string  `json:"deleted_at"`
}

type PurchaseResponse struct {
	models.CUDResponse
	Message          string  `json:"message"`
	TotalTransaction float64 `json:"total"`
}

type PurchaseRequest struct {
	Query string `json:"query" schema:"query"`
}
