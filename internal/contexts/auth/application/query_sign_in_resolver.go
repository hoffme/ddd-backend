package application

import (
	"context"
	"time"

	busDomain "github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

func (s *Service) CommandSignIn(ctx context.Context, command busDomain.Command[domain.CommandSignInData]) error {
	user, err := s.ports.UserRepository.FindByNick(ctx, command.Data.Nick)
	if err != nil {
		return domain.ErrorWrongNickOrPassword
	}

	samePassword, err := s.ports.Password.Verify(ctx, command.Data.Password, user.PasswordHash)
	if err != nil {
		return err
	}
	if !samePassword {
		return domain.ErrorWrongNickOrPassword
	}

	expireAt := time.Now().Add(s.configs.RefreshTokenExpiration).Unix()

	refreshToken, err := s.ports.JWT.EncodeRefreshToken(ctx, domain.JWTRefreshTokenClaims{
		UserID:   user.ID,
		ExpireAt: expireAt,
	})
	if err != nil {
		return err
	}

	event := domain.EventSignInDefinition.CreateEvent(domain.EventSignInData{
		CommandId:    command.ID,
		RefreshToken: refreshToken,
		ExpireAt:     expireAt,
	})

	s.ports.EventBusEmitter.Emit(ctx, event)

	return nil
}
