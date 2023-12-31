package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < KeySize {
		return nil, ErrInvalidSecretKeySize
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	var keyFunc jwt.Keyfunc = func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	payload := &Payload{}
	jwtToken, err := jwt.ParseWithClaims(token, payload, keyFunc)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			// if token expired we still return payload in case the payload used
			return payload, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
