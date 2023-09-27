package token

import (
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	require := require.New(t)

	maker, err := NewJWTMaker(util.RandomString(34))
	require.NoError(err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(err)
	require.NotEmpty(token)

	payload, err := maker.VerifyToken(token)
	require.NoError(err)
	require.NotEmpty(payload)

	require.NotZero(payload.ID)
	require.Equal(username, payload.Username)
	require.WithinDuration(issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestJWTExpiredToken(t *testing.T) {
	require := require.New(t)

	maker, err := NewJWTMaker(util.RandomString(34))
	require.NoError(err)

	username := util.RandomOwner()
	duration := time.Minute * -20

	expiredToken, err := maker.CreateToken(username, duration)
	require.NoError(err)
	require.NotEmpty(expiredToken)

	payload, err := maker.VerifyToken(expiredToken)
	require.ErrorIs(err, ErrTokenExpired)
	require.NotEmpty(payload)
}

func TestJWTInvalidToken(t *testing.T) {
	require := require.New(t)

	maker, err := NewJWTMaker(util.RandomString(34))
	require.NoError(err)

	invalidToken := util.RandomString(32)
	payload, err := maker.VerifyToken(invalidToken)
	require.ErrorIs(err, ErrInvalidToken)
	require.Nil(payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
