package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Builder struct {
	secret string
}

func NewBuilder(secret string) *Builder {
	return &Builder{
		secret: secret,
	}
}

func (b *Builder) CreateToken(id string) (string, error) {
	mySigningKey := []byte(b.secret)

	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   "todo",
		Issuer:    "todo",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
