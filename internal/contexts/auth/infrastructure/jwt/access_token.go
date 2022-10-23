package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"

	"github.com/golang-jwt/jwt/v4"
)

func accessClaimsFromDomain(claims domain.JWTAccessTokenClaims) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ID:        claims.UserID,
		ExpiresAt: jwt.NewNumericDate(time.Unix(claims.ExpireAt, 0)),
	}
}

func accessClaimsToDomain(claims jwt.RegisteredClaims) domain.JWTAccessTokenClaims {
	return domain.JWTAccessTokenClaims{
		UserID:   claims.ID,
		ExpireAt: claims.ExpiresAt.Unix(),
	}
}

func (a *Adapter) EncodeAccessToken(_ context.Context, claims domain.JWTAccessTokenClaims) (string, error) {
	registeredClaims := accessClaimsFromDomain(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)

	tokenString, err := token.SignedString([]byte(a.config.AccessTokenSecret))
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func (a *Adapter) DecodeAccessToken(_ context.Context, tokenString string) (domain.JWTAccessTokenClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.AccessTokenSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return domain.JWTAccessTokenClaims{}, domain.ErrorTokenInvalid
		} else {
			return domain.JWTAccessTokenClaims{}, domain.ErrorTokenInvalid
		}
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return domain.JWTAccessTokenClaims{}, domain.ErrorTokenInvalid
	}

	return accessClaimsToDomain(*claims), nil
}
