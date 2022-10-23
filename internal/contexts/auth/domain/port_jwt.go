package domain

import (
	"context"
	"time"
)

type JWTCrypt interface {
	EncodeRefreshToken(ctx context.Context, claims JWTRefreshTokenClaims) (string, error)
	DecodeRefreshToken(ctx context.Context, token string) (JWTRefreshTokenClaims, error)

	EncodeAccessToken(ctx context.Context, claims JWTAccessTokenClaims) (string, error)
	DecodeAccessToken(ctx context.Context, token string) (JWTAccessTokenClaims, error)
}

type JWTRefreshTokenClaims struct {
	UserID   string `json:"uid"`
	ExpireAt int64  `json:"expireAt"`
}

func (jwt JWTRefreshTokenClaims) Verify() error {
	if time.Now().After(time.Unix(jwt.ExpireAt, 0)) {
		return ErrorTokenExpired
	}

	return nil
}

type JWTAccessTokenClaims struct {
	UserID   string `json:"uid"`
	ExpireAt int64  `json:"expireAt"`
}

func (jwt JWTAccessTokenClaims) Verify() error {
	if time.Now().After(time.Unix(jwt.ExpireAt, 0)) {
		return ErrorTokenExpired
	}

	return nil
}
