package jwt

import "godem/domain/models/user"

type Claim struct {
	user.Users
	user.LoginResponse
}
