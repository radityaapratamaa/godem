package jwt

type Usecases interface {
	JWT() JWTs
}

type Usecase struct {
	jwt JWTs
}

func New(signingKey string) *Usecase {
	return &Usecase{
		jwt: NewJWT(signingKey),
	}
}

func (u *Usecase) JWT() JWTs {
	return u.jwt
}
