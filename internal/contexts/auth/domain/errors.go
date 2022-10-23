package domain

import "errors"

var (
	ErrorTokenInvalid = errors.New("token invalid")
	ErrorTokenExpired = errors.New("token expired")

	ErrorUserNotFound = errors.New("user not found")

	ErrorWrongNickOrPassword = errors.New("wrong nick or password")
)
