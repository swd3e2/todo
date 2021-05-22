package application

// Token Jwt токен
type Token struct {
	Token string
}

type TokenBuilder interface {
	CreateToken(userId string) (*Token, error)
}
