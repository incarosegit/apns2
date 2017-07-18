package apnskey

import "time"
import (
	"github.com/dgrijalva/jwt-go"
	"crypto/ecdsa"
)

type APNSToken struct {
	key *ecdsa.PrivateKey // Private key to sign JWT
	jwt *jwt.Token
	iat int64
}

func NewToken( keybytes []byte, kid string, iss string)  (*APNSToken, error){

	// Create the Claims
	iat := time.Now().Unix()
	claims := &jwt.StandardClaims{
		IssuedAt: iat,
		Issuer:   iss,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	token.Header["kid"] = kid

	privateKey, err := jwt.ParsePKCS8PrivateKey(keybytes)
	if err != nil {
		return nil, err
	}

	return &APNSToken{
		privateKey,
		token,
		iat,
	}, nil
}

func (t *APNSToken) Generate() (string, error) {
	return t.jwt.SignedString(t.key)
}