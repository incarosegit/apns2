package token

import "time"
import (
	"github.com/dgrijalva/jwt-go"
	"crypto/ecdsa"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"errors"
)

// Possible errors when parsing a .p8 file.
var (
	ErrAuthKeyNotPem   = errors.New("token: AuthKey must be a valid .p8 PEM file")
	ErrAuthKeyNotECDSA = errors.New("token: AuthKey must be of type ecdsa.PrivateKey")
//	ErrAuthKeyNil      = errors.New("token: AuthKey was nil")
)


type APNSToken struct {
	key *ecdsa.PrivateKey // Private key to sign JWT
	iat int64			// Issued at
	kid string			// Key ID
	iss string			// Issuer
	raw string			// The current token
}

func ECKeyFromFile(fileName string) (*ecdsa.PrivateKey, error){

	if bytes, err := ioutil.ReadFile(fileName); err != nil {
		return nil, err
	} else {
		return ECKeyFromBytes(bytes)
	}
}

func ECKeyFromBytes(bytes []byte) (*ecdsa.PrivateKey, error){

	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(bytes); block == nil {
		return nil, ErrAuthKeyNotPem
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, ErrAuthKeyNotECDSA
	}

	return pkey, nil
}


func NewToken( key *ecdsa.PrivateKey, kid string, iss string)  *APNSToken{

	// Create the Claims
	return &APNSToken{
		key: key,
		kid: kid,
		iss: iss,
	}
}

func (t *APNSToken) Generate() (bool, error) {

	// Check if we need to regenarte ?
	issuedAt := time.Now().Unix()
	jwtToken := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": t.kid,
		},
		Claims: jwt.MapClaims{
			"iss": t.iss,
			"iat": issuedAt,
		},
		Method: jwt.SigningMethodES256,
	}
	raw, err := jwtToken.SignedString(t.key)
	if err != nil {
		return false, err
	}
	t.iat = issuedAt
	t.raw = raw
	return true, nil
}

func (t *APNSToken) Raw() string {
	return t.raw
}