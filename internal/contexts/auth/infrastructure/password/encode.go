package password

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (a *Adapter) Encode(_ context.Context, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), a.config.Cost)
	if err != nil {
		panic(err)
	}

	return string(bytes), nil
}
