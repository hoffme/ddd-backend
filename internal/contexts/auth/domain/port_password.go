package domain

import "context"

type PasswordCrypt interface {
	Encode(ctx context.Context, password string) (string, error)
	Verify(ctx context.Context, password, hash string) (bool, error)
}
