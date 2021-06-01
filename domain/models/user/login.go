package user

type LoginRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password,omitempty" db:"passwd"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}
