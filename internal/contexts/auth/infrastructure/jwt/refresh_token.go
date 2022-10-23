package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"

	"github.com/golang-jwt/jwt/v4"
)

func refreshClaimsFromDomain(claims domain.JWTRefreshTokenClaims) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ID:        claims.UserID,
		ExpiresAt: jwt.NewNumericDate(time.Unix(claims.ExpireAt, 0)),
	}
}

func refreshClaimsToDomain(claims jwt.RegisteredClaims) domain.JWTRefreshTokenClaims {
	return domain.JWTRefreshTokenClaims{
		UserID:   claims.ID,
		ExpireAt: claims.ExpiresAt.Unix(),
	}
}

func (a *Adapter) EncodeRefreshToken(_ context.Context, claims domain.JWTRefreshTokenClaims) (string, error) {
	registeredClaims := refreshClaimsFromDomain(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)

	tokenString, err := token.SignedString([]byte(a.config.RefreshTokenSecret))
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func (a *Adapter) DecodeRefreshToken(_ context.Context, tokenString string) (domain.JWTRefreshTokenClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.RefreshTokenSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return domain.JWTRefreshTokenClaims{}, domain.ErrorTokenInvalid
		} else {
			return domain.JWTRefreshTokenClaims{}, domain.ErrorTokenInvalid
		}
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return domain.JWTRefreshTokenClaims{}, domain.ErrorTokenInvalid
	}

	return refreshClaimsToDomain(*claims), nil
}
