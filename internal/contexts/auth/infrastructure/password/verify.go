package password

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (a *Adapter) Verify(_ context.Context, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		panic(err)
	}

	return err == nil, nil
}
