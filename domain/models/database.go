package models

import "github.com/jmoiron/sqlx"

type Database struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

type SelectResponse struct {
	RequestParam  interface{} `json:"request_param"`
	Data          interface{} `json:"data"`
	TotalData     int64       `json:"total_data"`
	TotalFiltered int64       `json:"total_filtered"`
}

type CUDResponse struct {
	Status       bool  `json:"status"`
	RowsAffected int64 `json:"rows_affected"`
	LastInsertID int64 `json:"-"`
}
