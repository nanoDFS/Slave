package acl

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nanoDFS/Slave/utils/secrets"
)

type Claims struct {
	UserId string
	FileId string
	Mode   Mode
	Size   int64
	jwt.RegisteredClaims
}

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

func (t *JWT) Generate(claim *Claims) ([]byte, error) {
	secretKeyString := secrets.Get("JWT_SECRETE_KEY")
	secretKey := []byte(secretKeyString)

	claim.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    secrets.Get("JWT_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return []byte(signedToken), nil
}

func (t *JWT) Validate(tokenString string) (*Claims, error) {
	secretKeyString := secrets.Get("JWT_SECRETE_KEY")
	secretKey := []byte(secretKeyString)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
