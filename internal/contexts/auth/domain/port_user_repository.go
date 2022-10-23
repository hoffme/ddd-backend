package domain

import (
	"context"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (User, error)
	FindByNick(ctx context.Context, username string) (User, error)
	Save(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
}
