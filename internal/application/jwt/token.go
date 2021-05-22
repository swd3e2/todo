package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/swd3e2/todo/internal/application"
	"time"
)

type TokenBuilder struct {
	secret string
}

func NewTokenBuilder(secret string) *TokenBuilder {
	return &TokenBuilder{
		secret: secret,
	}
}

func (b *TokenBuilder) CreateToken(id string) (*application.Token, error) {
	mySigningKey := []byte(b.secret)

	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   "todo",
		Issuer:    "todo",
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := t.SignedString(mySigningKey)
	if err != nil {
		return nil, fmt.Errorf("can't sign token, err: %s", err)
	}

	return &application.Token{Token: signedToken}, nil
}
