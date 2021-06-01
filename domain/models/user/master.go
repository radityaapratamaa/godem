package user

import "github.com/volatiletech/null"

type Users struct {
	ID        int64       `json:"id" db:"id"`
	Username  string      `json:"username" db:"username"`
	Password  string      `json:"password,omitempty" db:"passwd"`
	Address   string      `json:"address" db:"address"`
	CreatedAt string      `json:"created_at" db:"created_at"`
	UpdatedAt null.String `json:"updated_at" db:"updated_at"`
	DeletedAt null.String `json:"deleted_at" db:"deleted_at"`
}

type UsersRequest struct {
	Query string `json:"query" schema:"query"`
}
